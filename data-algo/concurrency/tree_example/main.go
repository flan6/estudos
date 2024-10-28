package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree)

	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}

		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}

	walk(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	first := make(chan int, 10)
	second := make(chan int, 10)

	go Walk(t1, first)
	go Walk(t2, second)

	for {
		v1, ok1 := <-first
		v2, ok2 := <-second
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return true
}

func main() {
	values := make(chan int)

	go Walk(tree.New(1), values)

	for val := range values {
		fmt.Printf("%d ", val)
	}
	fmt.Println("")

	fmt.Scan()
	fmt.Println("should be true:", Same(tree.New(1), tree.New(1)))
	fmt.Println("should be false:", Same(tree.New(1), tree.New(2)))
}
