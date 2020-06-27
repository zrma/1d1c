FROM ubuntu:20.04

COPY ./docker/etc/sources.list /etc/apt

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install --no-install-recommends \
        apt-utils=2.0.2ubuntu0.1 \
        ca-certificates=20190110ubuntu1.1 \
        libgl1-mesa-dev=20.0.4-2ubuntu1 \
        xorg-dev=1:7.7+19ubuntu14 \
        gcc=4:9.3.0-1ubuntu2 -y && \
    rm -fr /var/lib/apt/lists/*

ADD https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz /usr/local/
RUN tar -zxf /usr/local/go1.14.4.linux-amd64.tar.gz -C /usr/local
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

RUN mkdir -p /go/src/app/
WORKDIR /go/src/app

COPY ./ /go/src/app

RUN go test ./...
