package sorting

import "cmp"

func MergeSort[T ~[]E, E cmp.Ordered](s T) {
	mergeSort(s)
}

func mergeSort[T ~[]E, E cmp.Ordered](s T) {
	if len(s) < 2 {
		return
	}

	mid := len(s) / 2

	left := make(T, mid)
	right := make(T, mid+len(s)%2)

	copy(left, s[:mid])
	copy(right, s[mid:])

	mergeSort(left)
	mergeSort(right)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			s[k] = left[i]
			i++
		} else {
			s[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		s[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		s[k] = right[j]
		j++
		k++
	}
}
