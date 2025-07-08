package queue

import "github.com/thefrost13/gollections/node"

type Queue[T any] struct {
	first *node.LinkedNode[T]
	last  *node.LinkedNode[T]
	size  int
}

func NewQueue[T any](sli []T) *Queue[T] {
	if sli == nil {
		return &Queue[T]{}
	}
	q := &Queue[T]{
		size: 0,
	}
	for _, v := range sli {
		q.Enqueue(v)
	}

	return q
}

func (q *Queue[T]) Enqueue(value T) {
	newItem := &node.LinkedNode[T]{
		Value: value,
	}

	if q.first == nil {
		q.first = newItem
	} else {
		q.last.Next = newItem
	}

	q.last = newItem
	q.size++
}

func (q *Queue[T]) Dequeue() T {
	var value T
	if q.first != nil {
		value = q.first.Value
		q.first = q.first.Next
		if q.first == nil {
			q.last = nil
		}
		q.size--
	}
	return value
}

func (q *Queue[T]) Peek() T {
	var value T
	if q.first != nil {
		value = q.first.Value
	}
	return value
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) ToSlice() []T {
	slice := make([]T, 0, q.size)
	current := q.first
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

func (q *Queue[T]) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}
