package algorithms

import "math"

func TwoCrystalBalls(input []bool) int {
 	jump := int(math.Sqrt(float64(len(input))))

	i := 0

	for ; i < len(input); i += jump {
		if input[i] {
			break
		}
	}

	i -= jump

	for j := i; j < i+jump && j < len(input); j++ {
		if input[j] {
			return j
		}
	}

	return -1
}
