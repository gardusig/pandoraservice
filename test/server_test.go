package test

import (
	"testing"

	"github.com/gardusig/pandoraservice/pandora"
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
	response, err := client.SendGuessRequest(0, 0)
	if err != nil {
		t.Fatalf("failed to make guess request: %v", err)
	}
	logrus.Debug("server response: ", response.Result)
}
