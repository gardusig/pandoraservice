package internal

import "math/rand"

const minThreshold = -4000000000000000000
const maxThreshold = +4000000000000000000

var randomNumber int64

var (
	equal   = "="
	less    = "<"
	greater = ">"
)

func init() {
	shuffleRandomNumber()
}

func validateGuess(guess int64) *string {
	if guess < randomNumber {
		return &less
	}
	if guess > randomNumber {
		return &greater
	}
	shuffleRandomNumber()
	return &equal
}

func shuffleRandomNumber() {
	randomNumber = rand.Int63n(maxThreshold-minThreshold+1) + minThreshold
}
