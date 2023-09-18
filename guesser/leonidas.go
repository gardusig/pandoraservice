package guesser

import (
	"context"
	"fmt"
	"time"

	pandoraproto "github.com/gardusig/pandoraproto/generated/go"
	"github.com/gardusig/pandoraservice/internal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

const maxRetryAttempt = 5

type Guesser struct {
	client pandoraproto.PandoraServiceClient

	level      uint32
	lowerBound int64
	upperBound int64
}

func NewGuesser(client pandoraproto.PandoraServiceClient) Guesser {
	return Guesser{
		client: client,
	}
}

func (g *Guesser) GetPandoraBox() (*pandoraproto.OpenedPandoraBox, error) {
	var lockedBox *pandoraproto.LockedPandoraBox
	var err error
	g.level = internal.LevelMinThreshold
	for g.level <= internal.LevelMaxThreshold {
		lockedBox, err = g.guessNumberByLevel()
		if err != nil {
			return nil, err
		}
		if lockedBox != nil {
			g.level += 1
			logrus.Debug("Passed to level: ", g.level, ", encryptedMessage: ", lockedBox.EncryptedMessage)
		}
	}
	return g.getOpenedPandoraBox(lockedBox)
}

func (g *Guesser) getOpenedPandoraBox(lockedPandoraBox *pandoraproto.LockedPandoraBox) (*pandoraproto.OpenedPandoraBox, error) {
	for attempt := 0; attempt < maxRetryAttempt; attempt += 1 {
		if attempt > 0 {
			time.Sleep(1 << time.Duration(attempt) * time.Second)
		}
		resp, err := g.client.OpenBox(context.Background(), lockedPandoraBox)
		if err == nil {
			return resp, nil
		}
		_, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
	}
	return nil, fmt.Errorf("Failed to make openBox request after %v attempts", maxRetryAttempt)
}

func (g *Guesser) guessNumberByLevel() (*pandoraproto.LockedPandoraBox, error) {
	logrus.Debug("attempt to guess number for level: ", g.level)
	g.lowerBound = internal.GuessMinThreshold
	g.upperBound = internal.GuessMaxThreshold
	for g.lowerBound <= g.upperBound {
		guess := g.lowerBound + ((g.upperBound - g.lowerBound) >> 1)
		logrus.Debug("lowerBound:", g.lowerBound, ", upperBound:", g.upperBound, ", guess:", guess)
		resp, err := g.sendGuessRequest(guess)
		if err != nil {
			return nil, err
		}
		logrus.Debug("server response:", resp.Result)
		switch resp.Result {
		case internal.Equal:
			return resp.LockedPandoraBox, nil
		case internal.Greater:
			g.upperBound = guess - 1
		case internal.Less:
			g.lowerBound = guess + 1
		default:
			return nil, fmt.Errorf("Unexpected response from server: %v", resp.Result)
		}
	}
	return nil, fmt.Errorf("Failed to guess the right number :/")
}

func (g *Guesser) sendGuessRequest(guess int64) (*pandoraproto.GuessNumberResponse, error) {
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
