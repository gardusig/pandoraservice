package internal

import (
	"context"
	"fmt"
	"time"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

const maxRetryAttempt = 3

type NumberGuesser struct {
	client pandoraproto.PandoraServiceClient

	lowerBound int64
	upperBound int64
}

func NewNumberGuesser(client pandoraproto.PandoraServiceClient) NumberGuesser {
	return NumberGuesser{
		client:     client,
		lowerBound: minThreshold,
		upperBound: maxThreshold,
	}
}

func (g *NumberGuesser) GetLockedPandoraBox() (*string, error) {
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

func (g *NumberGuesser) sendGuessRequest(guess int64) (*pandoraproto.GuessNumberResponse, error) {
	request := pandoraproto.GuessNumberRequest{
		Number: guess,
	}
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := g.client.GuessNumber(context.Background(), &request)
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
