package database

import (
	"math/rand"

	"github.com/gardusig/grpc_service/internal"
)

type SpecialNumberDb struct {
	randomNumberByLevel map[uint32]int64
}

func NewSpecialNumberDb() *SpecialNumberDb {
	return &SpecialNumberDb{
		randomNumberByLevel: getPopulatedDb(),
	}
}

func (db *SpecialNumberDb) ValidateGuess(level uint32, guess int64) string {
	randomNumber := db.randomNumberByLevel[level]
	if guess < randomNumber {
		return internal.Less
	}
	if guess > randomNumber {
		return internal.Greater
	}
	db.randomNumberByLevel[level] = getRandomNumber()
	return internal.Equal
}

func getRandomNumber() int64 {
	return rand.Int63n(internal.GuessMaxThreshold-internal.GuessMinThreshold+1) + internal.GuessMinThreshold
}

func getPopulatedDb() map[uint32]int64 {
	db := make(map[uint32]int64)
	for level := internal.LevelMinThreshold; level <= internal.LevelMaxThreshold; level += 1 {
		db[level] = getRandomNumber()
	}
	return db
}
