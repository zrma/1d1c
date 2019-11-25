package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

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

	go func() {
		for ctx.Err() == nil {
			func() {
				ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("Greeting: %s", r.GetMessage())

				time.Sleep(time.Second * 10)
			}()
		}
	}()

	waitCh := make(chan struct{})
	go func() {
		defer close(waitCh)

		errCount := 0
		const maxSleep = 30 * time.Second
		var cnt int64

		for ctx.Err() == nil {
			md := metadata.Pairs("id", fmt.Sprintf("%06d", idx))
			ctx2 := metadata.NewOutgoingContext(context.Background(), md)

			if errCount > 1 {
				// 에러가 발생하는 경우 재시도 간격을 지수적으로 증가시킨다:
				dur := time.Duration(1<<uint(errCount)) * 100 * time.Millisecond
				if dur > maxSleep {
					dur = maxSleep
				}

				select {
				case <-time.After(dur):
				case <-ctx2.Done():
				}
			}

			stream, err := c.SayHi(ctx2, &pb.HelloRequest{
				Name:     defaultName,
				StartCnt: cnt,
			})
			if err != nil {
				log.Printf("stream connection failed: %v", err)
			}

			for ctx2.Err() == nil {
				r, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					errCount++
					if statusCode := status.Code(err); statusCode != codes.Canceled {
						log.Println("stream recv err", err)
					}
					break
				}

				cnt = r.GetCnt()
				fmt.Println("recv", r.GetMessage(), cnt)
			}
		}

		if err := ctx.Err(); err != context.Canceled {
			log.Println("ctx err:", err)
		}
	}()

	<-waitCh

	time.Sleep(time.Second * 3)
	{
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	time.Sleep(time.Second * 3)
	{
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: defaultName})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	log.Println("client", idx, "closed")
}
