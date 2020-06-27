package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"

	"github.com/zrma/1d1go/cmd/micro/booking/data"
	auth "github.com/zrma/1d1go/cmd/micro/booking/srv/auth/proto"
)

// Auth struct save customers information
type Auth struct {
	customers map[string]*auth.Customer
}

// VerifyToken returns a customer from authentication token.
func (s *Auth) VerifyToken(ctx context.Context, req *auth.Request, res *auth.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	customer := s.customers[req.AuthToken]
	if customer == nil {
		return errors.New("invalid Token")
	}

	res.Customer = customer
	return nil
}

// loadCustomers loads customers from a JSON file.
func loadCustomerData(path string) map[string]*auth.Customer {
	file := data.MustAsset(path)
	var customers []*auth.Customer

	// unmarshal JSON
	if err := json.Unmarshal(file, &customers); err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)
	}

	// create customer lookup map
	cache := make(map[string]*auth.Customer)
	for _, c := range customers {
		cache[c.AuthToken] = c
	}
	return cache
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.auth"),
	)

	service.Init()

	if err := auth.RegisterAuthHandler(service.Server(), &Auth{
		customers: loadCustomerData("data/customers.json"),
	}); err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
