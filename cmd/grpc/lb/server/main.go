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

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	host = ":12345"
)

type Handler struct {
	proxy *httputil.ReverseProxy
	opts  []grpc.ServerOption
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.Header.Get("id")
	log.Println("id:", id)

	s := grpc.NewServer(h.opts...)
	pb.RegisterGreeterServer(s, &server{id: id})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go s.Serve(listener)

	tmp := strings.Split(listener.Addr().String(), ":")
	addr := "http://localhost:" + tmp[len(tmp)-1]
	endpoint, err := url.Parse(addr)
	if err != nil {
		log.Fatalf("invliad url: %v", endpoint)
	}

	p := httputil.NewSingleHostReverseProxy(endpoint)
	p.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
			log.Println("dialtls:", network, addr)
			ta, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}
			return net.DialTCP(network, nil, ta)
		},
	}
	p.ServeHTTP(w, req)
}

func main() {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &http.Server{
		Addr:    host,
		Handler: h2c.NewHandler(&Handler{opts: buildOpts()}, &http2.Server{}),
	}

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
	log.Println("end")
}

func buildOpts() []grpc.ServerOption {
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
