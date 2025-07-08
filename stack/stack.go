// Package stack provides a generic LIFO (Last In, First Out) stack implementation
// using a linked list for efficient push and pop operations.
package stack

import "github.com/thefrost13/gollections/node"

// Stack is a generic LIFO (Last In, First Out) data structure implemented using a linked list.
// It provides O(1) push and pop operations and maintains elements in reverse order of insertion.
// The zero value is ready to use but New should be preferred for initialization.
//
// Type parameters:
//   - T: the element type, can be any type
type Stack[T any] struct {
	top  *node.LinkedNode[T] // pointer to the top (most recently added) element
	size int                 // number of elements in the stack
}

// New creates and returns a new Stack initialized with the elements from the given slice.
// Elements are pushed in the order they appear in the slice, so the last element becomes the top.
// If the slice is nil, an empty stack is returned.
// Time complexity: O(n) where n is the length of the slice.
//
// Parameters:
//   - sli: slice of elements to initialize the stack with, can be nil
//
// Returns:
//   - a new Stack containing the elements from the slice
//
// Example:
//
//	stack := New([]int{1, 2, 3})  // Creates stack with 3 on top
//	emptyStack := New[string](nil)  // Creates empty stack
func New[T any](sli []T) *Stack[T] {
	if sli == nil {
		return &Stack[T]{}
	}
	s := &Stack[T]{
		size: 0,
	}
	for _, v := range sli {
		s.Push(v)
	}

	return s
}

// Push adds an element to the top of the stack.
// The element becomes the new top and will be the first to be popped.
// Time complexity: O(1).
//
// Parameters:
//   - value: the element to add to the stack
//
// Example:
//
//	stack.Push(42)      // Add 42 to top of stack
//	stack.Push("hello") // Add "hello" to top of stack
func (s *Stack[T]) Push(value T) {
	newItem := &node.LinkedNode[T]{
		Value: value,
	}

	if s.top == nil {
		s.top = newItem
	} else {
		newItem.Next = s.top
		s.top = newItem
	}
	s.size++
}

// Pop removes and returns the top element from the stack.
// If the stack is empty, returns the zero value of type T.
// Time complexity: O(1).
//
// Returns:
//   - the top element of the stack, or zero value if stack is empty
//
// Example:
//
//	value := stack.Pop()  // Remove and return top element
//	if !stack.IsEmpty() {
//	    next := stack.Pop()  // Get next element
//	}
func (s *Stack[T]) Pop() T {
	var value T
	if s.top != nil {
		value = s.top.Value
		s.top = s.top.Next
		s.size--
	}
	return value
}

// Peek returns the top element of the stack without removing it.
// If the stack is empty, returns the zero value of type T.
// Time complexity: O(1).
//
// Returns:
//   - the top element of the stack, or zero value if stack is empty
//
// Example:
//
//	top := stack.Peek()  // Look at top element without removing it
//	if !stack.IsEmpty() {
//	    fmt.Printf("Top element: %v\n", top)
//	}
func (s *Stack[T]) Peek() T {
	var value T
	if s.top != nil {
		value = s.top.Value
	}
	return value
}

// Size returns the number of elements in the stack.
// Time complexity: O(1).
//
// Returns:
//   - the number of elements in the stack
//
// Example:
//
//	count := stack.Size()
//	fmt.Printf("Stack has %d elements\n", count)
func (s *Stack[T]) Size() int {
	return s.size
}

// IsEmpty returns true if the stack contains no elements.
// Time complexity: O(1).
//
// Returns:
//   - true if the stack is empty, false otherwise
//
// Example:
//
//	if stack.IsEmpty() {
//	    fmt.Println("Stack is empty")
//	}
func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

// ToSlice returns a slice containing all elements in the stack in LIFO order.
// The first element in the slice is the top of the stack.
// The returned slice is a copy and modifications to it will not affect the stack.
// Time complexity: O(n) where n is the number of elements.
//
// Returns:
//   - a slice containing all stack elements in LIFO order
//
// Example:
//
//	elements := stack.ToSlice()  // Get all elements as slice
//	fmt.Printf("Stack contents: %v\n", elements)
func (s *Stack[T]) ToSlice() []T {
	slice := make([]T, 0, s.size)
	current := s.top
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

// Clear removes all elements from the stack, making it empty.
// Time complexity: O(1).
//
// Example:
//
//	stack.Clear()  // Remove all elements
//	fmt.Printf("Stack is now empty: %t\n", stack.IsEmpty())
func (s *Stack[T]) Clear() {
	s.top = nil
	s.size = 0
}
