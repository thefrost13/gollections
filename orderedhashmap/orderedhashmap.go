// Package orderedhashmap provides an ordered hash map implementation that maintains
// insertion order of key-value pairs while providing O(1) access, insertion, and deletion.
package orderedhashmap

import "github.com/thefrost13/gollections/node"

// OrderedHashMap is a generic hash map that maintains the insertion order of key-value pairs.
// It combines the fast access of a hash map with the ordered iteration of a linked list.
// The zero value is ready to use but New should be preferred for initialization.
//
// Type parameters:
//   - K: the key type, must be comparable
//   - V: the value type, can be any type
type OrderedHashMap[K comparable, V any] struct {
	first *node.LinkedNode[KVPair[K, V]]       // pointer to the first node in insertion order
	last  *node.LinkedNode[KVPair[K, V]]       // pointer to the last node in insertion order
	size  int                                  // number of key-value pairs in the map
	m     map[K]*node.LinkedNode[KVPair[K, V]] // hash map for O(1) key lookup
}

// KVPair represents a key-value pair stored in the OrderedHashMap.
// It is used in the linked list to maintain insertion order.
type KVPair[K comparable, V any] struct {
	Key   K // the key of the pair
	Value V // the value associated with the key
}

// New creates and returns a new empty OrderedHashMap.
// The returned map is ready to use and will maintain insertion order of key-value pairs.
//
// Example:
//
//	ohm := New[string, int]()
//	ohm.Set("first", 1)
//	ohm.Set("second", 2)
func New[K comparable, V any]() *OrderedHashMap[K, V] {
	return &OrderedHashMap[K, V]{
		m: make(map[K]*node.LinkedNode[KVPair[K, V]]),
	}
}

// Set inserts or updates a key-value pair in the OrderedHashMap.
// If the key already exists, its value is updated but its position in the insertion order is preserved.
// If the key is new, it is added to the end of the insertion order.
// Time complexity: O(1) average case.
//
// Parameters:
//   - key: the key to insert or update
//   - value: the value to associate with the key
//
// Example:
//
//	ohm.Set("name", "John")      // Insert new key-value pair
//	ohm.Set("name", "Jane")      // Update existing key, maintains position
func (ohm *OrderedHashMap[K, V]) Set(key K, value V) {
	if node, exists := ohm.m[key]; exists {
		node.Value.Value = value
		return
	}

	newItem := &node.LinkedNode[KVPair[K, V]]{
		Value: KVPair[K, V]{Key: key, Value: value},
	}

	if ohm.first == nil {
		ohm.first = newItem
	} else {
		ohm.last.Next = newItem
	}
	ohm.last = newItem
	ohm.size++
	ohm.m[key] = newItem
}

// Get retrieves the value associated with the given key from the OrderedHashMap.
// It returns the value and a boolean indicating whether the key was found.
// Time complexity: O(1) average case.
//
// Parameters:
//   - key: the key to look up
//
// Returns:
//   - value: the value associated with the key, or zero value if key not found
//   - exists: true if the key exists in the map, false otherwise
//
// Example:
//
//	value, exists := ohm.Get("name")
//	if exists {
//	    fmt.Printf("Found: %v\n", value)
//	}
func (ohm *OrderedHashMap[K, V]) Get(key K) (V, bool) {
	if node, exists := ohm.m[key]; exists {
		return node.Value.Value, true
	}
	var zeroValue V
	return zeroValue, false
}

// Delete removes the key-value pair with the given key from the OrderedHashMap.
// If the key doesn't exist, the operation is a no-op.
// The insertion order of remaining elements is preserved.
// Time complexity: O(n) worst case due to linked list traversal.
//
// Parameters:
//   - key: the key to remove from the map
//
// Example:
//
//	ohm.Delete("name")  // Remove the key-value pair with key "name"
func (ohm *OrderedHashMap[K, V]) Delete(key K) {
	if node, exists := ohm.m[key]; exists {
		delete(ohm.m, key)
		if node == ohm.first {
			ohm.first = node.Next
		} else {
			prev := ohm.first
			for prev != nil && prev.Next != node {
				prev = prev.Next
			}
			if prev != nil {
				prev.Next = node.Next
			}
		}
		if node == ohm.last {
			// Find the previous node to update last pointer
			if ohm.first == nil {
				ohm.last = nil
			} else {
				prev := ohm.first
				for prev != nil && prev.Next != nil {
					prev = prev.Next
				}
				ohm.last = prev
			}
		}
		ohm.size--
	}
}

// Size returns the number of key-value pairs in the OrderedHashMap.
// Time complexity: O(1).
//
// Returns:
//   - the number of key-value pairs in the map
//
// Example:
//
//	count := ohm.Size()
//	fmt.Printf("Map contains %d pairs\n", count)
func (ohm *OrderedHashMap[K, V]) Size() int {
	return ohm.size
}

// IsEmpty returns true if the OrderedHashMap contains no key-value pairs.
// Time complexity: O(1).
//
// Returns:
//   - true if the map is empty, false otherwise
//
// Example:
//
//	if ohm.IsEmpty() {
//	    fmt.Println("Map is empty")
//	}
func (ohm *OrderedHashMap[K, V]) IsEmpty() bool {
	return ohm.size == 0
}

// Keys returns a slice containing all keys in the OrderedHashMap in insertion order.
// The returned slice is a copy and modifications to it will not affect the map.
// Time complexity: O(n) where n is the number of key-value pairs.
//
// Returns:
//   - a slice containing all keys in insertion order
//
// Example:
//
//	keys := ohm.Keys()
//	fmt.Printf("Keys in order: %v\n", keys)
func (ohm *OrderedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		keys = append(keys, node.Value.Key)
	}
	return keys
}

// Values returns a slice containing all values in the OrderedHashMap in insertion order.
// The returned slice is a copy and modifications to it will not affect the map.
// Time complexity: O(n) where n is the number of key-value pairs.
//
// Returns:
//   - a slice containing all values in insertion order
//
// Example:
//
//	values := ohm.Values()
//	fmt.Printf("Values in order: %v\n", values)
func (ohm *OrderedHashMap[K, V]) Values() []V {
	values := make([]V, 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		values = append(values, node.Value.Value)
	}
	return values
}

// ToSlice returns a slice containing all key-value pairs in the OrderedHashMap in insertion order.
// Each element in the returned slice is a KVPair containing both the key and value.
// The returned slice is a copy and modifications to it will not affect the map.
// Time complexity: O(n) where n is the number of key-value pairs.
//
// Returns:
//   - a slice of KVPair containing all key-value pairs in insertion order
//
// Example:
//
//	pairs := ohm.ToSlice()
//	for _, pair := range pairs {
//	    fmt.Printf("%v: %v\n", pair.Key, pair.Value)
//	}
func (ohm *OrderedHashMap[K, V]) ToSlice() []KVPair[K, V] {
	slice := make([]KVPair[K, V], 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		slice = append(slice, node.Value)
	}
	return slice
}
