package database

import (
	"math/rand"

	"github.com/gardusig/grpc_service/internal"
)

type SpecialNumberDb struct {
	randomNumber int64
}

func NewSpecialNumberDb() *SpecialNumberDb {
	return &SpecialNumberDb{
		randomNumber: getRandomNumber(),
	}
}

func (db *SpecialNumberDb) ValidateGuess(guess int64) string {
	if guess < db.randomNumber {
		return internal.Less
	}
	if guess > db.randomNumber {
		return internal.Greater
	}
	db.randomNumber = getRandomNumber()
	return internal.Equal
}

func getRandomNumber() int64 {
	return rand.Int63n(internal.MaxThreshold-internal.MinThreshold+1) + internal.MinThreshold
}
