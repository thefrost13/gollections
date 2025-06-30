package gollections

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("create empty stack with nil slice", func(t *testing.T) {
		stack := NewStack[int](nil)
		if stack == nil {
			t.Fatal("NewStack should not return nil")
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size 0, got %d", stack.Size())
		}
		if stack.top != nil {
			t.Error("Expected top to be nil for empty stack")
		}
	})

	t.Run("create empty stack with empty slice", func(t *testing.T) {
		stack := NewStack([]int{})
		if stack == nil {
			t.Fatal("NewStack should not return nil")
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size 0, got %d", stack.Size())
		}
	})

	t.Run("create stack with values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		stack := NewStack(values)
		if stack == nil {
			t.Fatal("NewStack should not return nil")
		}
		if stack.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), stack.Size())
		}
		// Values should be pushed in order, so top should be last value
		if stack.Peek() != 5 {
			t.Errorf("Expected top value 5, got %d", stack.Peek())
		}
	})

	t.Run("create stack with string values", func(t *testing.T) {
		values := []string{"hello", "world", "test"}
		stack := NewStack(values)
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}
		if stack.Peek() != "test" {
			t.Errorf("Expected top value 'test', got '%s'", stack.Peek())
		}
	})
}

func TestStackPush(t *testing.T) {
	t.Run("push to empty stack", func(t *testing.T) {
		stack := NewStack[int](nil)
		stack.Push(42)
		if stack.Size() != 1 {
			t.Errorf("Expected size 1, got %d", stack.Size())
		}
		if stack.Peek() != 42 {
			t.Errorf("Expected top value 42, got %d", stack.Peek())
		}
	})

	t.Run("push multiple values", func(t *testing.T) {
		stack := NewStack[int](nil)
		values := []int{1, 2, 3, 4, 5}

		for i, v := range values {
			stack.Push(v)
			if stack.Size() != i+1 {
				t.Errorf("After pushing %d elements, expected size %d, got %d", i+1, i+1, stack.Size())
			}
			if stack.Peek() != v {
				t.Errorf("After pushing %d, expected top value %d, got %d", v, v, stack.Peek())
			}
		}
	})

	t.Run("push different types", func(t *testing.T) {
		stringStack := NewStack[string](nil)
		stringStack.Push("hello")
		stringStack.Push("world")
		if stringStack.Peek() != "world" {
			t.Errorf("Expected top value 'world', got '%s'", stringStack.Peek())
		}

		boolStack := NewStack[bool](nil)
		boolStack.Push(true)
		boolStack.Push(false)
		if boolStack.Peek() != false {
			t.Errorf("Expected top value false, got %v", boolStack.Peek())
		}
	})
}

func TestStackPop(t *testing.T) {
	t.Run("pop from empty stack", func(t *testing.T) {
		stack := NewStack[int](nil)
		value := stack.Pop()
		var zero int
		if value != zero {
			t.Errorf("Expected zero value %d when popping from empty stack, got %d", zero, value)
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size to remain 0, got %d", stack.Size())
		}
	})

	t.Run("pop single value", func(t *testing.T) {
		stack := NewStack[int](nil)
		stack.Push(42)
		value := stack.Pop()
		if value != 42 {
			t.Errorf("Expected popped value 42, got %d", value)
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size 0 after popping last element, got %d", stack.Size())
		}
	})

	t.Run("pop multiple values LIFO order", func(t *testing.T) {
		stack := NewStack[int](nil)
		values := []int{1, 2, 3, 4, 5}

		// Push values
		for _, v := range values {
			stack.Push(v)
		}

		// Pop values - should come out in reverse order (LIFO)
		for i := len(values) - 1; i >= 0; i-- {
			expected := values[i]
			actual := stack.Pop()
			if actual != expected {
				t.Errorf("Expected popped value %d, got %d", expected, actual)
			}
			if stack.Size() != i {
				t.Errorf("Expected size %d after pop, got %d", i, stack.Size())
			}
		}
	})

	t.Run("pop from stack created with slice", func(t *testing.T) {
		values := []int{1, 2, 3}
		stack := NewStack(values)

		// Should pop in reverse order of the slice
		expected := []int{3, 2, 1}
		for i, exp := range expected {
			actual := stack.Pop()
			if actual != exp {
				t.Errorf("Pop %d: expected %d, got %d", i, exp, actual)
			}
		}
	})
}

func TestStackPeek(t *testing.T) {
	t.Run("peek empty stack", func(t *testing.T) {
		stack := NewStack[int](nil)
		value := stack.Peek()
		var zero int
		if value != zero {
			t.Errorf("Expected zero value %d when peeking empty stack, got %d", zero, value)
		}
		if stack.Size() != 0 {
			t.Errorf("Expected size to remain 0 after peek, got %d", stack.Size())
		}
	})

	t.Run("peek does not modify stack", func(t *testing.T) {
		stack := NewStack[int](nil)
		stack.Push(42)

		// Peek multiple times
		for i := 0; i < 5; i++ {
			value := stack.Peek()
			if value != 42 {
				t.Errorf("Peek %d: expected 42, got %d", i, value)
			}
			if stack.Size() != 1 {
				t.Errorf("Peek %d: expected size to remain 1, got %d", i, stack.Size())
			}
		}
	})

	t.Run("peek after push and pop operations", func(t *testing.T) {
		stack := NewStack[string](nil)

		stack.Push("first")
		if stack.Peek() != "first" {
			t.Errorf("Expected 'first', got '%s'", stack.Peek())
		}

		stack.Push("second")
		if stack.Peek() != "second" {
			t.Errorf("Expected 'second', got '%s'", stack.Peek())
		}

		stack.Pop()
		if stack.Peek() != "first" {
			t.Errorf("Expected 'first' after pop, got '%s'", stack.Peek())
		}
	})
}

