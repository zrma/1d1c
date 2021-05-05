FROM ubuntu:focal-20210416

COPY ./docker/etc/sources.list /etc/apt

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get upgrade -y
RUN apt-get -y install --no-install-recommends ca-certificates=20210119~20.04.1
RUN apt-get -y install --no-install-recommends libgl1-mesa-dev=20.2.6-0ubuntu0.20.04.1 xorg-dev=1:7.7+19ubuntu14 gcc=4:9.3.0-1ubuntu2
RUN apt-get -y install --no-install-recommends wget=1.20.3-1ubuntu1 curl=7.68.0-1ubuntu2.5

RUN apt-get autoremove -y
RUN apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*

ARG GO_VERSION=1.16

ADD https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz /usr/local/
RUN tar -zxf /usr/local/go${GO_VERSION}.linux-amd64.tar.gz -C /usr/local
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.31.0

RUN mkdir -p /go/src/app/
WORKDIR /go/src/app

COPY ./ /go/src/app

RUN golangci-lint run --timeout 10m
RUN go test -race ./...
