package main

import (
	"log"
	"net"

	"github.com/zrma/1d1c/cmd/grpc/lb/server/proxy"
	"github.com/zrma/1d1c/cmd/grpc/lb/server/service"
)

const (
	host = ":12345"
)

func main() {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var svcFactory proxy.ServiceFactory = func(id string) *service.Service {
		return &service.Service{Id: id}
	}
	//s := proxy.ReverseProxyHTTP(listener.Addr().String(), svcFactory)
	s := proxy.ReverseProxyGRPC(svcFactory)

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
	defer s.Close()

	log.Println("end")
}
