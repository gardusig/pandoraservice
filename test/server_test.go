package test

import (
	"fmt"
	"testing"

	"github.com/gardusig/grpc_service/guesser"
	"github.com/gardusig/grpc_service/internal"
	"github.com/gardusig/grpc_service/pandora"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func TestServerSetup(t *testing.T) {
	server := pandora.NewPandoraServiceServer()
	logrus.Debug("pandora server: ", server)
	err := server.Start()
	if err != nil {
		t.Fatalf("Failed to start Pandora server: %v", err)
	}
	logrus.Debug("started server")
	client, err := pandora.NewPandoraServiceClient()
	if err != nil {
		panic(fmt.Errorf("failed to connect: %v", err))
	}
	logrus.Debug("started client")
	numberGuesser := guesser.NewNumberGuesser(client.ServiceClient)
	logrus.Debug("created number guesser")
	lockedPandoraBox, err := numberGuesser.GetLockedPandoraBox()
	if err != nil {
		panic(fmt.Sprintf("failed to guess right number: %v", err))
	}
	if *lockedPandoraBox != internal.EncryptedMessage {
		t.Fatalf("Expected message: %s, received message: %s", internal.EncryptedMessage, *lockedPandoraBox)
	}
}
