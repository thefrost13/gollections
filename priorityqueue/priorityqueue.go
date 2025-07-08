// Package priorityqueue provides a generic priority queue implementation using Go's container/heap.
// Elements are dequeued in priority order, with lower priority values having higher precedence.
package priorityqueue

import "container/heap"

// priorityQueueItem represents an item in the priority queue with its associated priority.
// Lower priority values indicate higher precedence.
type priorityQueueItem[T any] struct {
	value    T   // the actual value stored in the queue
	priority int // the priority of the item (lower = higher precedence)
}

// PriorityQueue is a generic priority queue that dequeues elements in priority order.
// It's implemented using Go's container/heap package for efficient operations.
// Lower priority values have higher precedence (min-heap behavior).
//
// Type parameters:
//   - T: the element type, can be any type
type PriorityQueue[T any] []priorityQueueItem[T]

// Len returns the number of elements in the priority queue.
// This method is required by the heap.Interface.
func (pq PriorityQueue[T]) Len() int { return len(pq) }

// Less compares two elements by their priority.
// Returns true if element i has higher priority than element j.
// This method is required by the heap.Interface.
func (pq PriorityQueue[T]) Less(i, j int) bool { return pq[i].priority < pq[j].priority }

// Swap exchanges the elements at positions i and j.
// This method is required by the heap.Interface.
func (pq PriorityQueue[T]) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

// Push adds an element to the priority queue.
// This method is required by the heap.Interface and should not be called directly.
// Use Enqueue instead.
func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(priorityQueueItem[T]))
}

// Pop removes and returns the element with the highest priority.
// This method is required by the heap.Interface and should not be called directly.
// Use Dequeue instead.
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[0 : n-1]
	return last
}

// NewPriorityQueue creates and returns a new empty PriorityQueue.
// The returned queue is ready to use and will maintain elements in priority order.
// Time complexity: O(1).
//
// Returns:
//   - a new empty PriorityQueue
//
// Example:
//
//	pq := NewPriorityQueue[string]()
//	pq.Enqueue("low priority", 10)
//	pq.Enqueue("high priority", 1)
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}

// Enqueue adds an element to the priority queue with the specified priority.
// Elements with lower priority values will be dequeued first.
// Time complexity: O(log n) where n is the number of elements.
//
// Parameters:
//   - value: the element to add to the queue
//   - priority: the priority of the element (lower = higher precedence)
//
// Example:
//
//	pq.Enqueue("urgent", 1)      // High priority
//	pq.Enqueue("normal", 5)      // Medium priority
//	pq.Enqueue("low", 10)        // Low priority
func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	item := priorityQueueItem[T]{value: value, priority: priority}
	heap.Push(pq, item)
}

// Dequeue removes and returns the element with the highest priority (lowest priority value).
// If the queue is empty, returns the zero value of type T.
// Time complexity: O(log n) where n is the number of elements.
//
// Returns:
//   - the element with the highest priority, or zero value if queue is empty
//
// Example:
//
//	value := pq.Dequeue()  // Remove and return highest priority element
//	if !pq.IsEmpty() {
//	    next := pq.Dequeue()  // Get next highest priority element
//	}
func (pq *PriorityQueue[T]) Dequeue() T {
	var value T
	if pq.Len() != 0 {
		item := heap.Pop(pq)
		return item.(priorityQueueItem[T]).value
	}
	return value
}

// Peek returns the element with the highest priority without removing it.
// If the queue is empty, returns the zero value of type T.
// Time complexity: O(1).
//
// Returns:
//   - the element with the highest priority, or zero value if queue is empty
//
// Example:
//
//	next := pq.Peek()  // Look at highest priority element without removing it
//	if !pq.IsEmpty() {
//	    fmt.Printf("Next element: %v\n", next)
//	}
func (pq *PriorityQueue[T]) Peek() T {
	var value T
	if pq.Len() != 0 {
		value = (*pq)[0].value
	}
	return value
}

// Size returns the number of elements in the priority queue.
// Time complexity: O(1).
//
// Returns:
//   - the number of elements in the priority queue
//
// Example:
//
//	count := pq.Size()
//	fmt.Printf("Priority queue has %d elements\n", count)
func (pq *PriorityQueue[T]) Size() int {
	return pq.Len()
}

// IsEmpty returns true if the priority queue contains no elements.
// Time complexity: O(1).
//
// Returns:
//   - true if the priority queue is empty, false otherwise
//
// Example:
//
//	if pq.IsEmpty() {
//	    fmt.Println("Priority queue is empty")
//	}
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

// Clear removes all elements from the priority queue, making it empty.
// Time complexity: O(1).
//
// Example:
//
//	pq.Clear()  // Remove all elements
//	fmt.Printf("Priority queue is now empty: %t\n", pq.IsEmpty())
func (pq *PriorityQueue[T]) Clear() {
	*pq = (*pq)[:0]
}

// ToSlice returns a slice containing all elements in the priority queue.
// The elements are returned in heap order, not priority order.
// For priority-ordered access, use repeated Dequeue operations.
// The returned slice is a copy and modifications to it will not affect the queue.
// Time complexity: O(n) where n is the number of elements.
//
// Returns:
//   - a slice containing all elements in the priority queue
//
// Example:
//
//	elements := pq.ToSlice()  // Get all elements as slice
//	fmt.Printf("Queue contents: %v\n", elements)
func (pq *PriorityQueue[T]) ToSlice() []T {
	slice := make([]T, pq.Len())
	for i, item := range *pq {
		slice[i] = item.value
	}
	return slice
}
