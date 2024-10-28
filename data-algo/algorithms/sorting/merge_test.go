package sorting_test

import (
	"math/rand"
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func TestMergeSort(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		ints := make([]int, 21)
		for i := range ints {
			ints[i] = rand.Intn(99)
		}

		tests := []Cases[int]{
			SortedCases(ints),
			SortedCases([]int{9, 0, 10, 99, 1, 88, 6, 5, 3, int(1 << 62)}),
			SortedCases([]int{1, 2, 0, 22, 6, 5, 3, 9, 10}),
			SortedCases([]int{111, 11, 1, 2, 1, 1111, 1}),
		}

		HelperSortOrdered[[]int](t, sorting.MergeSort, tests)
	})

	t.Run("strings", func(t *testing.T) {
		tests := []Cases[string]{
			SortedCases([]string{"z", "x", "c", "v", "b", "n"}),
			SortedCases([]string{"zebra", "abobora", "urso", "abelha", "limao", "batata"}),
		}

		HelperSortOrdered[[]string](t, sorting.MergeSort, tests)
	})

	t.Run("floats", func(t *testing.T) {
		tests := []Cases[float32]{
			SortedCases([]float32{1.123, 1.321, 2.0, 0.3, 0}),
		}

		HelperSortOrdered[[]float32](t, sorting.MergeSort, tests)
	})
}
