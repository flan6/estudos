package data_structures_test

import (
	"fmt"
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestAVLTree_Insert(t *testing.T) {
	tree := ds.AVLTree[int]{}

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(0)
	tree.Insert(-1)

	InOrderRes := (tree.InOrder())
	fmt.Println(InOrderRes)
}

func TestAVLTree_Delete(t *testing.T) {
	tree := ds.AVLTree[int]{}

	tree.Insert(9)
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(0)
	tree.Insert(6)
	tree.Insert(11)
	tree.Insert(-1)
	tree.Insert(1)
	tree.Insert(2)

	before := tree.PreOrder()
	fmt.Println(before)
	tree.Delete(10)
	after := (tree.PreOrder())
	fmt.Println(after)
}
