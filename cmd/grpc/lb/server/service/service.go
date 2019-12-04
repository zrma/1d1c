package service

import (
	"context"
	"log"
	"time"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

type Service struct {
	Id  string
	cnt int
}

// SayHello implements hello.GreeterServer
func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: s.Id + " Hello " + in.GetName()}, nil
}

func (s *Service) SayHi(req *pb.HelloRequest, stream pb.Greeter_SayHiServer) error {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	cnt := req.GetStartCnt()
	for stream.Context().Err() == nil {
		select {
		case <-ticker.C:
			cnt++
			if err := stream.Send(&pb.HelloReply{
				Message: "on " + s.Id + ":" + req.GetName(),
				Cnt:     cnt,
			}); err != nil {
				log.Println("stream sending failed", err)
			}
		}
	}
	return nil
}
