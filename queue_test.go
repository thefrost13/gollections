package gollections

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	t.Run("NewQueue with nil slice", func(t *testing.T) {
		q := NewQueue[int](nil)
		if q == nil {
			t.Fatal("Expected queue to be initialized, got nil")
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
		if q.first != nil {
			t.Error("Expected first to be nil")
		}
		if q.last != nil {
			t.Error("Expected last to be nil")
		}
	})

	t.Run("NewQueue with empty slice", func(t *testing.T) {
		q := NewQueue[int]([]int{})
		if q == nil {
			t.Fatal("Expected queue to be initialized, got nil")
		}
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("NewQueue with single element", func(t *testing.T) {
		q := NewQueue[int]([]int{42})
		if q.Size() != 1 {
			t.Errorf("Expected size 1, got %d", q.Size())
		}
		if q.Peek() != 42 {
			t.Errorf("Expected peek value 42, got %d", q.Peek())
		}
	})

	t.Run("NewQueue with multiple elements", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		q := NewQueue[int](values)
		if q.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), q.Size())
		}

		// Check that elements are in FIFO order
		for i, expected := range values {
			actual := q.Dequeue()
			if actual != expected {
				t.Errorf("At position %d, expected %d, got %d", i, expected, actual)
			}
		}
	})

	t.Run("NewQueue with string type", func(t *testing.T) {
		values := []string{"hello", "world", "test"}
		q := NewQueue[string](values)
		if q.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), q.Size())
		}
		if q.Peek() != "hello" {
			t.Errorf("Expected peek value 'hello', got '%s'", q.Peek())
		}
	})
}

func TestQueueEnqueue(t *testing.T) {
	t.Run("Enqueue to empty queue", func(t *testing.T) {
		q := NewQueue[int](nil)
		q.Enqueue(100)

		if q.Size() != 1 {
			t.Errorf("Expected size 1, got %d", q.Size())
		}
		if q.Peek() != 100 {
			t.Errorf("Expected peek value 100, got %d", q.Peek())
		}
		if q.first == nil {
			t.Error("Expected first to be set")
		}
		if q.last == nil {
			t.Error("Expected last to be set")
		}
		if q.first != q.last {
			t.Error("Expected first and last to be the same for single element")
		}
	})

	t.Run("Enqueue multiple elements", func(t *testing.T) {
		q := NewQueue[int](nil)
		values := []int{1, 2, 3, 4, 5}

		for i, value := range values {
			q.Enqueue(value)
			if q.Size() != i+1 {
				t.Errorf("After enqueue %d, expected size %d, got %d", value, i+1, q.Size())
			}
			if q.Peek() != values[0] {
				t.Errorf("After enqueue %d, expected peek value %d, got %d", value, values[0], q.Peek())
			}
		}
	})

	t.Run("Enqueue maintains FIFO order", func(t *testing.T) {
		q := NewQueue[int](nil)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		if q.Dequeue() != 1 {
			t.Error("Expected first dequeue to return 1")
		}
		if q.Dequeue() != 2 {
			t.Error("Expected second dequeue to return 2")
		}
		if q.Dequeue() != 3 {
			t.Error("Expected third dequeue to return 3")
		}
	})
}

