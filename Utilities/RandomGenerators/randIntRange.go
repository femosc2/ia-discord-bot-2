package utilities

import (
	"math/rand"
	"time"
)

// RandIntFromRange returns a true random int between 0 and param max
func RandIntFromRange(max int) int {
	rand.Seed(time.Now().UnixNano()) // Ensures proper random number generation
	return rand.Intn(max)
}
