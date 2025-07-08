package priorityqueue

import (
	"reflect"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	t.Run("create new priority queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		if pq == nil {
			t.Fatal("NewPriorityQueue should not return nil")
		}
		if pq.Size() != 0 {
			t.Errorf("Expected size 0, got %d", pq.Size())
		}
		if !pq.IsEmpty() {
			t.Error("Expected new priority queue to be empty")
		}
	})

	t.Run("create priority queue with different types", func(t *testing.T) {
		stringPQ := NewPriorityQueue[string]()
		if stringPQ.Size() != 0 {
			t.Errorf("Expected string priority queue size 0, got %d", stringPQ.Size())
		}

		boolPQ := NewPriorityQueue[bool]()
		if boolPQ.Size() != 0 {
			t.Errorf("Expected bool priority queue size 0, got %d", boolPQ.Size())
		}
	})
}

func TestPriorityQueueEnqueue(t *testing.T) {
	t.Run("enqueue single item", func(t *testing.T) {
		pq := NewPriorityQueue[string]()
		pq.Enqueue("test", 5)

		if pq.Size() != 1 {
			t.Errorf("Expected size 1, got %d", pq.Size())
		}
		if pq.IsEmpty() {
			t.Error("Expected priority queue to not be empty")
		}
		if pq.Peek() != "test" {
			t.Errorf("Expected peek value 'test', got '%s'", pq.Peek())
		}
	})

	t.Run("enqueue multiple items with different priorities", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		// Enqueue in non-priority order
		pq.Enqueue("medium", 5)
		pq.Enqueue("high", 1)
		pq.Enqueue("low", 10)
		pq.Enqueue("highest", 0)

		if pq.Size() != 4 {
			t.Errorf("Expected size 4, got %d", pq.Size())
		}

		// Highest priority (lowest number) should be at top
		if pq.Peek() != "highest" {
			t.Errorf("Expected peek value 'highest', got '%s'", pq.Peek())
		}
	})

	t.Run("enqueue items with same priority", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(1, 5)
		pq.Enqueue(2, 5)
		pq.Enqueue(3, 5)

		if pq.Size() != 3 {
			t.Errorf("Expected size 3, got %d", pq.Size())
		}

		// All have same priority, so any of them could be at top
		peek := pq.Peek()
		if peek != 1 && peek != 2 && peek != 3 {
			t.Errorf("Expected peek to be 1, 2, or 3, got %d", peek)
		}
	})
}

func TestPriorityQueueDequeue(t *testing.T) {
	t.Run("dequeue from empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		value := pq.Dequeue()

		var zero int
		if value != zero {
			t.Errorf("Expected zero value %d when dequeuing from empty queue, got %d", zero, value)
		}
		if pq.Size() != 0 {
			t.Errorf("Expected size to remain 0, got %d", pq.Size())
		}
	})

	t.Run("dequeue single item", func(t *testing.T) {
		pq := NewPriorityQueue[string]()
		pq.Enqueue("test", 1)

		value := pq.Dequeue()
		if value != "test" {
			t.Errorf("Expected dequeued value 'test', got '%s'", value)
		}
		if pq.Size() != 0 {
			t.Errorf("Expected size 0 after dequeue, got %d", pq.Size())
		}
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after dequeue")
		}
	})

	t.Run("dequeue in priority order", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		// Add items with various priorities
		items := []struct {
			value    string
			priority int
		}{
			{"medium", 5},
			{"high", 2},
			{"low", 8},
			{"highest", 1},
			{"lowest", 10},
		}

		for _, item := range items {
			pq.Enqueue(item.value, item.priority)
		}

		// Should dequeue in priority order (lowest priority number first)
		expected := []string{"highest", "high", "medium", "low", "lowest"}
		for i, exp := range expected {
			actual := pq.Dequeue()
			if actual != exp {
				t.Errorf("Dequeue %d: expected '%s', got '%s'", i, exp, actual)
			}
		}

		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after all dequeues")
		}
	})

	t.Run("dequeue with numeric values", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		// Add numbers with priorities
		pq.Enqueue(100, 3)
		pq.Enqueue(200, 1)
		pq.Enqueue(300, 2)

		// Should come out in priority order
		if pq.Dequeue() != 200 {
			t.Error("Expected first dequeue to be 200 (priority 1)")
		}
		if pq.Dequeue() != 300 {
			t.Error("Expected second dequeue to be 300 (priority 2)")
		}
		if pq.Dequeue() != 100 {
			t.Error("Expected third dequeue to be 100 (priority 3)")
		}
	})
}

