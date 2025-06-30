package gollections

type Stack[T any] struct {
	top  *LinkedNode[T]
	size int
}

func NewStack[T any](sli []T) *Stack[T] {
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

func (s *Stack[T]) Push(value T) {
	newItem := &LinkedNode[T]{
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

func (s *Stack[T]) Pop() T {
	var value T
	if s.top != nil {
		value = s.top.Value
		s.top = s.top.Next
		s.size--
	}
	return value
}

func (s *Stack[T]) Peek() T {
	var value T
	if s.top != nil {
		value = s.top.Value
	}
	return value
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) ToSlice() []T {
	slice := make([]T, 0, s.size)
	current := s.top
	for current != nil {
		slice = append(slice, current.Value)
		current = current.Next
	}
	return slice
}

func (s *Stack[T]) Clear() {
	s.top = nil
	s.size = 0
}
