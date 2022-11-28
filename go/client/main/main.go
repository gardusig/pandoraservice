package main

import (
	"context"
	"flag"
	"log"

	"github.com/gardusig/grpc-service/go/client"
	"github.com/gardusig/grpc-service/go/model/generated"
)

func main() {
	flag.Parse()
	client := client.NewGreeterClient("localhost:50051")
	r, err := client.Client.SayHello(
		context.Background(),
		&generated.HelloRequest{
			Name: "world",
		},
	)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("Response: %s", r.GetMessage())
}
