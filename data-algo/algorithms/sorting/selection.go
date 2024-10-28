package sorting

import "cmp"

func SelectionSort[T ~[]E, E cmp.Ordered](s T) {
	for i := 0; i < len(s)-1; i++ {
		min := i

		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}

		if min != i {
			Swap(&s[min], &s[i])
		}
	}
}

func SelectionSortVisual[T ~[]E, E cmp.Ordered](s T, f func()) {
	for i := 0; i < len(s)-1; i++ {
		min := i

		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
			f()
		}

		if min != i {
			Swap(&s[min], &s[i])
		}
	}
}
