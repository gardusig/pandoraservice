package pandora

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gardusig/grpc_service/internal"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
)

type PandoraServiceClient struct {
	connection    *grpc.ClientConn
	ServiceClient pandoraproto.PandoraServiceClient
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
		connection:    conn,
		ServiceClient: client,
	}, nil
}

func (c *PandoraServiceClient) CloseConnection() {
	c.connection.Close()
}
