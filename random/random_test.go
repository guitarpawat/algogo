package random

import (
	"github.com/ethereum/go-ethereum/common/math"
	"testing"
)

func isDistinct(slice []int) bool {
	m := map[int]struct{}{}
	for _, v := range slice {
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}

func isInBound(slice []int) bool {
	limit := len(slice)
	for _, v := range slice {
		if v < 1 || v > limit {
			return false
		}
	}
	return true
}

func TestRandDistinct(t *testing.T) {
	// Birthday problem: P(false) is nearly 1
	slice := RandDistinct(math.MaxInt16)
	ok := isDistinct(slice)
	if !ok {
		t.Error("the slice has duplicate number!")
	}
	ok = isInBound(slice)
	if !ok {
		t.Error("the slice is out of bound limit!")
	}
}

func BenchmarkRandDistinct10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(10)
	}
}

func BenchmarkRandDistinct100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(100)
	}
}

func BenchmarkRandDistinct1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(1000)
	}
}

func BenchmarkRandDistinct10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(10000)
	}
}

func BenchmarkRandDistinct100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(100000)
	}
}

func BenchmarkRandDistinct1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandDistinct(1000000)
	}
}