func TestQueueDequeue(t *testing.T) {
	t.Run("Dequeue from empty queue", func(t *testing.T) {
		q := NewQueue[int](nil)
		result := q.Dequeue()

		var expected int
		if result != expected {
			t.Errorf("Expected zero value %d, got %d", expected, result)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size to remain 0, got %d", q.Size())
		}
	})

	t.Run("Dequeue single element", func(t *testing.T) {
		q := NewQueue[int]([]int{42})
		result := q.Dequeue()

		if result != 42 {
			t.Errorf("Expected 42, got %d", result)
		}
		if q.first != nil {
			t.Error("Expected first to be nil after dequeue")
		}
		if q.last != nil {
			t.Error("Expected last to be nil after dequeue")
		}
	})

	t.Run("Dequeue multiple elements", func(t *testing.T) {
		values := []int{10, 20, 30, 40, 50}
		q := NewQueue[int](values)

		for i, expected := range values {
			result := q.Dequeue()
			if result != expected {
				t.Errorf("Dequeue %d: expected %d, got %d", i, expected, result)
			}
		}

		// Queue should be empty now
		if q.first != nil {
			t.Error("Expected first to be nil after all dequeues")
		}
		if q.last != nil {
			t.Error("Expected last to be nil after all dequeues")
		}
	})

	t.Run("Dequeue with string type", func(t *testing.T) {
		q := NewQueue[string]([]string{"first", "second", "third"})

		if q.Dequeue() != "first" {
			t.Error("Expected first dequeue to return 'first'")
		}
		if q.Dequeue() != "second" {
			t.Error("Expected second dequeue to return 'second'")
		}
		if q.Dequeue() != "third" {
			t.Error("Expected third dequeue to return 'third'")
		}

		// Dequeue from empty queue should return zero value
		result := q.Dequeue()
		if result != "" {
			t.Errorf("Expected empty string, got '%s'", result)
		}
	})
}

func TestQueuePeek(t *testing.T) {
	t.Run("Peek empty queue", func(t *testing.T) {
		q := NewQueue[int](nil)
		result := q.Peek()

		var expected int
		if result != expected {
			t.Errorf("Expected zero value %d, got %d", expected, result)
		}
		if q.Size() != 0 {
			t.Errorf("Expected size to remain 0, got %d", q.Size())
		}
	})

	t.Run("Peek single element", func(t *testing.T) {
		q := NewQueue[int]([]int{99})
		result := q.Peek()

		if result != 99 {
			t.Errorf("Expected 99, got %d", result)
		}
		// Peek should not modify the queue
		if q.Size() != 1 {
			t.Errorf("Expected size to remain 1, got %d", q.Size())
		}
		// Should be able to peek again with same result
		if q.Peek() != 99 {
			t.Error("Second peek should return the same value")
		}
	})

	t.Run("Peek multiple elements", func(t *testing.T) {
		values := []int{1, 2, 3, 4}
		q := NewQueue[int](values)

		// Peek should always return the first element
		for i := 0; i < 5; i++ {
			result := q.Peek()
			if result != values[0] {
				t.Errorf("Peek %d: expected %d, got %d", i, values[0], result)
			}
		}

		// Size should remain unchanged
		if q.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), q.Size())
		}
	})

	t.Run("Peek after operations", func(t *testing.T) {
		q := NewQueue[int]([]int{1, 2, 3})

		// Initial peek
		if q.Peek() != 1 {
			t.Error("Initial peek should return 1")
		}

		// After dequeue
		q.Dequeue()
		if q.Peek() != 2 {
			t.Error("After dequeue, peek should return 2")
		}

		// After enqueue
		q.Enqueue(4)
		if q.Peek() != 2 {
			t.Error("After enqueue, peek should still return 2")
		}
	})
}

func TestQueueSize(t *testing.T) {
	t.Run("Size of empty queue", func(t *testing.T) {
		q := NewQueue[int](nil)
		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})

	t.Run("Size updates with operations", func(t *testing.T) {
		q := NewQueue[int](nil)

		// After enqueues
		q.Enqueue(1)
		if q.Size() != 1 {
			t.Errorf("After first enqueue, expected size 1, got %d", q.Size())
		}

		q.Enqueue(2)
		if q.Size() != 2 {
			t.Errorf("After second enqueue, expected size 2, got %d", q.Size())
		}

		q.Enqueue(3)
		if q.Size() != 3 {
			t.Errorf("After third enqueue, expected size 3, got %d", q.Size())
		}

		// After dequeues
		q.Dequeue()
		if q.Size() != 2 {
			t.Errorf("After first dequeue, expected size 2, got %d", q.Size())
		}

		q.Dequeue()
		if q.Size() != 1 {
			t.Errorf("After second dequeue, expected size 1, got %d", q.Size())
		}

		q.Dequeue()
		if q.Size() != 0 {
			t.Errorf("After third dequeue, expected size 0, got %d", q.Size())
		}
	})
}

