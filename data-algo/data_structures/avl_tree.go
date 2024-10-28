package data_structures

import "cmp"

type AVLTree[T cmp.Ordered] struct {
	root *AVLNode[T]
}

type AVLNode[T cmp.Ordered] struct {
	value  T
	height int
	left   *AVLNode[T]
	right  *AVLNode[T]
}

func (t *AVLTree[T]) Insert(value T) {
	if t.root == nil {
		t.root = t.root.insert(value)
		return
	}

	t.root.insert(value)
}

func (t *AVLTree[T]) Delete(value T) {
	t.root = t.root.delete(value)
}

func (t AVLTree[T]) InOrder() []T {
	items := make([]T, 0)
	items = t.root.inOrder(items)

	return items
}

func (t AVLTree[T]) PreOrder() []T {
	items := make([]T, 0)
	items = t.root.preOrder(items)

	return items
}

func (t AVLTree[T]) PostOrder() []T {
	items := make([]T, 0)
	items = t.root.postOrder(items)

	return items
}

// Recursive function to insert a node and rebalance tree
func (n *AVLNode[T]) insert(value T) *AVLNode[T] {
	if n == nil {
		return &AVLNode[T]{value: value, height: 1}
	}

	if value < n.value {
		n.left = n.left.insert(value)
	} else if value > n.value {
		n.right = n.right.insert(value)
	} else {
		return n // Duplicate values are not allowed
	}

	// Update the balance factor of this ancestor node
	n.updateHeight()
	balance := n.getBalance()

	// left left
	if balance > 1 && value < n.left.value {
		return rotateRight(n)
	}

	// left right
	if balance > 1 && value > n.left.value {
		n.left = rotateLeft(n.left)
		return rotateRight(n)
	}

	// right right
	if balance < -1 && value > n.right.value {
		return rotateLeft(n)
	}

	// right left
	if balance < -1 && value < n.right.value {
		n.right = rotateRight(n.right)
		return rotateLeft(n)
	}

	return n
}

func (n *AVLNode[T]) delete(value T) *AVLNode[T] {
	if n == nil {
		return nil
	}

	if value < n.value {
		n.left = n.left.delete(value)
	} else if value > n.value {
		n.right = n.right.delete(value)
	} else {
		// Caso tenhamos apenas uma leaf
		if n.left == nil {
			n = n.right
			n.right = nil
		} else if n.right == nil {
			n = n.left
			n.left = nil
		} else {
			// Caso hajam duas leafs
			// Trocaremos pelo sucessor
			temp := n.right.minValue()

			n.value = temp.value

			n.right = n.right.delete(temp.value)
		}
	}

	if n == nil {
		return nil
	}

	n.updateHeight()

	balance := n.getBalance()

	// If this node becomes unbalanced, then there are 4 cases

	// Left Left Case
	if balance > 1 && n.left.getBalance() >= 0 {
		return rotateRight(n)
	}

	// Left Right Case
	if balance > 1 && n.left.getBalance() < 0 {
		n.left = rotateLeft(n.left)
		return rotateRight(n)
	}

	// Right Right Case
	if balance < -1 && n.right.getBalance() <= 0 {
		return rotateLeft(n)
	}

	// Right Left Case
	if balance < -1 && n.right.getBalance() > 0 {
		n.right = rotateRight(n.right)
		return rotateLeft(n)
	}

	return n
}

func (n *AVLNode[T]) inOrder(items []T) []T {
	if n == nil {
		return items
	}

	items = n.left.inOrder(items)
	items = append(items, n.value)
	items = n.right.inOrder(items)

	return items
}

func (n *AVLNode[T]) preOrder(items []T) []T {
	if n == nil {
		return items
	}

	items = append(items, n.value)
	items = n.left.preOrder(items)
	items = n.right.preOrder(items)

	return items
}

func (n *AVLNode[T]) postOrder(items []T) []T {
	if n == nil {
		return items
	}

	items = n.left.postOrder(items)
	items = n.right.postOrder(items)
	items = append(items, n.value)

	return items
}

// Get balance factor of node n
func (n *AVLNode[T]) getBalance() int {
	if n == nil {
		return 0
	}

	return n.left.getHeight() - n.right.getHeight()
}

// Helper function to get the height of the node
func (n *AVLNode[T]) getHeight() int {
	if n == nil {
		return 0
	}

	return n.height
}

// Helper function to update the height of the node
func (n *AVLNode[T]) updateHeight() {
	leftHeight := n.left.getHeight()
	rightHeight := n.right.getHeight()

	if leftHeight > rightHeight {
		n.height = leftHeight + 1
	} else {
		n.height = rightHeight + 1
	}
}

func (n *AVLNode[T]) minValue() *AVLNode[T] {
	current := n
	for current != nil && current.left != nil {
		current = current.left
	}

	return current
}

// Right rotation to balance the tree
func rotateRight[T cmp.Ordered](y *AVLNode[T]) *AVLNode[T] {
	x := y.left
	T2 := x.right

	// Perform rotation
	x.right = y
	y.left = T2

	// Update heights
	y.updateHeight()
	x.updateHeight()

	return x
}

// Left rotation to balance the tree
func rotateLeft[T cmp.Ordered](x *AVLNode[T]) *AVLNode[T] {
	y := x.right
	T2 := y.left

	// Perform rotation
	y.left = x
	x.right = T2

	// Update heights
	x.updateHeight()
	y.updateHeight()

	return y
}
