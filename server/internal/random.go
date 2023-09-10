package internal

import (
	"math/rand"
)

var randomNumber int64

func shuffleRandomNumber() {
	randomNumber = rand.Int63n(maxThreshold-minThreshold+1) + minThreshold
}
