package concurrency

import (
	"sync"
)

func ParallelQuickSort(values []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(values) <= 1 {
		return
	}

	pivotIndex := partition(values)

	wg.Add(2)
	go ParallelQuickSort(values[:pivotIndex], wg)
	go ParallelQuickSort(values[pivotIndex+1:], wg)
}

func partition(values []int) int {
	high := len(values) - 1
	pivot := values[high]

	i := -1
	for j := 0; j < high; j++ {
		if values[j] <= pivot {
			i++
			// swap
			values[i], values[j] = values[j], values[i]
		}
	}

	i++
	// swap
	values[i], values[high] = values[high], values[i]

	return i
}
