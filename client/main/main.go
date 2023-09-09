package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gardusig/grpc_service/generated"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	conn, err := grpc.Dial(
		fmt.Sprintf("server:%d", *port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := generated.NewPandoraServiceClient(conn)
	request := generated.GuessNumberRequest{
		Number: 420,
	}
	resp, err := client.GuessNumber(context.Background(), &request)
	if err != nil {
		log.Fatalf("failed to get stock price fluctuation: %v", err)
	}
	logrus.Debug("response: ", *resp.Result)
}
