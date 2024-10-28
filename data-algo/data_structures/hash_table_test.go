package data_structures_test

import (
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestHashTable_Insert(t *testing.T) {
	table := ds.NewHashTable(5)

	table.Insert(1)
	table.Insert(2)
	table.Insert(2)
	table.Insert(3)

	key := 2
	if found := table.Search(key); !found {
		t.Errorf("key %d not found in table", key)
	}
}

func TestHashTable_Delete(t *testing.T) {
	type cases struct {
		tableValues []int
		del         int
	}

	tests := map[string]cases{
		"first item on index": {[]int{1, 2, 3}, 3},
		"last item on index":  {[]int{1, 2, 3}, 1},
		"item does not exist": {[]int{1, 2, 3}, 4},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			table := ds.NewHashTable(len(test.tableValues) * 2)

			for v := range test.tableValues {
				table.Insert(test.tableValues[v])
			}

			table.Delete(test.del)
			if found := table.Search(test.del); found {
				t.Errorf("test.del %d found in table", test.del)
			}
		})
	}
}
