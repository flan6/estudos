package main

import (
	"errors"
	"log"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Sum[T Number](a, b any) (T, error) {
	var (
		result T
	)

	n, ok := a.(T)
	if !ok {
		return result, errors.New("a invalid type")
	}

	m, ok := b.(T)
	if !ok {
		return result, errors.New("b invalid type")
	}

	result = n + m

	return result, nil
}

func main() {
	log.Println(Sum[int](1, 2.68))
	log.Println(Sum[int]("1", 2.68))
	log.Println(Sum[int](1.2, float32(2.68)))
	log.Println(Sum[int](uint8(1), 2.68))
}
