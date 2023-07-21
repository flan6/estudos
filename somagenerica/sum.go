package somagenerica

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer | constraints.Float](a, b any) (T, error) {
	switch a.(type) {
	case int:
		if n, ok := b.(int); ok {
			return T(a.(int) + n), nil
		} else if n, ok := b.(float64); ok {
			return T(float64(a.(int)) + n), nil
		}

		return T(0), errors.New("unsuported type for b")

	case float64:
		if n, ok := b.(int); ok {
			return T(a.(float64) + float64(n)), nil
		} else if n, ok := b.(float64); ok {
			return T(a.(float64) + n), nil
		}

		return T(0), errors.New("unsuported type for b")

	default:
		return T(0), errors.New("unsuported type for a")
	}
}
