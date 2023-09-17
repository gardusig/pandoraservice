package test

import (
	"testing"

	"github.com/gardusig/grpc_service/guesser"
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
		t.Fatalf("failed to connect: %v", err)
	}
	logrus.Debug("started client")
	guesser := guesser.NewGuesser(client.ServiceClient)
	logrus.Debug("created number guesser")
	openedPandoraBox, err := guesser.GetPandoraBox()
	if err != nil {
		t.Fatalf("failed to guess right number: %v", err)
	}
	if openedPandoraBox == nil {
		t.Fatalf("expected opened pandora box, got nil instead")
	}
	logrus.Debug("message: ", openedPandoraBox.Message)
}