func TestStackSize(t *testing.T) {
	t.Run("size of empty stack", func(t *testing.T) {
		stack := NewStack[int](nil)
		if stack.Size() != 0 {
			t.Errorf("Expected size 0 for empty stack, got %d", stack.Size())
		}
	})

	t.Run("size changes with push operations", func(t *testing.T) {
		stack := NewStack[int](nil)

		for i := 1; i <= 10; i++ {
			stack.Push(i)
			if stack.Size() != i {
				t.Errorf("After pushing %d elements, expected size %d, got %d", i, i, stack.Size())
			}
		}
	})

	t.Run("size changes with pop operations", func(t *testing.T) {
		stack := NewStack([]int{1, 2, 3, 4, 5})
		initialSize := stack.Size()

		for i := 0; i < initialSize; i++ {
			stack.Pop()
			expectedSize := initialSize - i - 1
			if stack.Size() != expectedSize {
				t.Errorf("After %d pops, expected size %d, got %d", i+1, expectedSize, stack.Size())
			}
		}
	})

	t.Run("size with mixed operations", func(t *testing.T) {
		stack := NewStack[int](nil)

		// Push 3 elements
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)
		if stack.Size() != 3 {
			t.Errorf("Expected size 3, got %d", stack.Size())
		}

		// Pop 1 element
		stack.Pop()
		if stack.Size() != 2 {
			t.Errorf("Expected size 2, got %d", stack.Size())
		}

		// Push 2 more
		stack.Push(4)
		stack.Push(5)
		if stack.Size() != 4 {
			t.Errorf("Expected size 4, got %d", stack.Size())
		}
	})
}

func TestStackToSlice(t *testing.T) {
	t.Run("empty stack to slice", func(t *testing.T) {
		stack := NewStack[int](nil)
		slice := stack.ToSlice()
		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(slice))
		}
	})

	t.Run("stack to slice maintains LIFO order", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		stack := NewStack[int](values)
		slice := stack.ToSlice()

		// ToSlice should return elements in LIFO order (top to bottom)
		expected := []int{5, 4, 3, 2, 1}

		if len(slice) != len(values) {
			t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
		}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v (LIFO order), got %v", expected, slice)
		}
	})

	t.Run("string stack to slice", func(t *testing.T) {
		values := []string{"first", "second", "third"}
		stack := NewStack[string](values)
		slice := stack.ToSlice()

		// ToSlice should return elements in LIFO order
		expected := []string{"third", "second", "first"}

		if len(slice) != len(values) {
			t.Errorf("Expected slice length %d, got %d", len(values), len(slice))
		}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})

	t.Run("slice independence", func(t *testing.T) {
		stack := NewStack[int]([]int{1, 2, 3})
		slice := stack.ToSlice()

		// Modify the slice
		slice[0] = 999

		// Original stack should be unchanged (top should still be 3)
		if stack.Peek() != 3 {
			t.Error("Modifying returned slice should not affect original stack")
		}
	})

	t.Run("partial stack to slice", func(t *testing.T) {
		stack := NewStack[int]([]int{1, 2, 3, 4, 5})

		// Pop some elements (removes 5 and 4)
		stack.Pop()
		stack.Pop()

		slice := stack.ToSlice()
		expected := []int{3, 2, 1} // LIFO order from top to bottom

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})

	t.Run("single element stack to slice", func(t *testing.T) {
		stack := NewStack[int](nil)
		stack.Push(42)

		slice := stack.ToSlice()
		expected := []int{42}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})
}

func TestStackIntegration(t *testing.T) {
	t.Run("complex workflow", func(t *testing.T) {
		// Create stack with initial values
		stack := NewStack([]int{10, 20, 30})

		// Verify initial state
		if stack.Size() != 3 {
			t.Errorf("Expected initial size 3, got %d", stack.Size())
		}
		if stack.Peek() != 30 {
			t.Errorf("Expected initial top 30, got %d", stack.Peek())
		}

		// Pop one value
		value := stack.Pop()
		if value != 30 {
			t.Errorf("Expected popped value 30, got %d", value)
		}

		// Push new values
		stack.Push(40)
		stack.Push(50)

		// Verify state
		if stack.Size() != 4 {
			t.Errorf("Expected size 4, got %d", stack.Size())
		}
		if stack.Peek() != 50 {
			t.Errorf("Expected top 50, got %d", stack.Peek())
		}

		// Pop all remaining values and verify order
		expected := []int{50, 40, 20, 10}
		for i, exp := range expected {
			actual := stack.Pop()
			if actual != exp {
				t.Errorf("Pop %d: expected %d, got %d", i, exp, actual)
			}
		}

		// Verify empty
		if stack.Size() != 0 {
			t.Errorf("Expected final size 0, got %d", stack.Size())
		}
	})
}

// Benchmark tests
func BenchmarkStackPush(b *testing.B) {
	stack := NewStack[int](nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := NewStack[int](nil)
	// Pre-populate the stack
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}

func BenchmarkStackPeek(b *testing.B) {
	stack := NewStack([]int{1, 2, 3, 4, 5})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stack.Peek()
	}
}
