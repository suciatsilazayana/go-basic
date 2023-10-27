package utility

import (
	"math/rand"
	"time"
)

func GenerateRandomNumber(min int, max int) (number int) {
	// rand.Seed(time.Now().UnixNano())
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	return generator.Intn(max-min) + min
}
