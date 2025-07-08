package main

import (
	"fmt"

	"github.com/thefrost13/gollections/hashset"
	"github.com/thefrost13/gollections/orderedhashmap"
	"github.com/thefrost13/gollections/priorityqueue"
	"github.com/thefrost13/gollections/queue"
	"github.com/thefrost13/gollections/stack"
)

func main() {
	fmt.Println("=== Gollections Demo ===")

	// HashSet Demo
	fmt.Println("\n1. HashSet Demo:")
	set := hashset.New([]int{1, 2, 3, 2, 4}) // Duplicates removed
	set.Add(5)
	fmt.Printf("Set contains 3: %t\n", set.Contains(3))
	fmt.Printf("Set size: %d\n", set.Size())
	fmt.Printf("Set elements: %v\n", set.ToSlice())

	// Stack Demo
	fmt.Println("\n2. Stack Demo:")
	stack := stack.New([]string{"first", "second"})
	stack.Push("third")
	fmt.Printf("Top element: %s\n", stack.Peek())
	fmt.Printf("Stack size: %d\n", stack.Size())
	for !stack.IsEmpty() {
		fmt.Printf("Popped: %s\n", stack.Pop())
	}

	// Queue Demo
	fmt.Println("\n3. Queue Demo:")
	queue := queue.New([]int{1, 2, 3})
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Printf("Front element: %d\n", queue.Peek())
	fmt.Printf("Queue size: %d\n", queue.Size())
	for !queue.IsEmpty() {
		fmt.Printf("Dequeued: %d\n", queue.Dequeue())
	}

	// PriorityQueue Demo
	fmt.Println("\n4. PriorityQueue Demo:")
	pq := priorityqueue.New[string]()
	pq.Enqueue("low priority", 10)
	pq.Enqueue("high priority", 1)
	pq.Enqueue("medium priority", 5)
	fmt.Printf("Next element: %s\n", pq.Peek())
	fmt.Printf("Priority queue size: %d\n", pq.Size())
	for !pq.IsEmpty() {
		fmt.Printf("Processing: %s\n", pq.Dequeue())
	}

	// OrderedHashMap Demo
	fmt.Println("\n5. OrderedHashMap Demo:")
	ohm := orderedhashmap.New[string, int]()
	ohm.Set("first", 1)
	ohm.Set("second", 2)
	ohm.Set("third", 3)
	ohm.Set("second", 20) // Update existing key

	fmt.Printf("Keys in order: %v\n", ohm.Keys())
	fmt.Printf("Values in order: %v\n", ohm.Values())
	fmt.Printf("Key-value pairs: %v\n", ohm.ToSlice())

	ohm.Delete("second")
	fmt.Printf("Size after deletion: %d\n", ohm.Size())

	fmt.Println("\n=== Demo Complete ===")
}
