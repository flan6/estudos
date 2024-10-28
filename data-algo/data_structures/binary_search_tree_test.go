package data_structures_test

import (
	"fmt"
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestInsert(t *testing.T) {
	tree := ds.BinarySearchTree[int]{}

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(0)
	tree.Insert(-1)

	InOrderRes := (tree.InOrder())
	fmt.Println(InOrderRes)
	PreOrderRes := (tree.PreOrder())
	fmt.Println(PreOrderRes)
	PostOrderRes := (tree.PostOrder())
	fmt.Println(PostOrderRes)
}

func TestDelete(t *testing.T) {
	tree := ds.BinarySearchTree[int]{}

	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(0)
	tree.Insert(-1)

	before := tree.PreOrder()
	fmt.Println(before)
	tree.Delete(3)
	after := (tree.PreOrder())
	fmt.Println(after)
}

func TestSearch(t *testing.T) {
	tree := ds.BinarySearchTree[rune]{}

	tree.Insert('A')
	tree.Insert('B')
	tree.Insert('C')
	tree.Insert('D')
	tree.Insert('E')

	node, found := tree.Search('C')
	if !found {
		t.Fatalf("not found item %v in tree", 'C')
	}

	if node.Value() != 'C' {
		t.Errorf("wrong value from node. got: %v expected: C", node.Value())
	}
}
