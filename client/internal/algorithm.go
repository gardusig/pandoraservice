package internal

import (
	"context"
	"fmt"

	"github.com/gardusig/grpc_service/generated"
	"github.com/sirupsen/logrus"
)

func GetLockedPandoraBox(client generated.PandoraServiceClient) (*string, error) {
	lowerBound := minThreshold
	upperBound := maxThreshold
	for lowerBound <= upperBound {
		mid := lowerBound + ((upperBound - lowerBound) >> 1)
		logrus.Debug("lowerBound:", lowerBound, ", upperBound:", upperBound, ", mid:", mid)
		request := generated.GuessNumberRequest{
			Number: int64(mid),
		}
		resp, err := client.GuessNumber(context.Background(), &request)
		if err != nil {
			return nil, err
		}
		logrus.Debug("server response:", *resp.Result)
		if *resp.Result == equal {
			return resp.LockedPandoraBox, nil
		}
		if *resp.Result == greater {
			upperBound = mid - 1
		} else if *resp.Result == less {
			lowerBound = mid + 1
		} else {
			return nil, fmt.Errorf("Unexpected response from server: %v", *resp.Result)
		}
	}
	return nil, fmt.Errorf("Failed to get right number :/")
}
