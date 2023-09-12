package internal

import (
	"context"
	"fmt"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/sirupsen/logrus"
)

func init() {
	shuffleRandomNumber()
}

type PandoraService struct {
	pandoraproto.UnimplementedPandoraServiceServer
}

func (s *PandoraService) GuessNumber(ctx context.Context, req *pandoraproto.GuessNumberRequest) (*pandoraproto.GuessNumberResponse, error) {
	guess := req.Number
	logrus.Debug("Received request at server with number: ", guess)
	if guess < minThreshold {
		return nil, fmt.Errorf("Guess must be at least %v", minThreshold)
	}
	if guess > maxThreshold {
		return nil, fmt.Errorf("Guess must be at most %v", maxThreshold)
	}
	result := validateGuess(req.Number)
	response := pandoraproto.GuessNumberResponse{
		Result:           result,
		LockedPandoraBox: nil,
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
