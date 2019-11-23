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

type ProxyHandler struct {
	opts []grpc.ServerOption
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.Header.Get("id")
	log.Println("id:", id)

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

	w.Header().Set("X-Forwarded-For", req.Host)
	p.ServeHTTP(w, req)
}

type Server interface {
	Serve(l net.Listener) error
}

func main() {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := reverseProxyHTTP()
	//s := reverseProxyGRPC()

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}

	log.Println("end")
}

func reverseProxyHTTP() Server {
	return &http.Server{
		Addr:    host,
		Handler: h2c.NewHandler(&ProxyHandler{opts: buildFrontOpts()}, &http2.Server{}),
	}
}

func reverseProxyGRPC() Server {
	opts := buildFrontOpts()
	opts = append(opts, grpc.CustomCodec(proxy.Codec()))

	s := grpc.NewServer(opts...)
	proxy.RegisterService(s, func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		outCtx := metadata.NewOutgoingContext(ctx, md.Copy())

		if ok {
			id := md["id"][0]

			s := grpc.NewServer(buildBackOpts()...)
			pb.RegisterGreeterServer(s, &server{id: id})

			listener, err := net.Listen("tcp", ":0")
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}

			go s.Serve(listener)

			conn, err := grpc.DialContext(ctx, listener.Addr().String(), grpc.WithCodec(proxy.Codec()), grpc.WithInsecure())
			return outCtx, conn, err
		}
		return nil, nil, status.Errorf(codes.Unimplemented, "Unknown method")
	},
		"Greeter",
		pb.GetSvcDesc()...)
	return s
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
	return opts
}

type server struct {
	id string
}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHi(req *pb.HelloRequest, stream pb.Greeter_SayHiServer) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for stream.Context().Err() == nil {
		select {
		case <-ticker.C:
			if err := stream.Send(&pb.HelloReply{
				Message: "on " + s.id + ":" + req.GetName(),
			}); err != nil {
				log.Println("stream sending failed", err)
			}
		}
	}
	return nil
}
