package data_structures

import "cmp"

// MaxHeap is a max-heap.
// Implements heap.Interface
type MaxHeap[T any] []Item[T]

type Item[T any] struct {
	Value    T
	Priority int
}

func (h MaxHeap[T]) Less(i, j int) bool {
	return h[i].Priority < h[j].Priority
}

func (h MaxHeap[T]) Len() int {
	return len(h)
}

func (h MaxHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap[T]) Push(x any) {
	*h = append(*h, x.(Item[T]))
}

func (h *MaxHeap[T]) Pop() any {
	old := *h

	n := len(old)

	item := old[n-1]

	*h = old[0 : n-1]

	return item
}

type MinHeap[T cmp.Ordered] []T

func (m *MinHeap[T]) Push(val T) {
	*m = append(*m, val)
	m.heapifyUp(len(*m) - 1)
}

func (m *MinHeap[T]) Pull() T {
	if len(*m) == 0 {
		return *new(T)
	}

	heap := *m

	out := heap[0]
	if len(*m) == 1 {
		*m = make(MinHeap[T], 0)
		return out
	}

	heap[0] = heap[len(heap)-1]
	*m = heap
	m.heapifyDown(0)
	*m = heap[0 : len(heap)-1]
	return out
}

func (m MinHeap[T]) heapifyDown(index int) {
	if index >= len(m) {
		return
	}

	leftIndex := leftChild(index)
	rightIndex := rightChild(index)

	if leftIndex >= len(m) {
		return
	}

	// TODO: right child panics

	ourValue := m[index]
	leftValue, rightValue := m[leftIndex], m[rightIndex]

	if leftValue > rightValue && ourValue > rightValue {
		m[index] = rightValue
		m[rightIndex] = ourValue
		m.heapifyDown(rightIndex)
	} else if rightValue > leftValue && ourValue > leftValue {
		m[index] = leftValue
		m[leftIndex] = ourValue
		m.heapifyDown(leftIndex)
	}
}

func (m MinHeap[T]) heapifyUp(index int) {
	if index == 0 {
		return
	}

	parentIndex := parent(index)
	parentValue := m[parentIndex]
	ourValue := m[index]

	if parentValue > ourValue {
		m[index] = parentValue
		m[parentIndex] = ourValue
		m.heapifyUp(parentIndex)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}
