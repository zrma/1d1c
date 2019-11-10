package main

import (
	"context"
	"encoding/json"
	"log"

	"golang.org/x/net/trace"

	"github.com/zrma/1d1c/cmd/micro/booking/data"
	"github.com/zrma/1d1c/cmd/micro/booking/srv/rate/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

type stay struct {
	HotelID string
	InDate  string
	OutDate string
}

type Rate struct {
	rateTable map[stay]*rate.RatePlan
}

// GetRates gets rates for hotels for specific date range.
func (s *Rate) GetRates(ctx context.Context, req *rate.Request, res *rate.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	for _, hotelID := range req.HotelIds {
		data := stay{
			HotelID: hotelID,
			InDate:  req.InDate,
			OutDate: req.OutDate,
		}
		if s.rateTable[data] != nil {
			res.RatePlans = append(res.RatePlans, s.rateTable[data])
		}
	}
	return nil
}

// loadRates loads rate codes from JSON file.
func loadRateTable(path string) map[stay]*rate.RatePlan {
	file := data.MustAsset("data/rates.json")

	var rates []*rate.RatePlan
	if err := json.Unmarshal(file, &rates); err != nil {
		log.Fatalf("Failed to load json: %v", err)
	}

	rateTable := make(map[stay]*rate.RatePlan)
	for _, ratePlan := range rates {
		data := stay{
			HotelID: ratePlan.HotelId,
			InDate:  ratePlan.InDate,
			OutDate: ratePlan.OutDate,
		}
		rateTable[data] = ratePlan
	}
	return rateTable
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.rate"),
	)

	service.Init()

	if err := rate.RegisterRateHandler(service.Server(), &Rate{
		rateTable: loadRateTable("data/rates.json"),
	}); err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
