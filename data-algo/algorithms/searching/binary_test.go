package searching_test

import (
	"testing"

	"github.com/flan6/data-algo/algorithms/searching"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string]testCase[int]{
		"first item":               {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 0, expectedIndex: 0},
		"last item":                {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 9, expectedIndex: 9},
		"not present":              {input: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, value: 10, expectedIndex: -1},
		"single element":           {input: []int{0}, value: 0, expectedIndex: 0},
		"two elements":             {input: []int{0, 1}, value: 1, expectedIndex: 1},
		"two elements not present": {input: []int{0, 1}, value: 69, expectedIndex: -1},
		"three elements":           {input: []int{0, 1, 2}, value: 1, expectedIndex: 1},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			res := searching.BinarySearch(test.input, test.value)
			if res != test.expectedIndex {
				t.Errorf("unexpected result, got: %d expected: %d", res, test.expectedIndex)
			}
		})
	}
}
