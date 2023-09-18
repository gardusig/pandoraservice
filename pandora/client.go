package pandora

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/gardusig/pandoraservice/internal"
)

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

func (c *PandoraServiceClient) CloseConnection() {
	c.connection.Close()
}
