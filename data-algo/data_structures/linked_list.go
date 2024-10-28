package data_structures

import "errors"

type List[T any] interface {
	Length() int
	InsertAt(index int, item T)
	Remove(index int) T
	Append(item T)
	Prepend(item T)
	Get(index int) T
}

type DoubleLinkedList[T any] struct {
	head *nodeDList[T]
	tail *nodeDList[T]
	len  uint
}

type nodeDList[T any] struct {
	value T
	next  *nodeDList[T]
	prev  *nodeDList[T]
}

var ErrIndexOutOfRange = errors.New("index out of range")

func NewLinkedList[T any]() DoubleLinkedList[T] {
	return DoubleLinkedList[T]{}
}

// Length returns the current size of the list
func (l DoubleLinkedList[T]) Length() uint {
	return l.len
}

// Append inserts item at the end of the list
func (l *DoubleLinkedList[T]) Append(item T) {
	n := &nodeDList[T]{value: item}
	l.len++

	if l.tail == nil {
		l.head = n
		l.tail = n
		return
	}

	n.prev = l.tail
	l.tail.next = n
	l.tail = n
}

// Prepend inserts item in the begining of the list
func (l *DoubleLinkedList[T]) Prepend(item T) {
	n := &nodeDList[T]{value: item}
	l.len++

	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	n.next = l.head
	l.head.prev = n
	l.head = n
}

// Get gets the value at index
// If list is empty, zero value is returned and nil error
// Returns ErrIndexOutOfRange in case index does not exist in the list
func (l DoubleLinkedList[T]) Get(index uint) (T, error) {
	if index > l.len {
		return *new(T), ErrIndexOutOfRange
	}

	if index == 0 && l.len == 0 {
		return *new(T), nil // TODO: should return error empty list?
	}

	return l.get(index).value, nil
}

// InsertAt inserts item at index
// Returns ErrIndexOutOfRange in case index does not exist in the list
func (l *DoubleLinkedList[T]) InsertAt(index uint, item T) error {
	if index > l.len {
		return ErrIndexOutOfRange
	}

	if index == l.len {
		l.Append(item)
		return nil
	}

	if index == 0 {
		l.Prepend(item)
		return nil
	}

	curr := l.get(index)

	node := &nodeDList[T]{
		value: item,
		next:  curr,
		prev:  curr.prev,
	}
	curr.prev = node
	node.prev.next = node

	l.len++

	return nil
}

// Set sets the new value at index
// returns ErrIndexOutOfRange in case index does not exist in the list
func (l *DoubleLinkedList[T]) Set(index uint, value T) error {
	if index > l.len {
		return ErrIndexOutOfRange
	}

	if index == 0 && l.len == 0 {
		l.Append(value)
	}

	l.get(index).value = value

	return nil
}

// Remove removes the node at index
// returns ErrIndexOutOfRange in case index does not exist in the list
func (l *DoubleLinkedList[T]) Remove(index uint) error {
	if index > l.len {
		return ErrIndexOutOfRange
	}

	if l.head == nil {
		return nil
	}

	curr := l.get(index)

	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}

	if l.head == curr {
		l.head = curr.next
	}

	if l.tail == curr {
		l.tail = curr.prev
	}

	curr.prev = nil
	curr.next = nil
	l.len--

	return nil
}

func (l DoubleLinkedList[T]) get(index uint) *nodeDList[T] {
	if index == l.len {
		return l.tail
	}

	curr := l.head
	for i := uint(0); i < index && curr != nil; i++ {
		curr = curr.next
	}

	return curr
}

type LinkedListEval[T comparable] struct {
	DoubleLinkedList[T]
}

func NewLinkedListEval[T comparable]() LinkedListEval[T] {
	return LinkedListEval[T]{}
}

// RemoveValue removes the first ocurrence of value
// No op if not found
func (l *LinkedListEval[T]) RemoveValue(value T) {
	// TODO: travels two times the list. must proper implement single traversal later
	idx := l.IndexOf(value)
	if idx < 0 {
		return
	}

	l.Remove(uint(idx))
}

// IndexOf finds the index of value, returns -1 if not found
func (l LinkedListEval[T]) IndexOf(value T) int {
	idx := 0
	n := l.walk(l.head, value, &idx)
	if n != nil {
		return idx
	}

	return -1
}

func (l LinkedListEval[T]) walk(n *nodeDList[T], value T, index *int) *nodeDList[T] {
	if n == nil {
		return nil
	}

	if n.value == value {
		return n
	}

	if index != nil {
		*index++
	}

	return l.walk(n.next, value, index)
}
