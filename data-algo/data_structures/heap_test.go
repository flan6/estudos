package data_structures_test

import (
	"container/heap"
	"fmt"
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestMaxHeap(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple":  2,
		"pear":   4,
		"jujuba": 1,
	}

	maxHeap := make(ds.MaxHeap[string], len(items))
	i := 0
	for value, priority := range items {
		maxHeap[i] = ds.Item[string]{
			Value:    value,
			Priority: priority,
		}
		i++
	}
	heap.Init(&maxHeap)

	// Insert a new item.
	heap.Push(&maxHeap, ds.Item[string]{
		Value:    "orange",
		Priority: 5,
	})

	heap.Push(&maxHeap, ds.Item[string]{
		Value:    "banana",
		Priority: 1,
	})

	// Take the items out; they arrive in decreasing priority order.
	for maxHeap.Len() > 0 {
		item := heap.Pop(&maxHeap).(ds.Item[string])

		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}

func TestMinHeap(t *testing.T) {
	items := []int{3, 2, 4, 1}

	minHeap := make(ds.MinHeap[int], 0)
	for index := range items {
		minHeap.Push(items[index])
	}

	// Insert a new item.
	minHeap.Push(5)

	// Take the items out; they arrive in decreasing priority order.
	for len(minHeap) > 0 {
		item := minHeap.Pull()

		fmt.Printf("%d", item)
	}
}