func TestPriorityQueuePeek(t *testing.T) {
	t.Run("peek empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		value := pq.Peek()

		var zero int
		if value != zero {
			t.Errorf("Expected zero value %d when peeking empty queue, got %d", zero, value)
		}
		if pq.Size() != 0 {
			t.Errorf("Expected size to remain 0 after peek, got %d", pq.Size())
		}
	})

	t.Run("peek does not modify queue", func(t *testing.T) {
		pq := NewPriorityQueue[string]()
		pq.Enqueue("test", 1)

		// Peek multiple times
		for i := 0; i < 5; i++ {
			value := pq.Peek()
			if value != "test" {
				t.Errorf("Peek %d: expected 'test', got '%s'", i, value)
			}
			if pq.Size() != 1 {
				t.Errorf("Peek %d: expected size to remain 1, got %d", i, pq.Size())
			}
		}
	})

	t.Run("peek returns highest priority", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		pq.Enqueue(10, 5)
		if pq.Peek() != 10 {
			t.Error("Expected peek to return 10")
		}

		pq.Enqueue(20, 3)
		if pq.Peek() != 20 {
			t.Error("Expected peek to return 20 (higher priority)")
		}

		pq.Enqueue(30, 7)
		if pq.Peek() != 20 {
			t.Error("Expected peek to still return 20 (highest priority)")
		}

		pq.Enqueue(40, 1)
		if pq.Peek() != 40 {
			t.Error("Expected peek to return 40 (new highest priority)")
		}
	})
}

func TestPriorityQueueSize(t *testing.T) {
	t.Run("size of empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		if pq.Size() != 0 {
			t.Errorf("Expected size 0 for empty queue, got %d", pq.Size())
		}
	})

	t.Run("size changes with enqueue operations", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		for i := 1; i <= 10; i++ {
			pq.Enqueue(i*10, i)
			if pq.Size() != i {
				t.Errorf("After enqueuing %d elements, expected size %d, got %d", i, i, pq.Size())
			}
		}
	})

	t.Run("size changes with dequeue operations", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		// Add 5 elements
		for i := 1; i <= 5; i++ {
			pq.Enqueue(i*10, i)
		}

		// Remove elements and check size
		for i := 4; i >= 0; i-- {
			pq.Dequeue()
			if pq.Size() != i {
				t.Errorf("After dequeue, expected size %d, got %d", i, pq.Size())
			}
		}
	})
}

func TestPriorityQueueIsEmpty(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		if !pq.IsEmpty() {
			t.Error("Expected new queue to be empty")
		}
	})

	t.Run("non-empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(42, 1)
		if pq.IsEmpty() {
			t.Error("Expected queue with elements to not be empty")
		}
	})

	t.Run("empty after clear", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)
		pq.Clear()
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after clear")
		}
	})

	t.Run("empty after dequeuing all", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(42, 1)
		pq.Dequeue()
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after dequeuing all elements")
		}
	})
}

