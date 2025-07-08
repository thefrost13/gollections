// Package queue provides a generic FIFO (First In, First Out) queue implementation
// using a linked list for efficient enqueue and dequeue operations.
package queue

import "github.com/thefrost13/gollections/node"

// Queue is a generic FIFO (First In, First Out) data structure implemented using a linked list.
// It provides O(1) enqueue and dequeue operations and maintains elements in insertion order.
// The zero value is ready to use but New should be preferred for initialization.
//
// Type parameters:
//   - T: the element type, can be any type
type Queue[T any] struct {
	first *node.LinkedNode[T] // pointer to the first (oldest) element
	last  *node.LinkedNode[T] // pointer to the last (newest) element
	size  int                 // number of elements in the queue
}

// New creates and returns a new Queue initialized with the elements from the given slice.
// Elements are enqueued in the order they appear in the slice.
// If the slice is nil, an empty queue is returned.
// Time complexity: O(n) where n is the length of the slice.
//
// Parameters:
//   - sli: slice of elements to initialize the queue with, can be nil
//
// Returns:
//   - a new Queue containing the elements from the slice
//
// Example:
//
//	queue := New([]int{1, 2, 3})  // Creates queue with 1 at front
//	emptyQueue := New[string](nil)  // Creates empty queue
func New[T any](sli []T) *Queue[T] {
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

// Enqueue adds an element to the back of the queue.
// The element will be the last to be dequeued (FIFO order).
// Time complexity: O(1).
//
// Parameters:
//   - value: the element to add to the queue
//
// Example:
//
//	queue.Enqueue(42)      // Add 42 to back of queue
//	queue.Enqueue("hello") // Add "hello" to back of queue
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

// Dequeue removes and returns the front element from the queue.
// If the queue is empty, returns the zero value of type T.
// Time complexity: O(1).
//
// Returns:
//   - the front element of the queue, or zero value if queue is empty
//
// Example:
//
//	value := queue.Dequeue()  // Remove and return front element
//	if !queue.IsEmpty() {
//	    next := queue.Dequeue()  // Get next element
//	}
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

// Peek returns the front element of the queue without removing it.
// If the queue is empty, returns the zero value of type T.
// Time complexity: O(1).
//
// Returns:
//   - the front element of the queue, or zero value if queue is empty
//
// Example:
//
//	front := queue.Peek()  // Look at front element without removing it
//	if !queue.IsEmpty() {
//	    fmt.Printf("Front element: %v\n", front)
//	}
func (q *Queue[T]) Peek() T {
	var value T
	if q.first != nil {
		value = q.first.Value
	}
	return value
}

// Size returns the number of elements in the queue.
// Time complexity: O(1).
//
// Returns:
//   - the number of elements in the queue
//
// Example:
//
//	count := queue.Size()
//	fmt.Printf("Queue has %d elements\n", count)
func (q *Queue[T]) Size() int {
	return q.size
}

// IsEmpty returns true if the queue contains no elements.
// Time complexity: O(1).
//
// Returns:
//   - true if the queue is empty, false otherwise
//
// Example:
//
//	if queue.IsEmpty() {
//	    fmt.Println("Queue is empty")
//	}
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// ToSlice returns a slice containing all elements in the queue in FIFO order.
// The first element in the slice is the front of the queue.
// The returned slice is a copy and modifications to it will not affect the queue.
// Time complexity: O(n) where n is the number of elements.
//
// Returns:
//   - a slice containing all queue elements in FIFO order
//
// Example:
//
//	elements := queue.ToSlice()  // Get all elements as slice
//	fmt.Printf("Queue contents: %v\n", elements)
func (q *Queue[T]) ToSlice() []T {
	slice := make([]T, 0, q.size)
	current := q.first
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

// Clear removes all elements from the queue, making it empty.
// Time complexity: O(1).
//
// Example:
//
//	queue.Clear()  // Remove all elements
//	fmt.Printf("Queue is now empty: %t\n", queue.IsEmpty())
func (q *Queue[T]) Clear() {
	q.first = nil
	q.last = nil
	q.size = 0
}
