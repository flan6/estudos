package concurrency_test

import (
	"math/rand"
	"testing"

	"github.com/flan6/data-algo/concurrency"
)

func TestParallelQuickSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		ints := make([]int, 5000000)
		for i := range ints {
			ints[i] = rand.Intn(99)
		}

		tests := []Cases[int]{
			SortedCases(ints),
			// SortedCases([]int{9, 0, 10, 99, 1, 88, 6, 5, 3, int(1 << 62)}),
			// SortedCases([]int{1, 2, 0, 22, 6, 5, 3, 9, 10}),
			// SortedCases([]int{111, 11, 1, 2, 1, 1111, 1}),
		}

		HelperParallelSortOrdered(t, concurrency.ParallelQuickSort, tests)
	})
}
