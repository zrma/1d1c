package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	address     = "localhost:12345"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			// keepalive settings - https://github.com/grpc/grpc/blob/master/doc/keepalive.md
			Time:                10 * time.Second,
			Timeout:             5 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		log.Fatalf("grpc connection failed: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	waitCh := make(chan struct{})
	go func() {
		defer close(waitCh)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		stream, err := c.SayHi(ctx, &pb.HelloRequest{
			Name: defaultName,
		})
		if err != nil {
			log.Printf("stream connection failed: %v", err)
		}

		for ctx.Err() == nil {
			r, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("stream recv err", err)
				return
			}
			fmt.Println("recv", r.GetMessage())
		}
	}()

	<-waitCh
	log.Println("end")
	time.Sleep(time.Second)
}
