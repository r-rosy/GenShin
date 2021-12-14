package cardpool

import (
	"math/rand"
	"time"
)

func Random() int {
	r := generateRandnum()
	if r <= 6 {
		return 5
	}
	if r <= 51 {
		return 4
	}
	return 3
}
func generateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1000)
	return randNum
}
