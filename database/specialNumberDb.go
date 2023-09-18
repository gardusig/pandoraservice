package database

import (
	"fmt"
	"math/rand"

	"github.com/gardusig/pandoraservice/internal"
)

const (
	defaultEncryptedMessage = "someEncryptedMessage"
	defaultDecryptedMessage = "someDecryptedMessage"
)

type SpecialNumberDb struct {
	randomNumberByLevel     map[uint32]int64
	encryptedMessageByLevel map[uint32]string
	decryptedMessageByLevel map[uint32]string
}

func NewSpecialNumberDb() *SpecialNumberDb {
	return &SpecialNumberDb{
		randomNumberByLevel:     getPopulatedRandomNumberByLevel(),
		encryptedMessageByLevel: getPopulatedEncryptedMessageByLevel(),
		decryptedMessageByLevel: getPopulatedDecryptedMessageByLevel(),
	}
}

func (db *SpecialNumberDb) ValidateGuess(level uint32, guess int64) (string, *string, error) {
	randomNumber := db.randomNumberByLevel[level]
	if guess < randomNumber {
		return internal.Less, nil, nil
	}
	if guess > randomNumber {
		return internal.Greater, nil, nil
	}
	db.randomNumberByLevel[level] = getRandomNumber()
	encryptedMessage := db.encryptedMessageByLevel[level]
	return internal.Equal, &encryptedMessage, nil
}

func (db *SpecialNumberDb) ValidateLockedPandoraBox(level uint32, encryptedMessage string) (string, error) {
	if encryptedMessage != db.encryptedMessageByLevel[level] {
		return "", fmt.Errorf("wrong encrypted message")
	}
	return db.decryptedMessageByLevel[level], nil
}

func getRandomNumber() int64 {
	return rand.Int63n(internal.GuessMaxThreshold-internal.GuessMinThreshold+1) + internal.GuessMinThreshold
}

func getPopulatedRandomNumberByLevel() map[uint32]int64 {
	db := make(map[uint32]int64)
	for level := internal.LevelMinThreshold; level <= internal.LevelMaxThreshold; level += 1 {
		db[level] = getRandomNumber()
	}
	return db
}

func getPopulatedEncryptedMessageByLevel() map[uint32]string {
	db := make(map[uint32]string)
	for level := internal.LevelMinThreshold; level <= internal.LevelMaxThreshold; level += 1 {
		db[level] = defaultEncryptedMessage
	}
	return db
}

func getPopulatedDecryptedMessageByLevel() map[uint32]string {
	db := make(map[uint32]string)
	for level := internal.LevelMinThreshold; level <= internal.LevelMaxThreshold; level += 1 {
		db[level] = defaultDecryptedMessage
	}
	return db
}
