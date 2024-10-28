package data_structures_test

import (
	"math/rand"
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestLinkedList(t *testing.T) {
	// Setup
	list := ds.NewLinkedListEval[int]()

	val := []int{0, -10, 69, 420}
	for i := range val {
		list.Append(val[i])
	}

	// Length
	expectedSize := len(val)
	listSize := list.Length()
	if expectedSize != int(listSize) {
		t.Fatalf("Length: wrong size for the list, expected %d got %d", expectedSize, listSize)
	}

	// Get and Append from setup
	for i := range val {
		res, err := list.Get(uint(i))
		if res != val[i] || err != nil {
			t.Fatalf("Append: got %d expected %d at index %d", res, val[i], i)
		}
	}

	// Prepend
	v := 69420
	list.Prepend(v)
	if res, err := list.Get(0); res != v || err != nil {
		t.Fatalf("Prepend: got %d but expected %d at head", res, v)
	}

	// IndexOf
	idx := list.IndexOf(v)
	if idx != 0 {
		t.Fatalf("IndexOf: got index %d but expected 0 for value %d", idx, v)
	}

	// Set
	v = 123
	err := list.Set(0, v)
	if err != nil {
		t.Fatalf("Set: coult not set %d at index 0, got err %s", v, err)
	}
	if res, err := list.Get(0); res != v || err != nil {
		t.Fatalf("Set: got %d but expected %d at head", res, v)
	}
}

func FuzzLinkedListEval_RemoveValue(f *testing.F) {
	f.Add("test", "fuzz", "random")
	f.Add("tset", "flan", "modnar")

	f.Fuzz(func(t *testing.T, a, b, c string) {
		list := ds.NewLinkedListEval[string]()

		values := []string{a, b, c}
		vOccurrences := make(map[string]int)
		for i := range values {
			vOccurrences[values[i]]++
			list.Append(values[i])

			// Assures that append is working
			res, err := list.Get(uint(i))
			if err != nil || res != values[i] {
				t.Errorf("%q %v", res, err)
			}
		}

		vToRemove := values[rand.Intn(2)]
		list.RemoveValue(vToRemove)

		lOccurrences := make(map[string]int)
		for i := uint(0); i < list.Length(); i++ {
			res, err := list.Get(i)
			if err != nil {
				t.Errorf("failed to get %v: %v", res, err)
			}

			lOccurrences[res]++
		}

		if vOccurrences[vToRemove]-1 != lOccurrences[vToRemove] {
			t.Errorf("failed to remove %v", vToRemove)
		}
	})
}

func TestLinkedListEval_RemoveValue(t *testing.T) {
	tests := map[string]struct {
		items    []string
		toRemove string
		expected []string
	}{
		"remove last item": {
			items:    []string{"fizz", "buzz", "blaus"},
			toRemove: "blaus",
			expected: []string{"fizz", "buzz"},
		},

		"item does not exist": {
			items:    []string{"fizz", "buzz", "blaus"},
			toRemove: "flan",
			expected: []string{"fizz", "buzz", "blaus"},
		},

		"item exist more than once, should exclude first": {
			items:    []string{"flan", "flan", "blaus"},
			toRemove: "flan",
			expected: []string{"flan", "blaus"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			list := ds.NewLinkedListEval[string]()
			for i, item := range test.items {
				list.Append(item)

				// Assures append is working
				res, err := list.Get(uint(i))
				if err != nil || res != item {
					t.Errorf("failed to append item %s got %s: %v", item, res, err)
				}
			}

			list.RemoveValue(test.toRemove)

			for i := uint(0); i < list.Length(); i++ {
				res, err := list.Get(i)
				if err != nil {
					t.Fatal(err)
				}

				if res != test.expected[i] {
					t.Fatalf("unexpected item found %s on list index %d, expected:%v", res, i, test.expected)
				}
			}
		})
	}
}
