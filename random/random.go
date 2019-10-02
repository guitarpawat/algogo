package random

import (
	"github.com/guitarpawat/algogo/util"
	"math/rand"
)

// RandDistinct generates slice of uniquely random from 1 to limit
// by implementing Fisher-Yates algorithm.
func RandDistinct(limit int) []int {
	// Create slice from 1 to limit
	slice := make([]int, limit, limit)
	for i := 0; i < limit; i++ {
		slice[i] = i + 1
	}
	// Implements Fisher-Yates algorithm
	for i := limit - 1; i > 0; i-- {
		target := rand.Intn(i + 1)
		util.SwapInt(slice, i, target)
	}

	return slice
}
