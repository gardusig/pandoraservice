package guesser

import (
	"context"
	"fmt"
	"time"

	"github.com/gardusig/grpc_service/internal"
	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

const maxRetryAttempt = 5

type NumberGuesser struct {
	client pandoraproto.PandoraServiceClient

	level      uint32
	lowerBound int64
	upperBound int64
}

func NewNumberGuesser(client pandoraproto.PandoraServiceClient) NumberGuesser {
	return NumberGuesser{
		client: client,
	}
}

func (g *NumberGuesser) GetLockedPandoraBox() (encryptedMessage *string, err error) {
	g.level = internal.LevelMinThreshold
	for g.level <= internal.LevelMaxThreshold {
		encryptedMessage, err = g.guessNumberByLevel()
		if err != nil {
			return nil, err
		}
		logrus.Debug("Passed level: ", g.level, ", message: ", *encryptedMessage)
		g.level += 1
	}
	return encryptedMessage, nil
}

func (g *NumberGuesser) guessNumberByLevel() (*string, error) {
	logrus.Debug("attempt to guess number for level: ", g.level)
	g.lowerBound = internal.GuessMinThreshold
	g.upperBound = internal.GuessMaxThreshold
	for g.lowerBound <= g.upperBound {
		encryptedMessage, err := g.makeNextGuess()
		if encryptedMessage != nil || err != nil {
			return encryptedMessage, err
		}
	}
	return nil, fmt.Errorf("Failed to guess the right number :/")
}

func (g *NumberGuesser) makeNextGuess() (*string, error) {
	guess := g.lowerBound + ((g.upperBound - g.lowerBound) >> 1)
	logrus.Debug("lowerBound:", g.lowerBound, ", upperBound:", g.upperBound, ", guess:", guess)
	resp, err := g.sendGuessRequest(guess)
	if resp == nil || err != nil {
		return nil, err
	}
	logrus.Debug("server response:", resp.Result)
	if resp.Result == internal.Equal {
		return resp.LockedPandoraBox, nil
	}
	return nil, g.updateBoundaries(guess, resp.Result)
}

func (g *NumberGuesser) updateBoundaries(guess int64, response string) error {
	if response == internal.Greater {
		g.upperBound = guess - 1
		return nil
	}
	if response == internal.Less {
		g.lowerBound = guess + 1
		return nil
	}
	return fmt.Errorf("Unexpected response from server: %v", response)
}

func (g *NumberGuesser) sendGuessRequest(guess int64) (*pandoraproto.GuessNumberResponse, error) {
	request := pandoraproto.GuessNumberRequest{
		Number: guess,
		Level:  g.level,
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
