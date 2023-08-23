package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"client/generated/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf("server:%d", *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance using the gRPC connection
	client := models.NewStockPickerServiceClient(conn)

	// Call the GetStockPriceFluctuation method
	req := &models.GetStockFluctuationRequest{
		// Set the required fields in the request message
	}
	resp, err := client.GetStockPriceFluctuation(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to get stock price fluctuation: %v", err)
	}

	// Process the response
	fmt.Printf("Before: %f, After: %f\n", resp.Before, resp.After)
}
