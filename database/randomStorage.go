package database

import (
	"math/rand"

	"github.com/gardusig/grpc_service/internal"
)

var randomNumber int64

func ShuffleSpecialNumber() {
	randomNumber = rand.Int63n(internal.MaxThreshold-internal.MinThreshold+1) + internal.MinThreshold
}

func ValidateGuess(guess int64) string {
	if guess < randomNumber {
		return internal.Less
	}
	if guess > randomNumber {
		return internal.Greater
	}
	ShuffleSpecialNumber()
	return internal.Equal
}
