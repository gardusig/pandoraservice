# Summary

Repository used as a foundation example of how to setup client and server using gRPC. Execution details at GitHub actions.

## Proto

#### Go

Install:

```bash
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Generate:
```bash
protoc --proto_path=./proto --go_out=./generated/go --go-grpc_out=./generated/go ./proto/stock_picker.proto
```

It should generate a few `*.pb.go` at `/generated/go/*`.


### Usage

1. [Docker Compose](#docker-compose) (C'mon, much easier life)
2. [Docker](#docker) (Otherwise...)

#### Docker Compose

As simple as that

```bash
docker-compose up
```

#### Docker

Since we're establishing a connection, it's important to create our own network :)

```bash
docker network create grpc-network
```

##### Server

```bash
docker build . -t server --progress=plain
docker run -d --name server --network grpc-network -p 50051:50051 server
```

##### Client

```bash
docker build . -t client --progress=plain
docker run --network grpc-network client
```
