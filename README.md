# Gollections

A generic collections library for Go that provides type-safe data structures using Go generics.

## Features

- **HashSet**: A hash-based set implementation for unique values
- **Stack**: A LIFO (Last In, First Out) stack implementation using linked list
- **Queue**: A FIFO (First In, First Out) queue implementation using linked list  
- **PriorityQueue**: A priority queue implementation using Go's container/heap
- **OrderedHashMap**: A hash map that maintains insertion order of key-value pairs

## Installation

```bash
go get github.com/thefrost13/gollections
```

## Usage

### HashSet

```go
package main

import (
    "fmt"
    "github.com/thefrost13/gollections"
)

func main() {
    // Create a new HashSet
    set := gollections.NewHashSet([]int{1, 2, 3, 2, 4}) // Duplicates are automatically removed
    
    // Add elements
    set.Add(5)
    
    // Check if element exists
    if set.Contains(3) {
        fmt.Println("Set contains 3")
    }
    
    // Get size
    fmt.Printf("Set size: %d\n", set.Size())
    
    // Convert to slice
    slice := set.ToSlice()
    fmt.Printf("Set elements: %v\n", slice)
    
    // Remove element
    set.Remove(2)
    
    // Clear all elements
    set.Clear()
}
```

### Stack

```go
package main

import (
    "fmt"
    "github.com/thefrost13/gollections"
)

func main() {
    // Create a new Stack
    stack := gollections.NewStack([]string{"first", "second"})
    
    // Push elements
    stack.Push("third")
    stack.Push("fourth")
    
    // Peek at top element (doesn't remove)
    top := stack.Peek()
    fmt.Printf("Top element: %s\n", top)
    
    // Pop elements (LIFO order)
    for !stack.IsEmpty() {
        element := stack.Pop()
        fmt.Printf("Popped: %s\n", element)
    }
}
```

### Queue

```go
package main

import (
    "fmt"
    "github.com/thefrost13/gollections"
)

func main() {
    // Create a new Queue
    queue := gollections.NewQueue([]int{1, 2, 3})
    
    // Enqueue elements
    queue.Enqueue(4)
    queue.Enqueue(5)
    
    // Peek at front element (doesn't remove)
    front := queue.Peek()
    fmt.Printf("Front element: %d\n", front)
    
    // Dequeue elements (FIFO order)
    for !queue.IsEmpty() {
        element := queue.Dequeue()
        fmt.Printf("Dequeued: %d\n", element)
    }
}
```

### PriorityQueue

```go
package main

import (
    "fmt"
    "github.com/thefrost13/gollections"
)

func main() {
    // Create a new PriorityQueue
    pq := gollections.NewPriorityQueue[string]()
    
    // Enqueue elements with priorities (lower number = higher priority)
    pq.Enqueue("low priority", 10)
    pq.Enqueue("high priority", 1)
    pq.Enqueue("medium priority", 5)
    
    // Peek at highest priority element
    next := pq.Peek()
    fmt.Printf("Next element: %s\n", next)
    
    // Dequeue elements in priority order
    for !pq.IsEmpty() {
        element := pq.Dequeue()
        fmt.Printf("Processing: %s\n", element)
    }
}
```

### OrderedHashMap

```go
package main

import (
    "fmt"
    "github.com/thefrost13/gollections"
)

func main() {
    // Create a new OrderedHashMap
    ohm := gollections.NewOrderedHashMap[string, int]()
    
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
```

## API Reference

### Common Methods

Most collections implement these common methods:

- `Size() int` - Returns the number of elements
- `IsEmpty() bool` - Returns true if the collection is empty

Some collections also implement:
- `Clear()` - Removes all elements (HashSet, Stack, Queue, PriorityQueue)

### HashSet Methods

- `NewHashSet[T comparable](slice []T) HashSet[T]` - Creates a new HashSet
- `Add(value T)` - Adds an element to the set
- `Remove(value T)` - Removes an element from the set
- `Contains(value T) bool` - Checks if an element exists in the set
- `ToSlice() []T` - Returns a slice containing all elements
- `Equals(other HashSet[T]) bool` - Checks if two sets are equal

### Stack Methods

- `NewStack[T any](slice []T) *Stack[T]` - Creates a new Stack
- `Push(value T)` - Pushes an element onto the stack
- `Pop() T` - Pops and returns the top element
- `Peek() T` - Returns the top element without removing it

### Queue Methods

- `NewQueue[T any](slice []T) *Queue[T]` - Creates a new Queue
- `Enqueue(value T)` - Adds an element to the back of the queue
- `Dequeue() T` - Removes and returns the front element
- `Peek() T` - Returns the front element without removing it

### PriorityQueue Methods

- `NewPriorityQueue[T any]() *PriorityQueue[T]` - Creates a new PriorityQueue
- `Enqueue(value T, priority int)` - Adds an element with a priority
- `Dequeue() T` - Removes and returns the highest priority element
- `Peek() T` - Returns the highest priority element without removing it
- `ToSlice() []T` - Returns a slice containing all elements

### OrderedHashMap Methods

- `NewOrderedHashMap[K comparable, V any]() *OrderedHashMap[K, V]` - Creates a new OrderedHashMap
- `Set(key K, value V)` - Sets a key-value pair (maintains insertion order)
- `Get(key K) (V, bool)` - Gets a value by key, returns value and existence flag
- `Delete(key K)` - Removes a key-value pair from the map
- `Keys() []K` - Returns all keys in insertion order
- `Values() []V` - Returns all values in insertion order
- `ToSlice() []KVPair[K, V]` - Returns all key-value pairs as a slice in insertion order

## Requirements

- Go 1.18 or later (for generics support)

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.