package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"server/generated/models"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type StockService struct {
	models.UnimplementedStockPickerServiceServer
}

func (s *StockService) GetStockPriceFluctuation(ctx context.Context, req *models.GetStockFluctuationRequest) (*models.GetStockFluctuationResponse, error) {
	// Your implementation to handle the gRPC request goes here
	// For example, you can perform some computations and return the response
	log.Printf("Received request :)")
	return &models.GetStockFluctuationResponse{
		Before: 100.0,
		After:  120.0,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	models.RegisterStockPickerServiceServer(s, &StockService{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
