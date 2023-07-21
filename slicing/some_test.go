package main

import (
	"testing"
)

func benchmarkSlicingWithIf(someSlice []int, count int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SlicingWithIf(someSlice, count)
	}
}

func benchmarkSlicingWithMin(someSlice []int, count int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		SlicingWithMin(someSlice, count)
	}
}

func BenchmarkSlicingWithIfShort(b *testing.B) {
	someSlice := make([]int, 10)
	count := 5
	benchmarkSlicingWithIf(someSlice, count, b)
}

func BenchmarkSlicingWithIfLong(b *testing.B) {
	someSlice := make([]int, 1000)
	count := 500
	benchmarkSlicingWithIf(someSlice, count, b)
}

func BenchmarkSlicingWithMinShort(b *testing.B) {
	someSlice := make([]int, 10)
	count := 5
	benchmarkSlicingWithMin(someSlice, count, b)
}

func BenchmarkSlicingWithMinLong(b *testing.B) {
	someSlice := make([]int, 1000)
	count := 500
	benchmarkSlicingWithMin(someSlice, count, b)
}
