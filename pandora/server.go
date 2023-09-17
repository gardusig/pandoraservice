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

func (s *PandoraServiceServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", internal.PandoraServicePort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	logrus.Debug("started listener at: ", lis.Addr())
	server := grpc.NewServer()
	pandoraproto.RegisterPandoraServiceServer(server, s)
	go startServer(server, lis)
	return nil
}

func (s *PandoraServiceServer) GuessNumber(ctx context.Context, req *pandoraproto.GuessNumberRequest) (*pandoraproto.GuessNumberResponse, error) {
	logrus.Debug("Received request at server with number: ", req.Number)
	err := validateGuessNumberRequest(req)
	if err != nil {
		return nil, err
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

func startServer(server *grpc.Server, listener net.Listener) {
	if err := server.Serve(listener); err != nil {
		panic(fmt.Errorf("pandora server failed to serve: %v", err))
	}
}

func validateGuessNumberRequest(req *pandoraproto.GuessNumberRequest) error {
	if req.Number < internal.MinThreshold {
		return fmt.Errorf("Guess must be at least %v", internal.MinThreshold)
	}
	if req.Number > internal.MaxThreshold {
		return fmt.Errorf("Guess must be at most %v", internal.MaxThreshold)
	}
	return nil
}
