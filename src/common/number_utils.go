package common

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber
}
