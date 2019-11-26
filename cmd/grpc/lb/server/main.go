package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/mwitkow/grpc-proxy/proxy"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	host = ":12345"
)

type proxyHandler struct {
	opts []grpc.ServerOption

	mutex   sync.Mutex
	servers map[string]*httputil.ReverseProxy
}

func (h *proxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.Header.Get("id")
	log.Println("id:", id)

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if p, ok := h.servers[id]; ok {
		p.ServeHTTP(w, req)
		return
	}

	s := grpc.NewServer(h.opts...)
	pb.RegisterGreeterServer(s, &server{id: id})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go s.Serve(listener)

	port := strings.Split(listener.Addr().String(), ":")
	addr := "http://localhost:" + port[len(port)-1]
	u, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("invliad url: %v", u)
	}

	p := httputil.NewSingleHostReverseProxy(u)
	p.FlushInterval = -time.Second // default = 무한대, 설정 안 하면 stream send 내용을 클라로 전달하지 않게 됨.
	p.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			log.Println("dialtls:", network, addr)
			return net.Dial(network, addr)
		},
	}

	p.ServeHTTP(w, req)

	log.Println("create server :", id)
	h.servers[id] = p
}

type Server interface {
	Serve(l net.Listener) error
	Close() error
}

func main() {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//s := reverseProxyHTTP()
	s := reverseProxyGRPC()

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
	defer s.Close()
	
	log.Println("end")
}

func reverseProxyHTTP() Server {
	return &http.Server{
		Addr: host,
		Handler: h2c.NewHandler(&proxyHandler{
			opts:    buildFrontOpts(),
			servers: make(map[string]*httputil.ReverseProxy),
		}, &http2.Server{}),
	}
}

type serverAdapter struct {
	*grpc.Server
}

func (s *serverAdapter) Close() error {
	s.GracefulStop()
	return nil
}

func reverseProxyGRPC() Server {
	opts := buildFrontOpts()
	opts = append(opts, grpc.CustomCodec(proxy.Codec()))

	var mutex sync.Mutex
	servers := make(map[string]string)

	s := grpc.NewServer(opts...)
	proxy.RegisterService(s, func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		outCtx := metadata.NewOutgoingContext(context.Background(), md.Copy())

		if !ok {
			return nil, nil, status.Errorf(codes.Unimplemented, "Unknown method")
		}

		idPair, ok := md["id"]
		if !ok || len(idPair) == 0 {
			return nil, nil, status.Errorf(codes.Unimplemented, "Unknown method")
		}
		id := idPair[0]

		mutex.Lock()
		defer mutex.Unlock()

		if endpoint, ok := servers[id]; ok {
			conn, err := grpc.DialContext(outCtx, endpoint, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			return outCtx, conn, err
		}

		s := grpc.NewServer(buildBackOpts()...)
		pb.RegisterGreeterServer(s, &server{id: id})

		listener, err := net.Listen("tcp", ":0")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		go s.Serve(listener)

		endpoint := listener.Addr().String()
		conn, err := grpc.DialContext(outCtx, endpoint, grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())

		log.Println("create server :", id)

		servers[id] = endpoint
		return outCtx, conn, err
	},
		"Greeter",
		pb.GetSvcDesc()...)
	return &serverAdapter{s}
}

func buildFrontOpts() []grpc.ServerOption {
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     30 * time.Second,
			MaxConnectionAge:      1 * time.Minute,
			MaxConnectionAgeGrace: 5 * time.Second,
			// pings the client to see if the transport is still alive.
			Time:    20 * time.Second,
			Timeout: 5 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             12 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	var unaryServerInterceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("front method", info.FullMethod)

		h, err := handler(ctx, req)
		return h, err
	}
	var streamServerInterceptor grpc.StreamServerInterceptor = func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("front stream", info.FullMethod, info.IsClientStream, info.IsServerStream)

		return handler(srv, ss)
	}
	opts = append(opts, grpc.UnaryInterceptor(unaryServerInterceptor), grpc.StreamInterceptor(streamServerInterceptor))

	return opts
}

func buildBackOpts() []grpc.ServerOption {
	var opts []grpc.ServerOption
	opts = append(opts,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			// pings the client to see if the transport is still alive.
			Time:    30 * time.Second,
			Timeout: 10 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             12 * time.Second,
			PermitWithoutStream: true,
		}),
	)

	var unaryServerInterceptor grpc.UnaryServerInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("back method", info.FullMethod)

		h, err := handler(ctx, req)
		return h, err
	}
	var streamServerInterceptor grpc.StreamServerInterceptor = func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		log.Println("back stream", info.FullMethod, info.IsClientStream, info.IsServerStream)

		return handler(srv, ss)
	}
	opts = append(opts, grpc.UnaryInterceptor(unaryServerInterceptor), grpc.StreamInterceptor(streamServerInterceptor))

	return opts
}

type server struct {
	id  string
	cnt int
}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: s.id + " Hello " + in.GetName()}, nil
}

func (s *server) SayHi(req *pb.HelloRequest, stream pb.Greeter_SayHiServer) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	cnt := req.GetStartCnt()
	for stream.Context().Err() == nil {
		select {
		case <-ticker.C:
			cnt++
			if err := stream.Send(&pb.HelloReply{
				Message: "on " + s.id + ":" + req.GetName(),
				Cnt:     cnt,
			}); err != nil {
				log.Println("stream sending failed", err)
			}
		}
	}
	return nil
}
