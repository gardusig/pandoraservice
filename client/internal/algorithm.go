package internal

import (
	"context"
	"fmt"

	"github.com/gardusig/grpc_service/generated"
	"github.com/sirupsen/logrus"
)

const maxRetryAttempt = 3

type NumberGuesser struct {
	client generated.PandoraServiceClient

	lowerBound int64
	upperBound int64
}

func NewNumberGuesser(client generated.PandoraServiceClient) NumberGuesser {
	return NumberGuesser{
		client:     client,
		lowerBound: minThreshold,
		upperBound: maxThreshold,
	}
}

func (g NumberGuesser) GetLockedPandoraBox() (*string, error) {
	for g.lowerBound <= g.upperBound {
		response, err := g.makeNextGuess()
		if err != nil {
			return nil, err
		}
		if response != nil {
			return response, nil
		}
	}
	return nil, fmt.Errorf("Failed to get right number :/")
}

func (g *NumberGuesser) makeNextGuess() (*string, error) {
	guess := g.lowerBound + ((g.upperBound - g.lowerBound) >> 1)
	logrus.Debug("lowerBound:", g.lowerBound, ", upperBound:", g.upperBound, ", guess:", guess)
	resp, err := g.sendGuessRequest(guess)
	if resp == nil || err != nil {
		return nil, err
	}
	logrus.Debug("server response:", resp.Result)
	if resp.Result == equal {
		return resp.LockedPandoraBox, nil
	}
	return nil, g.updateBoundaries(guess, resp.Result)
}

func (g *NumberGuesser) updateBoundaries(guess int64, response string) error {
	if response == greater {
		g.upperBound = guess - 1
		return nil
	}
	if response == less {
		g.lowerBound = guess + 1
		return nil
	}
	return fmt.Errorf("Unexpected response from server: %v", response)
}

func (g NumberGuesser) sendGuessRequest(guess int64) (*generated.GuessNumberResponse, error) {
	request := generated.GuessNumberRequest{
		Number: guess,
	}
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		resp, err := g.client.GuessNumber(context.Background(), &request)
		if err == nil {
			return resp, nil
		}
		// if grpc.StatusCode(err) == grpc.StatusCode.Unavailable {
		// 	continue
		// }
		return nil, err
	}
	return nil, fmt.Errorf("Failed to make guess request after %v attempts", maxRetryAttempt)
}