func TestPriorityQueueClear(t *testing.T) {
	t.Run("clear empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Clear()
		if pq.Size() != 0 {
			t.Errorf("Expected size 0 after clearing empty queue, got %d", pq.Size())
		}
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after clear")
		}
	})

	t.Run("clear non-empty queue", func(t *testing.T) {
		pq := NewPriorityQueue[string]()
		pq.Enqueue("a", 1)
		pq.Enqueue("b", 2)
		pq.Enqueue("c", 3)

		pq.Clear()
		if pq.Size() != 0 {
			t.Errorf("Expected size 0 after clear, got %d", pq.Size())
		}
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty after clear")
		}

		// Verify elements are gone
		value := pq.Peek()
		if value != "" {
			t.Errorf("Expected empty string after clear, got '%s'", value)
		}
	})

	t.Run("use after clear", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)
		pq.Clear()

		// Should be able to use normally after clear
		pq.Enqueue(42, 1)
		if pq.Size() != 1 {
			t.Errorf("Expected size 1 after clear and enqueue, got %d", pq.Size())
		}
		if pq.Peek() != 42 {
			t.Errorf("Expected peek 42 after clear and enqueue, got %d", pq.Peek())
		}
	})
}

func TestPriorityQueueToSlice(t *testing.T) {
	t.Run("empty queue to slice", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		slice := pq.ToSlice()
		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(slice))
		}
	})

	t.Run("queue to slice", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		values := []int{10, 20, 30, 40}
		priorities := []int{2, 1, 4, 3}

		for i, v := range values {
			pq.Enqueue(v, priorities[i])
		}

		slice := pq.ToSlice()
		if len(slice) != len(values) {
			t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
		}

		// Convert to map for easier checking since order may vary
		sliceMap := make(map[int]bool)
		for _, v := range slice {
			sliceMap[v] = true
		}

		for _, v := range values {
			if !sliceMap[v] {
				t.Errorf("Expected slice to contain %d", v)
			}
		}
	})

	t.Run("string queue to slice", func(t *testing.T) {
		pq := NewPriorityQueue[string]()
		pq.Enqueue("hello", 2)
		pq.Enqueue("world", 1)
		pq.Enqueue("test", 3)

		slice := pq.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}

		// Check all values are present
		expected := map[string]bool{"hello": true, "world": true, "test": true}
		sliceMap := make(map[string]bool)
		for _, v := range slice {
			sliceMap[v] = true
		}

		if !reflect.DeepEqual(sliceMap, expected) {
			t.Errorf("Expected slice to contain %v, got %v", expected, sliceMap)
		}
	})

	t.Run("slice independence", func(t *testing.T) {
		pq := NewPriorityQueue[int]()
		pq.Enqueue(1, 1)
		pq.Enqueue(2, 2)

		slice := pq.ToSlice()
		// Modify the slice
		slice[0] = 999

		// Original queue should be unchanged
		originalSlice := pq.ToSlice()
		for _, v := range originalSlice {
			if v == 999 {
				t.Error("Modifying returned slice should not affect original queue")
			}
		}
	})
}

