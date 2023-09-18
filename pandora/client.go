package pandora

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/gardusig/pandoraservice/internal"
)

const maxRetryAttempt = 5

type PandoraServiceClient struct {
	pandoraproto.PandoraServiceClient

	connection *grpc.ClientConn
}

func NewPandoraServiceClient() (*PandoraServiceClient, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("localhost:%s", internal.PandoraServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}
	client := pandoraproto.NewPandoraServiceClient(conn)
	return &PandoraServiceClient{
		PandoraServiceClient: client,
		connection:           conn,
	}, nil
}

func (c *PandoraServiceClient) SendGuessRequest(level uint32, guess int64) (*pandoraproto.GuessNumberResponse, error) {
	request := pandoraproto.GuessNumberRequest{
		Level:  level,
		Number: guess,
	}
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := c.GuessNumber(context.Background(), &request)
		if err == nil {
			return resp, nil
		}
		_, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("Failed to make guess request after %v attempts", maxRetryAttempt)
}

func (c *PandoraServiceClient) SendOpenBoxRequest(lockedPandoraBox *pandoraproto.LockedPandoraBox) (*pandoraproto.OpenedPandoraBox, error) {
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := c.OpenBox(context.Background(), lockedPandoraBox)
		if err == nil {
			return resp, nil
		}
		_, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("Failed to make openBox request after %v attempts", maxRetryAttempt)
}

func (c *PandoraServiceClient) CloseConnection() {
	c.connection.Close()
}
