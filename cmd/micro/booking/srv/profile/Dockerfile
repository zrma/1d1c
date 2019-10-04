FROM golang:latest AS builder
RUN mkdir /go/src/app
ADD . /go/src/app/
WORKDIR /go/src/app/

RUN GO111MODULE=on go mod init
RUN GO111MODULE=on go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o /app

FROM alpine:3.2
COPY --from=builder /app /app
WORKDIR /
ENTRYPOINT [ "/app" ]
