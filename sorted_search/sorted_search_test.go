package sorted_search

import "testing"

func BenchmarkLinearSearch1(b *testing.B) {
	benchFunc(b, 1, LinearSearch)
}

func BenchmarkLinearSearch10(b *testing.B) {
	benchFunc(b, 10, LinearSearch)
}

func BenchmarkLinearSearch100(b *testing.B) {
	benchFunc(b, 100, LinearSearch)
}

func BenchmarkLinearSearch1000(b *testing.B) {
	benchFunc(b, 1000, LinearSearch)
}

func BenchmarkLinearSearch10000(b *testing.B) {
	benchFunc(b, 10000, LinearSearch)
}

func BenchmarkLinearSearch100000(b *testing.B) {
	benchFunc(b, 100000, LinearSearch)
}

func BenchmarkBinarySearch1(b *testing.B) {
	benchFunc(b, 1, BinarySearch)
}

func BenchmarkBinarySearch10(b *testing.B) {
	benchFunc(b, 10, BinarySearch)
}

func BenchmarkBinarySearch100(b *testing.B) {
	benchFunc(b, 100, BinarySearch)
}

func BenchmarkBinarySearch1000(b *testing.B) {
	benchFunc(b, 1000, BinarySearch)
}

func BenchmarkBinarySearch10000(b *testing.B) {
	benchFunc(b, 10000, BinarySearch)
}

func BenchmarkBinarySearch100000(b *testing.B) {
	benchFunc(b, 100000, BinarySearch)
}

func BenchmarkBinarySearch1000000(b *testing.B) {
	benchFunc(b, 1000000, BinarySearch)
}
