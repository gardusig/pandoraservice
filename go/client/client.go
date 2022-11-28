package client

import (
	"log"

	model "github.com/gardusig/template-go/grpc/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GreeterClient struct {
	Client     model.GreeterClient
	connection *grpc.ClientConn
}

func NewGreeterClient(address string) *GreeterClient {
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &GreeterClient{
		Client:     model.NewGreeterClient(conn),
		connection: conn,
	}
}

func (c *GreeterClient) Close() {
	c.connection.Close()
}
