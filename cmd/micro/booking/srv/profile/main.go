package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/zrma/1d1go/cmd/micro/booking/data"
	"github.com/zrma/1d1go/cmd/micro/booking/srv/profile/proto"

	"golang.org/x/net/trace"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

// Profile struct serve hotel list
type Profile struct {
	hotels map[string]*profile.Hotel
}

// GetProfiles returns hotel profiles for requested IDs
func (s *Profile) GetProfiles(ctx context.Context, req *profile.Request, res *profile.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	for _, i := range req.HotelIds {
		res.Hotels = append(res.Hotels, s.hotels[i])
	}
	return nil
}

// loadProfiles loads hotel profiles from a JSON file.
func loadProfiles(path string) map[string]*profile.Hotel {
	file := data.MustAsset(path)

	// unmarshal json profiles
	var hotels []*profile.Hotel
	if err := json.Unmarshal(file, &hotels); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	profiles := make(map[string]*profile.Hotel)
	for _, hotel := range hotels {
		profiles[hotel.Id] = hotel
	}
	return profiles
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.profile"),
	)

	service.Init()

	if err := profile.RegisterProfileHandler(service.Server(), &Profile{
		hotels: loadProfiles("data/profiles.json"),
	}); err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
