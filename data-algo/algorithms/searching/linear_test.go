package searching_test

import (
	"testing"

	"github.com/flan6/data-algo/algorithms/searching"
)

type testCase[T comparable] struct {
	input         []T
	value         T
	expectedIndex int
}

func TestLinearSearch(t *testing.T) {
	tests := map[string]testCase[int]{
		"ordered first item": {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 0, expectedIndex: 0},
		"ordered last item":  {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 9, expectedIndex: 9},
		"not present":        {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 10, expectedIndex: -1},
		"unordered":          {input: []int{28, 30, 1, 69, 420, 24, 32, 44, 31}, value: 69, expectedIndex: 3},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := searching.LinearSearch(test.input, test.value)
			if res != test.expectedIndex {
				t.Errorf("unexpected result, got: %d expected: %d", res, test.expectedIndex)
			}
		})
	}
}
