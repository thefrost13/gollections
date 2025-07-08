// Package hashset provides a generic hash-based set implementation that stores
// unique values and provides fast membership testing, insertion, and deletion.
package hashset

// HashSet is a generic set data structure that stores unique values of comparable types.
// It's implemented using Go's built-in map for O(1) average-case operations.
// The zero value is ready to use but NewHashSet should be preferred for initialization.
//
// Type parameters:
//   - T: the element type, must be comparable
type HashSet[T comparable] map[T]bool

// NewHashSet creates and returns a new HashSet initialized with the elements from the given slice.
// Duplicate elements in the slice are automatically removed.
// If the slice is nil, an empty set is returned.
// Time complexity: O(n) where n is the length of the slice.
//
// Parameters:
//   - sli: slice of elements to initialize the set with, can be nil
//
// Returns:
//   - a new HashSet containing the unique elements from the slice
//
// Example:
//
//	set := NewHashSet([]int{1, 2, 3, 2, 1})  // Creates set with {1, 2, 3}
//	emptySet := NewHashSet[string](nil)       // Creates empty set
func NewHashSet[T comparable](sli []T) HashSet[T] {
	set := make(HashSet[T])
	for _, v := range sli {
		set[v] = true
	}
	return set
}

// Add inserts an element into the set.
// If the element already exists, the operation is a no-op.
// Time complexity: O(1) average case.
//
// Parameters:
//   - value: the element to add to the set
//
// Example:
//
//	set.Add(42)      // Add 42 to the set
//	set.Add("hello") // Add "hello" to the set
func (s HashSet[T]) Add(value T) {
	s[value] = true
}

// Remove deletes an element from the set.
// If the element doesn't exist, the operation is a no-op.
// Time complexity: O(1) average case.
//
// Parameters:
//   - value: the element to remove from the set
//
// Example:
//
//	set.Remove(42)      // Remove 42 from the set
//	set.Remove("hello") // Remove "hello" from the set
func (s HashSet[T]) Remove(value T) {
	delete(s, value)
}

// Contains checks if an element exists in the set.
// Time complexity: O(1) average case.
//
// Parameters:
//   - value: the element to check for membership
//
// Returns:
//   - true if the element exists in the set, false otherwise
//
// Example:
//
//	if set.Contains(42) {
//	    fmt.Println("Set contains 42")
//	}
func (s HashSet[T]) Contains(value T) bool {
	_, exists := s[value]
	return exists
}

// Size returns the number of elements in the set.
// Time complexity: O(1).
//
// Returns:
//   - the number of elements in the set
//
// Example:
//
//	count := set.Size()
//	fmt.Printf("Set has %d elements\n", count)
func (s HashSet[T]) Size() int {
	return len(s)
}

// IsEmpty returns true if the set contains no elements.
// Time complexity: O(1).
//
// Returns:
//   - true if the set is empty, false otherwise
//
// Example:
//
//	if set.IsEmpty() {
//	    fmt.Println("Set is empty")
//	}
func (s HashSet[T]) IsEmpty() bool {
	return len(s) == 0
}

// Clear removes all elements from the set, making it empty.
// Time complexity: O(n) where n is the number of elements.
//
// Example:
//
//	set.Clear()  // Remove all elements
//	fmt.Printf("Set is now empty: %t\n", set.IsEmpty())
func (s HashSet[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

// ToSlice returns a slice containing all elements in the set.
// The order of elements in the slice is not guaranteed to be consistent.
// The returned slice is a copy and modifications to it will not affect the set.
// Time complexity: O(n) where n is the number of elements.
//
// Returns:
//   - a slice containing all elements in the set
//
// Example:
//
//	elements := set.ToSlice()  // Get all elements as slice
//	fmt.Printf("Set contents: %v\n", elements)
func (s HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(s))
	for key := range s {
		slice = append(slice, key)
	}
	return slice
}

// Equals checks if this set contains exactly the same elements as another set.
// The comparison is based on set membership, not on the order of elements.
// Time complexity: O(n) where n is the number of elements in the larger set.
//
// Parameters:
//   - other: the HashSet to compare with
//
// Returns:
//   - true if both sets contain the same elements, false otherwise
//
// Example:
//
//	set1 := NewHashSet([]int{1, 2, 3})
//	set2 := NewHashSet([]int{3, 2, 1})  // Different order
//	if set1.Equals(set2) {
//	    fmt.Println("Sets are equal")  // This will print
//	}
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
