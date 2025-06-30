package gollections

type HashSet[T comparable] map[T]bool

func NewHashSet[T comparable](sli []T) HashSet[T] {
	set := make(HashSet[T])
	for _, v := range sli {
		set[v] = true
	}
	return set
}

func (s HashSet[T]) Add(value T) {
	s[value] = true
}

func (s HashSet[T]) Remove(value T) {
	delete(s, value)
}

func (s HashSet[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

func (s HashSet[T]) Size() int {
	return len(s)
}

func (s HashSet[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s HashSet[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

func (s HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}

func (s HashSet[T]) Equals(other HashSet[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
