package internal

import (
	"context"
	"fmt"

	"github.com/gardusig/grpc_service/generated"
	"github.com/sirupsen/logrus"
)

func init() {
	shuffleRandomNumber()
}

type PandoraService struct {
	generated.UnimplementedPandoraServiceServer
}

func (s *PandoraService) GetStockPriceFluctuation(ctx context.Context, req *generated.GuessNumberRequest) (*generated.GuessNumberResponse, error) {
	guess := req.Number
	logrus.Debug("Received request at server with number: ", guess)
	if guess < minThreshold {
		return nil, fmt.Errorf("Guess must be at least %v", minThreshold)
	}
	if guess > maxThreshold {
		return nil, fmt.Errorf("Guess must be at most %v", maxThreshold)
	}
	result := validateGuess(req.Number)
	response := generated.GuessNumberResponse{
		Result: &result,
	}
	if result == equal {
		response.LockedPandoraBox = &encryptedMessage
	}
	return &response, nil
}

func validateGuess(guess int64) string {
	if guess < randomNumber {
		return less
	}
	if guess > randomNumber {
		return greater
	}
	shuffleRandomNumber()
	return equal
}
