package pandora

import (
	"context"
	"fmt"
	"net"

	"github.com/gardusig/grpc_service/database"
	"github.com/gardusig/grpc_service/internal"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type PandoraServiceServer struct {
	pandoraproto.UnimplementedPandoraServiceServer

	db *database.SpecialNumberDb
}

func NewPandoraServiceServer() *PandoraServiceServer {
	return &PandoraServiceServer{
		db: database.NewSpecialNumberDb(),
	}
}

func StartPandoraServer() error {
	serverChan := make(chan error, 1)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", internal.PandoraServicePort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	logrus.Debug("started listener at: ", lis.Addr())
	server := grpc.NewServer()
	pandoraproto.RegisterPandoraServiceServer(server, &PandoraServiceServer{})
	logrus.Debug("created channel")
	go startServer(server, lis, serverChan)
	go monitorServerStatus(serverChan)
	return nil
}

func (s *PandoraServiceServer) GuessNumber(ctx context.Context, req *pandoraproto.GuessNumberRequest) (*pandoraproto.GuessNumberResponse, error) {
	guess := req.Number
	logrus.Debug("Received request at server with number: ", guess)
	if guess < internal.MinThreshold {
		return nil, fmt.Errorf("Guess must be at least %v", internal.MinThreshold)
	}
	if guess > internal.MaxThreshold {
		return nil, fmt.Errorf("Guess must be at most %v", internal.MaxThreshold)
	}
	result := s.db.ValidateGuess(req.Number)
	response := pandoraproto.GuessNumberResponse{
		Result:           result,
		LockedPandoraBox: nil,
	}
	if result == internal.Equal {
		response.LockedPandoraBox = &internal.EncryptedMessage
	}
	return &response, nil
}

func startServer(server *grpc.Server, listener net.Listener, serverChan chan error) {
	serverChan <- server.Serve(listener)
}

func monitorServerStatus(serverChan chan error) {
	select {
	case err := <-serverChan:
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
