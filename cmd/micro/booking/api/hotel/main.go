package main

import (
	"context"
	"encoding/base64"
	"errors"
	_ "expvar"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strings"

	"golang.org/x/net/trace"

	"github.com/google/uuid"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	. "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"

	"github.com/zrma/1d1go/cmd/micro/booking/api/hotel/proto"
	"github.com/zrma/1d1go/cmd/micro/booking/srv/auth/proto"
	"github.com/zrma/1d1go/cmd/micro/booking/srv/geo/proto"
	"github.com/zrma/1d1go/cmd/micro/booking/srv/profile/proto"
	"github.com/zrma/1d1go/cmd/micro/booking/srv/rate/proto"
)

const (
	// BasicSchema is prefix of basic schema
	BasicSchema string = "Basic "

	// BearerSchema is prefix of auth bearer
	BearerSchema string = "Bearer "
)

type profileResults struct {
	hotels []*profile.Hotel
	err    error
}

type rateResults struct {
	ratePlans []*rate.RatePlan
	err       error
}

// Hotel struct implements booking server
type Hotel struct {
	Client client.Client
}

// Rates method register client's order
func (s *Hotel) Rates(ctx context.Context, req *hotel.Request, res *hotel.Response) error {
	// tracing
	tr := trace.New("api.v1", "Hotel.Rates")
	defer tr.Finish()

	// context
	ctx = trace.NewContext(ctx, tr)

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	// add a unique request id to context
	traceID := uuid.New()
	// make copy
	tmd := metadata.Metadata{}
	for k, v := range md {
		tmd[k] = v
	}

	tmd["traceID"] = traceID.String()
	tmd["fromName"] = "api.v1"
	ctx = metadata.NewContext(ctx, tmd)

	// token from request headers
	token, err := getToken(md)
	if err != nil {
		return Forbidden("api.hotel.rates", err.Error())
	}

	// verify token w/ auth service
	authClient := auth.NewAuthService("go.micro.srv.auth", s.Client)
	if _, err = authClient.VerifyToken(ctx, &auth.Request{AuthToken: token}); err != nil {
		return Unauthorized("api.hotel.rates", "Unauthorized")
	}

	// checkin and checkout date query params
	inDate, outDate := req.InDate, req.OutDate
	if inDate == "" || outDate == "" {
		return BadRequest("api.hotel.rates", "Please specify inDate/outDate params")
	}

	// finds nearby hotels
	// TODO(hw): use lat/lon from request params
	geoClient := geo.NewGeoService("go.micro.srv.geo", s.Client)
	nearby, err := geoClient.Nearby(ctx, &geo.Request{
		Lat: 51.502973,
		Lon: -0.114723,
	})
	if err != nil {
		return InternalServerError("api.hotel.rates", err.Error())
	}

	// make requests for profiles and rates
	profileCh := getHotelProfiles(ctx, s.Client, nearby.HotelIds)
	rateCh := getRatePlans(ctx, s.Client, nearby.HotelIds, inDate, outDate)

	// wait on profiles reply
	profileReply := <-profileCh
	if err := profileReply.err; err != nil {
		return InternalServerError("api.hotel.rates", err.Error())
	}

	// wait on rates reply
	rateReply := <-rateCh
	if err := rateReply.err; err != nil {
		return InternalServerError("api.hotel.rates", err.Error())
	}

	res.Hotels = profileReply.hotels
	res.RatePlans = rateReply.ratePlans
	return nil
}

func getToken(md metadata.Metadata) (string, error) {
	// Grab the raw authorization header
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", errors.New("authorization header required")
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BasicSchema) && !strings.HasPrefix(authHeader, BearerSchema) {
		return "", errors.New("authorization requires Basic/Bearer scheme")
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, BasicSchema) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(BasicSchema):])
		if err != nil {
			return "", errors.New("base64 encoding issue")
		}
		credential := strings.Split(string(str), ":")
		return credential[0], nil
	}

	return authHeader[len(BearerSchema):], nil
}

func getRatePlans(ctx context.Context, c client.Client, hotelIDs []string, inDate string, outDate string) chan rateResults {
	rateClient := rate.NewRateService("go.micro.srv.rate", c)
	ch := make(chan rateResults, 1)

	go func() {
		var ratePlans []*rate.RatePlan
		res, err := rateClient.GetRates(ctx, &rate.Request{
			HotelIds: hotelIDs,
			InDate:   inDate,
			OutDate:  outDate,
		})
		if res != nil {
			ratePlans = append(ratePlans, res.RatePlans...)
		}
		ch <- rateResults{ratePlans, err}
	}()

	return ch
}

func getHotelProfiles(ctx context.Context, c client.Client, hotelIDs []string) chan profileResults {
	profileClient := profile.NewProfileService("go.micro.srv.profile", c)
	ch := make(chan profileResults, 1)

	go func() {
		var hotels []*profile.Hotel
		res, err := profileClient.GetProfiles(ctx, &profile.Request{
			HotelIds: hotelIDs,
			Locale:   "en",
		})
		if res != nil {
			hotels = append(hotels, res.Hotels...)
		}
		ch <- profileResults{hotels, err}
	}()

	return ch
}

func main() {
	// trace library patched for demo purposes.
	// https://github.com/golang/net/blob/master/trace/trace.go#L94
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	service := micro.NewService(
		micro.Name("go.micro.api.hotel"),
	)

	service.Init()
	if err := hotel.RegisterHotelHandler(service.Server(), &Hotel{service.Client()}); err != nil {
		log.Fatalln(err)
	}

	if err := service.Run(); err != nil {
		log.Fatalln(err)
	}
}
