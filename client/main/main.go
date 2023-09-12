package main

import (
	"flag"
	"fmt"

	"github.com/gardusig/grpc_service/client/internal"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
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
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect: %v", err))
	}
	defer conn.Close()
	client := pandoraproto.NewPandoraServiceClient(conn)
	numberGuesser := internal.NewNumberGuesser(client)
	lockedPandoraBox, err := numberGuesser.GetLockedPandoraBox()
	if err != nil {
		panic(fmt.Sprintf("failed to guess right number: %v", err))
	}
	logrus.Debug("encrypted message: ", *lockedPandoraBox)
}
