# Helloworld Service

This is the Helloworld service

Generated with

```
micro new helloworld
```

## Requirement

```
# protoc
(linux) brew install protobuf
(windows) scoop install protobuf

# protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go

# protoc-gen-micro
go get github.com/micro/micro/v3/cmd/protoc-gen-micro@master
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- Alias: helloworld

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service

```
./helloworld
```

Build a docker image

```
make docker
```
