module helloworld

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro/v3 v3.0.0-beta.0.20200817090452-870a1ebc07bb
	github.com/micro/micro/v3 v3.0.0-beta.0.20200817131840-755a1e3d06ce
	google.golang.org/protobuf v1.25.0
)
