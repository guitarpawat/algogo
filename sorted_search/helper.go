package sorted_search

import (
	"testing"
)

func createSlice(size int) []int {
	numbers := make([]int, size)
	for i:=0; i<size; i++ {
		numbers[i] = i
	}
	return numbers
}

func doSearch(numbers []int, f func(int, []int) bool) {
	size := len(numbers)
	for i:=0; i<size; i++ {
		f(i, numbers)
	}
}

func benchFunc(b *testing.B, size int, f func(int, []int) bool) {
	b.StopTimer()
	numbers := createSlice(size)
	b.StartTimer()
	for i:=0; i<b.N; i++ {
		doSearch(numbers, f)
	}
}