package booking

// NOTE - run this command in $GOPATH/src

//go:generate protoc github.com/zrma/1d1c/cmd/micro/booking/srv/rate/proto/rate.proto --go_out=plugins=grpc:. --micro_out=plugins=grpc:.
//go:generate protoc github.com/zrma/1d1c/cmd/micro/booking/srv/profile/proto/profile.proto --go_out=plugins=grpc:. --micro_out=plugins=grpc:.
//go:generate protoc github.com/zrma/1d1c/cmd/micro/booking/srv/auth/proto/auth.proto --go_out=plugins=grpc:. --micro_out=plugins=grpc:.
//go:generate protoc github.com/zrma/1d1c/cmd/micro/booking/srv/geo/proto/geo.proto --go_out=plugins=grpc:. --micro_out=plugins=grpc:.
//go:generate protoc github.com/zrma/1d1c/cmd/micro/booking/api/hotel/proto/hotel.proto --go_out=plugins=grpc:. --micro_out=plugins=grpc:.
