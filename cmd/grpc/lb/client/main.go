package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"

	"github.com/zrma/1d1c/cmd/grpc/lb/pb"
)

const (
	address     = ":12345"
	defaultName = "world"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			time.Sleep(time.Second * time.Duration(idx))
			connect(ctx, idx)
			wg.Done()
		}(i)
	}

	wg.Wait()
	log.Println("end")
	time.Sleep(time.Second)
}

func connect(ctx context.Context, idx int) {
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

	md := metadata.Pairs("id", fmt.Sprintf("%06d", idx))
	ctx = metadata.NewOutgoingContext(ctx, md)

	{
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	waitCh := make(chan struct{})
	go func() {
		defer close(waitCh)

		ctx, cancel := context.WithCancel(ctx)
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

	log.Println("client", idx, "closed")
}
