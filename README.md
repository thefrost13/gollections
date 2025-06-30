# Gollections

A generic collections library for Go that provides type-safe data structures using Go generics.

## Features

- **HashSet**: A hash-based set implementation for unique values
- **Stack**: A LIFO (Last In, First Out) stack implementation using linked list
- **Queue**: A FIFO (First In, First Out) queue implementation using linked list  
- **PriorityQueue**: A priority queue implementation using Go's container/heap

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

## API Reference

### Common Methods

All collections implement these common methods:

- `Size() int` - Returns the number of elements
- `IsEmpty() bool` - Returns true if the collection is empty
- `Clear()` - Removes all elements

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

## Requirements

- Go 1.18 or later (for generics support)

## License

This project is licensed under the BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.