func TestQueueToSlice(t *testing.T) {
	t.Run("empty queue to slice", func(t *testing.T) {
		q := NewQueue[int](nil)
		slice := q.ToSlice()
		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(slice))
		}
	})

	t.Run("queue to slice maintains FIFO order", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		q := NewQueue[int](values)
		slice := q.ToSlice()

		if len(slice) != len(values) {
			t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
		}

		if !reflect.DeepEqual(slice, values) {
			t.Errorf("Expected slice %v, got %v", values, slice)
		}
	})

	t.Run("string queue to slice", func(t *testing.T) {
		values := []string{"first", "second", "third"}
		q := NewQueue[string](values)
		slice := q.ToSlice()

		if len(slice) != len(values) {
			t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
		}

		if !reflect.DeepEqual(slice, values) {
			t.Errorf("Expected slice %v, got %v", values, slice)
		}
	})

	t.Run("slice independence", func(t *testing.T) {
		q := NewQueue[int]([]int{1, 2, 3})
		slice := q.ToSlice()

		// Modify the slice
		slice[0] = 999

		// Original queue should be unchanged
		if q.Peek() != 1 {
			t.Error("Modifying returned slice should not affect original queue")
		}
	})

	t.Run("partial queue to slice", func(t *testing.T) {
		q := NewQueue[int]([]int{1, 2, 3, 4, 5})

		// Dequeue some elements
		q.Dequeue()
		q.Dequeue()

		slice := q.ToSlice()
		expected := []int{3, 4, 5}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})
}

func TestQueueBehavior(t *testing.T) {
	t.Run("FIFO behavior", func(t *testing.T) {
		q := NewQueue[int](nil)
		expected := []int{1, 2, 3, 4, 5}

		// Enqueue all elements
		for _, value := range expected {
			q.Enqueue(value)
		}

		// Dequeue all elements and verify order
		var result []int
		for q.Size() > 0 {
			result = append(result, q.Dequeue())
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Mixed operations", func(t *testing.T) {
		q := NewQueue[int](nil)

		q.Enqueue(1)
		q.Enqueue(2)

		if q.Dequeue() != 1 {
			t.Error("First dequeue should return 1")
		}

		q.Enqueue(3)
		q.Enqueue(4)

		if q.Dequeue() != 2 {
			t.Error("Second dequeue should return 2")
		}
		if q.Dequeue() != 3 {
			t.Error("Third dequeue should return 3")
		}
		if q.Dequeue() != 4 {
			t.Error("Fourth dequeue should return 4")
		}

		if q.Size() != 0 {
			t.Errorf("Expected empty queue, size is %d", q.Size())
		}
	})

	t.Run("Queue with custom struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		q := NewQueue[Person](nil)
		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Bob", Age: 25}

		q.Enqueue(p1)
		q.Enqueue(p2)

		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}

		result1 := q.Dequeue()
		if result1.Name != "Alice" || result1.Age != 30 {
			t.Errorf("Expected Alice(30), got %s(%d)", result1.Name, result1.Age)
		}

		result2 := q.Peek()
		if result2.Name != "Bob" || result2.Age != 25 {
			t.Errorf("Expected Bob(25), got %s(%d)", result2.Name, result2.Age)
		}
	})
}

// Benchmark tests
func BenchmarkQueueEnqueue(b *testing.B) {
	q := NewQueue[int](nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func BenchmarkQueueDequeue(b *testing.B) {
	q := NewQueue[int](nil)
	// Pre-populate the queue
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Dequeue()
	}
}

func BenchmarkQueuePeek(b *testing.B) {
	q := NewQueue[int]([]int{1, 2, 3, 4, 5})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Peek()
	}
}

func BenchmarkQueueMixedOperations(b *testing.B) {
	q := NewQueue[int](nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
		if i%2 == 0 && q.Size() > 0 {
			q.Dequeue()
		}
		q.Peek()
	}
}
