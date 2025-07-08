package main

import (
	"fmt"

	"github.com/thefrost13/gollections/orderedhashmap"
)

func main() {
	// Create a new OrderedHashMap
	ohm := orderedhashmap.NewOrderedHashMap[string, int]()

	// Set key-value pairs (maintains insertion order)
	ohm.Set("first", 1)
	ohm.Set("second", 2)
	ohm.Set("third", 3)

	// Get values by key
	value, exists := ohm.Get("second")
	if exists {
		fmt.Printf("Value for 'second': %d\n", value)
	}

	// Update existing key (maintains original position)
	ohm.Set("second", 20)

	// Get all keys in insertion order
	keys := ohm.Keys()
	fmt.Printf("Keys in order: %v\n", keys)

	// Get all values in insertion order
	values := ohm.Values()
	fmt.Printf("Values in order: %v\n", values)

	// Convert to slice of key-value pairs
	pairs := ohm.ToSlice()
	fmt.Printf("Key-value pairs: %v\n", pairs)

	// Delete a key
	ohm.Delete("second")

	// Check size
	fmt.Printf("Size: %d\n", ohm.Size())
}
