package data_structures

type Queue[T any] struct {
	data DoubleLinkedList[T]
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{
		NewLinkedList[T](),
	}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data.Prepend(value)
}

func (q *Queue[T]) Dequeue() T {
	res, _ := q.data.Get(q.data.Length() - 1)

	q.data.Remove(q.data.Length())

	return res
}
