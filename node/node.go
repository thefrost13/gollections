package node

type LinkedNode[T any] struct {
	Value T
	Next  *LinkedNode[T]
}
