// Package node provides a generic linked list node implementation
// used by other data structures in the gollections library.
package node

// LinkedNode represents a node in a singly linked list.
// It contains a value of type T and a pointer to the next node.
//
// Type parameters:
//   - T: the value type, can be any type
type LinkedNode[T any] struct {
	Value T              // the value stored in this node
	Next  *LinkedNode[T] // pointer to the next node, nil if this is the last node
}
