package pandora

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gardusig/grpc_service/guesser"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
)

func StartClient() {
	conn, err := grpc.Dial(
		fmt.Sprintf("server:%s", pandoraServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to connect: %v", err))
	}
	defer conn.Close()
	client := pandoraproto.NewPandoraServiceClient(conn)
	numberGuesser := guesser.NewNumberGuesser(client)
	lockedPandoraBox, err := numberGuesser.GetLockedPandoraBox()
	if err != nil {
		panic(fmt.Sprintf("failed to guess right number: %v", err))
	}
	logrus.Debug("encrypted message: ", *lockedPandoraBox)
}
