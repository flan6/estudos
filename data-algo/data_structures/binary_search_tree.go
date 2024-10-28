package data_structures

import (
	"cmp"
)

type BinarySearchTree[T cmp.Ordered] struct {
	root *BSTNode[T]
}

type BSTNode[T cmp.Ordered] struct {
	value T
	left  *BSTNode[T]
	right *BSTNode[T]
}

func (n BSTNode[T]) Value() T {
	return n.value
}

func (b *BinarySearchTree[T]) Insert(val T) {
	if b.root == nil {
		b.root = b.root.Insert(val)
		return
	}

	b.root.Insert(val)
}

func (n *BSTNode[T]) Insert(val T) *BSTNode[T] {
	if n == nil {
		return &BSTNode[T]{
			value: val,
		}
	}

	if val <= n.value {
		n.left = n.left.Insert(val)
	} else {
		n.right = n.right.Insert(val)
	}

	return n
}

func (b *BinarySearchTree[T]) Delete(val T) {
	b.root.Delete(val)
}

func (n *BSTNode[T]) Delete(val T) *BSTNode[T] {
	if n == nil {
		return nil
	}

	if val < n.value {
		n.left = n.left.Delete(val)
	} else if val > n.value {
		n.right = n.right.Delete(val)
	} else {
		// Caso tenhamos apenas uma leaf
		if n.left == nil {
			temp := n.right
			n = nil

			return temp
		} else if n.right == nil {
			temp := n.left
			n = nil

			return temp
		}

		// Caso hajam duas leafs
		// Trocaremos pelo sucessor
		temp := n.right.minValue()

		n.value = temp.value

		n.right = n.right.Delete(temp.value)
	}

	return n
}

func (n *BSTNode[T]) minValue() *BSTNode[T] {
	current := n
	for current != nil && current.left != nil {
		current = current.left
	}

	return current
}

// Depth First Search
func (b *BinarySearchTree[T]) Search(val T) (*BSTNode[T], bool) {
	return b.root.Search(val)
}

// Depth First Search
func (n *BSTNode[T]) Search(val T) (*BSTNode[T], bool) {
	if n == nil {
		return nil, false
	}

	if n.value == val {
		return n, true
	}

	if val < n.value {
		return n.left.Search(val)
	} else {
		return n.right.Search(val)
	}
}

// Return tree node values ordered
func (b BinarySearchTree[T]) InOrder() []T {
	items := make([]T, 0)
	items = b.root.InOrder(items)

	return items
}

func (n *BSTNode[T]) InOrder(items []T) []T {
	if n == nil {
		return items
	}

	items = n.left.InOrder(items)
	items = append(items, n.value)
	items = n.right.InOrder(items)

	return items
}

func (b BinarySearchTree[T]) PreOrder() []T {
	items := make([]T, 0)
	items = b.root.PreOrder(items)

	return items
}

func (n *BSTNode[T]) PreOrder(items []T) []T {
	// Base case
	if n == nil {
		return items
	}

	// pre
	items = append(items, n.value)

	// recursion
	items = n.left.PreOrder(items)
	items = n.right.PreOrder(items)

	// post
	return items
}

func (b BinarySearchTree[T]) PostOrder() []T {
	items := make([]T, 0)
	items = b.root.PostOrder(items)

	return items
}

func (n *BSTNode[T]) PostOrder(items []T) []T {
	if n == nil {
		return items
	}

	items = n.left.PostOrder(items)
	items = n.right.PostOrder(items)

	items = append(items, n.value)

	return items
}
