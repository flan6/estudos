package concurrency_test

import (
	"cmp"
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"testing"

	"github.com/flan6/data-algo/concurrency"
)

func TestParallelMergeSort(t *testing.T) {
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

		HelperParallelSortOrdered(t, concurrency.ParallelMergeSort, tests)
	})
}

type Cases[T cmp.Ordered] struct {
	input    []T
	expected []T
}

func SortedCases[E []T, T cmp.Ordered](input E) Cases[T] {
	expec := make(E, len(input))
	copy(expec, input)
	slices.Sort(expec)
	return Cases[T]{input: input, expected: expec}
}

func HelperParallelSortOrdered[E ~[]T, T cmp.Ordered](t *testing.T, orderFunction func(s E, wg *sync.WaitGroup), c []Cases[T]) {
	for name, test := range c {
		t.Run(fmt.Sprintf("case: %d", name), func(t *testing.T) {
			t.Parallel()

			wg := new(sync.WaitGroup)

			wg.Add(1)
			go orderFunction(test.input, wg)
			wg.Wait()

			for i := range test.input {
				if test.input[i] != test.expected[i] {
					t.Fatalf("input not sorted, got: %v expected: %v", test.input, test.expected)
				}
			}
		})
	}
}
