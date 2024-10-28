package sorting

import "cmp"

func QuickSort[T ~[]E, E cmp.Ordered](s T) {
	quickSort(s, 0, len(s)-1)
}

func quickSort[T ~[]E, E cmp.Ordered](s T, low, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(s, low, high)

	quickSort(s, low, pivotIdx-1)
	quickSort(s, pivotIdx+1, high)
}

// Picks the pivot and moves every item bigger than P
// to right and lower or equal to left
func partition[T ~[]E, E cmp.Ordered](s T, low, high int) int {
	pivot := s[high]

	idx := low - 1
	for i := low; i < high; i++ {
		if s[i] <= pivot {
			idx++
			Swap(&s[i], &s[idx])
		}
	}

	idx++
	Swap(&s[high], &s[idx])

	return idx
}
