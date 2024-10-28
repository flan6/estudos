package concurrency

import (
	"sync"
)

// ParallelMergeSort recursively sorts the array in parallel.
func ParallelMergeSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	left := make([]int, mid)
	right := make([]int, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	leftWG := sync.WaitGroup{}
	rightWG := sync.WaitGroup{}

	leftWG.Add(1)
	rightWG.Add(1)

	go ParallelMergeSort(left, &leftWG)
	go ParallelMergeSort(right, &rightWG)

	leftWG.Wait()
	rightWG.Wait()

	copy(arr, merge(left, right))
}

// merge combines two sorted halves into one sorted array.
func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	// Copy the remaining elements of left, if there are any
	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	// Copy the remaining elements of right, if there are any
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}
