package sorting

import (
	"cmp"
)

func BubbleSort[E ~[]T, T cmp.Ordered](s E) {
	for i := 0; i < len(s)-1; i++ {
		swapped := false

		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(&s[j], &s[j+1])
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}

func BubbleSortVisualization[E ~[]T, T cmp.Ordered](s E, f func()) {
	for i := 0; i < len(s)-1; i++ {
		swapped := false

		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(&s[j], &s[j+1])
				swapped = true
			}
			f()
		}

		if !swapped {
			break
		}
	}
}
