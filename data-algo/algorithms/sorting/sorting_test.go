package sorting_test

import (
	"cmp"
	"fmt"
	"slices"
	"testing"

	"github.com/flan6/data-algo/algorithms/sorting"
)

func TestSwap(t *testing.T) {
	a, b := 2, 1

	sorting.Swap(&a, &b)

	if a != 1 || b != 2 {
		t.Errorf("failed to swap elements, 'a' expected: 1 got: %d, 'b' expected: 2 got: %d", a, b)
	}
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

func HelperSortOrdered[E ~[]T, T cmp.Ordered](t *testing.T, f func(s E), c []Cases[T]) {
	for name, test := range c {
		t.Run(fmt.Sprintf("case: %d", name), func(t *testing.T) {
			t.Parallel()

			f(test.input)

			for i := range test.input {
				if test.input[i] != test.expected[i] {
					t.Fatalf("input not sorted, got: %v expected: %v", test.input, test.expected)
				}
			}
		})
	}
}
