package booking

//go:generate protoc srv/rate/proto/rate.proto --go_out=plugins=grpc:.
//go:generate protoc srv/profile/proto/profile.proto --go_out=plugins=grpc:.
//go:generate protoc api/hotel/proto/hotel.proto --go_out=plugins=grpc:.
