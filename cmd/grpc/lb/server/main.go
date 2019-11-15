package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	port = ":12345"
)

type Handler struct {
	opts []grpc.ServerOption
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.Header.Get("id")
	log.Println("id:", id)

	s := grpc.NewServer(h.opts...)
	pb.RegisterGreeterServer(s, &server{id: id})

	s.ServeHTTP(w, req)
}

func main() {
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

	//if err := s.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}

	log.Fatalln(http.ListenAndServe(port,
		h2c.NewHandler(&Handler{opts: opts}, &http2.Server{}),
	))
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
