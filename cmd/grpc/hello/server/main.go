package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"1d1go/cmd/grpc/hello/pb"
)

const (
	endpoint = "127.0.0.1:12345"
)

func main() {
	listener, err := net.Listen("tcp", endpoint)
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

	svr := &server{}
	s := grpc.NewServer(opts...)
	pb.RegisterGreeterService(s, &pb.GreeterService{
		SayHello:      svr.SayHello,
		SayHelloAgain: svr.SayHelloAgain,
	})
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

// SayHelloAgain implements hello.GreeterServer
func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.Name}, nil
}
