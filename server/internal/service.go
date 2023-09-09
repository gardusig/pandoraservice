package internal

import (
	"context"

	"github.com/gardusig/grpc_service/generated"
	"github.com/sirupsen/logrus"
)

type PandoraService struct {
	generated.UnimplementedPandoraServiceServer
}

func (s *PandoraService) GetStockPriceFluctuation(ctx context.Context, req *generated.GuessNumberRequest) (*generated.GuessNumberResponse, error) {
	logrus.Debug("Received request at server with number: ", req.Number)
	return &generated.GuessNumberResponse{
		Result: validateGuess(req.Number),
	}, nil
}
