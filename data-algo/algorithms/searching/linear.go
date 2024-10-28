package searching

func LinearSearch[T ~[]E, E comparable](input T, value E) int {
	for i, val := range input {
		if val == value {
			return i
		}
	}

	return -1
}
