package main

import (
	"log"

	"context"
	"github.com/micro/go-micro"
	api "github.com/micro/micro/api/proto"
)

// Redirect implements url route
type Redirect struct{}

// URL function does redirection
func (r *Redirect) URL(ctx context.Context, req *api.Request, res *api.Response) error {
	res.StatusCode = int32(301)
	res.Header = map[string]*api.Pair{
		"Location": {
			Key:    "Location",
			Values: []string{"https://google.com"},
		},
	}
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.redirect"),
	)

	// parse command line flags
	service.Init()

	if err := service.Server().Handle(
		service.Server().NewHandler(new(Redirect)),
	); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
