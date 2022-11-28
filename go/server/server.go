package server

import (
	"context"
	"log"

	"github.com/gardusig/template-go/example/grpc/model"
)

type GreeterServer struct {
	model.UnimplementedGreeterServer
}

func NewGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

func (s *GreeterServer) SayHello(ctx context.Context, request *model.HelloRequest) (*model.HelloReply, error) {
	log.Printf("Received sayHello request with name: %v", request.GetName())
	return &model.HelloReply{
			Message: getSayHelloFormattedMessage(request),
		},
		nil
}

func getSayHelloFormattedMessage(request *model.HelloRequest) string {
	return "Hello " + request.GetName()
}
