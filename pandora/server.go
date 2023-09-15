package pandora

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/gardusig/grpc_service/database"
	"github.com/gardusig/grpc_service/internal"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	database.ShuffleSpecialNumber()
}

func StartPandoraServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", pandoraServicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pandoraproto.RegisterPandoraServiceServer(s, &PandoraServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type PandoraServiceServer struct {
	pandoraproto.UnimplementedPandoraServiceServer
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
	result := database.ValidateGuess(req.Number)
	response := pandoraproto.GuessNumberResponse{
		Result:           result,
		LockedPandoraBox: nil,
	}
	if result == internal.Equal {
		response.LockedPandoraBox = &internal.EncryptedMessage
	}
	return &response, nil
}
