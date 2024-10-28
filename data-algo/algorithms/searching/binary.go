package searching

import "cmp"

// Input slice must have elements sorted from lower to greater
func BinarySearch[T ~[]E, E cmp.Ordered](input T, value E) int {
	low, high := 0, len(input)

	for {
		middle := low + (high-low)/2

		searchValue := input[middle]
		if value == searchValue {
			return middle
		} else if value < searchValue {
			high = middle
		} else {
			low = middle + 1
		}

		if low >= high {
			break
		}
	}

	return -1
}
