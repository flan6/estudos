package sorting

func Swap[T any](a, b *T) {
	temp := *a
	*a = *b
	*b = temp
}
