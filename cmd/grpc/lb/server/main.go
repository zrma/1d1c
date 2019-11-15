package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	port = ":12345"
)

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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

	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
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
				Message: req.GetName(),
			}); err != nil {
				log.Println("stream sending failed", err)
			}
		}
	}
	return nil
}