func TestPriorityQueueBehavior(t *testing.T) {
	t.Run("priority ordering", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		// Add items in random priority order
		items := []struct {
			value    string
			priority int
		}{
			{"E", 5},
			{"B", 2},
			{"A", 1},
			{"D", 4},
			{"C", 3},
		}

		for _, item := range items {
			pq.Enqueue(item.value, item.priority)
		}

		// Should dequeue in priority order
		expected := []string{"A", "B", "C", "D", "E"}
		var result []string

		for !pq.IsEmpty() {
			result = append(result, pq.Dequeue())
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected priority order %v, got %v", expected, result)
		}
	})

	t.Run("mixed operations", func(t *testing.T) {
		pq := NewPriorityQueue[int]()

		pq.Enqueue(10, 3)
		pq.Enqueue(20, 1)

		if pq.Dequeue() != 20 {
			t.Error("First dequeue should return 20 (priority 1)")
		}

		pq.Enqueue(30, 2)
		pq.Enqueue(40, 1)

		if pq.Dequeue() != 40 {
			t.Error("Second dequeue should return 40 (priority 1)")
		}
		if pq.Dequeue() != 30 {
			t.Error("Third dequeue should return 30 (priority 2)")
		}
		if pq.Dequeue() != 10 {
			t.Error("Fourth dequeue should return 10 (priority 3)")
		}

		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty")
		}
	})

	t.Run("negative priorities", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		pq.Enqueue("normal", 0)
		pq.Enqueue("high", -1)
		pq.Enqueue("higher", -2)

		if pq.Dequeue() != "higher" {
			t.Error("Expected highest priority (most negative) first")
		}
		if pq.Dequeue() != "high" {
			t.Error("Expected second highest priority")
		}
		if pq.Dequeue() != "normal" {
			t.Error("Expected lowest priority last")
		}
	})

	t.Run("custom struct type", func(t *testing.T) {
		type Task struct {
			Name string
			ID   int
		}

		pq := NewPriorityQueue[Task]()

		t1 := Task{Name: "Low", ID: 1}
		t2 := Task{Name: "High", ID: 2}
		t3 := Task{Name: "Medium", ID: 3}

		pq.Enqueue(t1, 3)
		pq.Enqueue(t2, 1)
		pq.Enqueue(t3, 2)

		result1 := pq.Dequeue()
		if result1.Name != "High" || result1.ID != 2 {
			t.Errorf("Expected High(2), got %s(%d)", result1.Name, result1.ID)
		}

		result2 := pq.Peek()
		if result2.Name != "Medium" || result2.ID != 3 {
			t.Errorf("Expected Medium(3), got %s(%d)", result2.Name, result2.ID)
		}
	})
}

func TestPriorityQueueIntegration(t *testing.T) {
	t.Run("complex workflow", func(t *testing.T) {
		pq := NewPriorityQueue[string]()

		// Add initial items
		pq.Enqueue("task1", 5)
		pq.Enqueue("task2", 1)
		pq.Enqueue("task3", 3)

		// Verify initial state
		if pq.Size() != 3 {
			t.Errorf("Expected initial size 3, got %d", pq.Size())
		}
		if pq.Peek() != "task2" {
			t.Errorf("Expected initial peek 'task2', got '%s'", pq.Peek())
		}

		// Process highest priority
		processed := pq.Dequeue()
		if processed != "task2" {
			t.Errorf("Expected dequeue 'task2', got '%s'", processed)
		}

		// Add more items
		pq.Enqueue("urgent", 0)
		pq.Enqueue("normal", 4)

		// Verify state
		if pq.Size() != 4 {
			t.Errorf("Expected size 4, got %d", pq.Size())
		}
		if pq.Peek() != "urgent" {
			t.Errorf("Expected peek 'urgent', got '%s'", pq.Peek())
		}

		// Process remaining in priority order
		expected := []string{"urgent", "task3", "normal", "task1"}
		for i, exp := range expected {
			actual := pq.Dequeue()
			if actual != exp {
				t.Errorf("Process %d: expected '%s', got '%s'", i, exp, actual)
			}
		}

		// Verify empty
		if !pq.IsEmpty() {
			t.Error("Expected queue to be empty")
		}

		// Use after emptying
		pq.Enqueue("new", 1)
		if pq.Size() != 1 {
			t.Errorf("Expected size 1 after re-use, got %d", pq.Size())
		}
	})
}

// Benchmark tests
func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	pq := NewPriorityQueue[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i%1000)
	}
}

func BenchmarkPriorityQueueDequeue(b *testing.B) {
	pq := NewPriorityQueue[int]()
	// Pre-populate the queue
	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i%1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Dequeue()
	}
}

func BenchmarkPriorityQueuePeek(b *testing.B) {
	pq := NewPriorityQueue[int]()
	for i := 0; i < 1000; i++ {
		pq.Enqueue(i, i)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Peek()
	}
}

func BenchmarkPriorityQueueMixedOperations(b *testing.B) {
	pq := NewPriorityQueue[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i, i%100)
		if i%3 == 0 && !pq.IsEmpty() {
			pq.Dequeue()
		}
		if i%5 == 0 {
			pq.Peek()
		}
	}
}
