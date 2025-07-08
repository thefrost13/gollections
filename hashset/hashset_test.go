package hashset

import (
	"reflect"
	"sort"
	"testing"
)

func TestNewHashSet(t *testing.T) {
	t.Run("create empty hashset with nil slice", func(t *testing.T) {
		set := NewHashSet[int](nil)
		if set == nil {
			t.Fatal("NewHashSet should not return nil")
		}
		if set.Size() != 0 {
			t.Errorf("Expected size 0, got %d", set.Size())
		}
		if !set.IsEmpty() {
			t.Error("Expected empty set")
		}
	})

	t.Run("create empty hashset with empty slice", func(t *testing.T) {
		set := NewHashSet([]int{})
		if set == nil {
			t.Fatal("NewHashSet should not return nil")
		}
		if set.Size() != 0 {
			t.Errorf("Expected size 0, got %d", set.Size())
		}
		if !set.IsEmpty() {
			t.Error("Expected empty set")
		}
	})

	t.Run("create hashset with values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		set := NewHashSet(values)
		if set.Size() != len(values) {
			t.Errorf("Expected size %d, got %d", len(values), set.Size())
		}
		for _, v := range values {
			if !set.Contains(v) {
				t.Errorf("Expected set to contain %d", v)
			}
		}
	})

	t.Run("create hashset with duplicate values", func(t *testing.T) {
		values := []int{1, 2, 2, 3, 3, 3, 4}
		set := NewHashSet(values)
		expectedSize := 4 // unique values: 1, 2, 3, 4
		if set.Size() != expectedSize {
			t.Errorf("Expected size %d (duplicates removed), got %d", expectedSize, set.Size())
		}
		for _, v := range []int{1, 2, 3, 4} {
			if !set.Contains(v) {
				t.Errorf("Expected set to contain %d", v)
			}
		}
	})

	t.Run("create hashset with string values", func(t *testing.T) {
		values := []string{"hello", "world", "test", "hello"}
		set := NewHashSet(values)
		expectedSize := 3 // unique values: "hello", "world", "test"
		if set.Size() != expectedSize {
			t.Errorf("Expected size %d, got %d", expectedSize, set.Size())
		}
		for _, v := range []string{"hello", "world", "test"} {
			if !set.Contains(v) {
				t.Errorf("Expected set to contain '%s'", v)
			}
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("add to empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		set.Add(42)
		if set.Size() != 1 {
			t.Errorf("Expected size 1, got %d", set.Size())
		}
		if !set.Contains(42) {
			t.Error("Expected set to contain 42")
		}
		if set.IsEmpty() {
			t.Error("Expected set to not be empty")
		}
	})

	t.Run("add multiple values", func(t *testing.T) {
		set := NewHashSet[int](nil)
		values := []int{1, 2, 3, 4, 5}

		for i, v := range values {
			set.Add(v)
			if set.Size() != i+1 {
				t.Errorf("After adding %d elements, expected size %d, got %d", i+1, i+1, set.Size())
			}
			if !set.Contains(v) {
				t.Errorf("Expected set to contain %d", v)
			}
		}
	})

	t.Run("add duplicate values", func(t *testing.T) {
		set := NewHashSet[int](nil)
		set.Add(42)
		set.Add(42)
		set.Add(42)

		if set.Size() != 1 {
			t.Errorf("Expected size 1 after adding duplicates, got %d", set.Size())
		}
		if !set.Contains(42) {
			t.Error("Expected set to contain 42")
		}
	})

	t.Run("add different types", func(t *testing.T) {
		stringSet := NewHashSet[string](nil)
		stringSet.Add("hello")
		stringSet.Add("world")
		if stringSet.Size() != 2 {
			t.Errorf("Expected string set size 2, got %d", stringSet.Size())
		}

		boolSet := NewHashSet[bool](nil)
		boolSet.Add(true)
		boolSet.Add(false)
		boolSet.Add(true) // duplicate
		if boolSet.Size() != 2 {
			t.Errorf("Expected bool set size 2, got %d", boolSet.Size())
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("remove from empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		set.Remove(42)
		if set.Size() != 0 {
			t.Errorf("Expected size to remain 0, got %d", set.Size())
		}
	})

	t.Run("remove existing value", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		set.Remove(2)
		if set.Size() != 2 {
			t.Errorf("Expected size 2, got %d", set.Size())
		}
		if set.Contains(2) {
			t.Error("Expected set to not contain 2")
		}
		if !set.Contains(1) || !set.Contains(3) {
			t.Error("Expected set to still contain 1 and 3")
		}
	})

	t.Run("remove non-existing value", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		originalSize := set.Size()
		set.Remove(99)
		if set.Size() != originalSize {
			t.Errorf("Expected size to remain %d, got %d", originalSize, set.Size())
		}
	})

	t.Run("remove all values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		set := NewHashSet(values)

		for i, v := range values {
			set.Remove(v)
			expectedSize := len(values) - i - 1
			if set.Size() != expectedSize {
				t.Errorf("After removing %d elements, expected size %d, got %d", i+1, expectedSize, set.Size())
			}
			if set.Contains(v) {
				t.Errorf("Expected set to not contain %d after removal", v)
			}
		}

		if !set.IsEmpty() {
			t.Error("Expected set to be empty after removing all elements")
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("contains in empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		if set.Contains(42) {
			t.Error("Expected empty set to not contain any value")
		}
	})

	t.Run("contains existing values", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		set := NewHashSet(values)

		for _, v := range values {
			if !set.Contains(v) {
				t.Errorf("Expected set to contain %d", v)
			}
		}
	})

	t.Run("contains non-existing values", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		nonExisting := []int{0, 4, 99, -1}

		for _, v := range nonExisting {
			if set.Contains(v) {
				t.Errorf("Expected set to not contain %d", v)
			}
		}
	})

	t.Run("contains after add and remove", func(t *testing.T) {
		set := NewHashSet[string](nil)

		set.Add("test")
		if !set.Contains("test") {
			t.Error("Expected set to contain 'test' after adding")
		}

		set.Remove("test")
		if set.Contains("test") {
			t.Error("Expected set to not contain 'test' after removing")
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("size of empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		if set.Size() != 0 {
			t.Errorf("Expected size 0 for empty set, got %d", set.Size())
		}
	})

	t.Run("size changes with add operations", func(t *testing.T) {
		set := NewHashSet[int](nil)

		for i := 1; i <= 10; i++ {
			set.Add(i)
			if set.Size() != i {
				t.Errorf("After adding %d elements, expected size %d, got %d", i, i, set.Size())
			}
		}
	})

	t.Run("size changes with remove operations", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3, 4, 5})
		initialSize := set.Size()

		for i := 0; i < initialSize; i++ {
			set.Remove(i + 1)
			expectedSize := initialSize - i - 1
			if set.Size() != expectedSize {
				t.Errorf("After %d removals, expected size %d, got %d", i+1, expectedSize, set.Size())
			}
		}
	})

	t.Run("size with duplicates", func(t *testing.T) {
		set := NewHashSet[int](nil)
		set.Add(1)
		set.Add(1)
		set.Add(1)
		if set.Size() != 1 {
			t.Errorf("Expected size 1 with duplicates, got %d", set.Size())
		}
	})
}

func TestIsEmpty(t *testing.T) {
	t.Run("empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		if !set.IsEmpty() {
			t.Error("Expected new set to be empty")
		}
	})

	t.Run("non-empty hashset", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		if set.IsEmpty() {
			t.Error("Expected set with values to not be empty")
		}
	})

	t.Run("empty after clear", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		set.Clear()
		if !set.IsEmpty() {
			t.Error("Expected set to be empty after clear")
		}
	})

	t.Run("empty after removing all", func(t *testing.T) {
		set := NewHashSet([]int{42})
		set.Remove(42)
		if !set.IsEmpty() {
			t.Error("Expected set to be empty after removing all elements")
		}
	})
}

func TestClear(t *testing.T) {
	t.Run("clear empty hashset", func(t *testing.T) {
		set := NewHashSet[int](nil)
		set.Clear()
		if set.Size() != 0 {
			t.Errorf("Expected size 0 after clearing empty set, got %d", set.Size())
		}
		if !set.IsEmpty() {
			t.Error("Expected set to be empty after clear")
		}
	})

	t.Run("clear non-empty hashset", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3, 4, 5})
		set.Clear()
		if set.Size() != 0 {
			t.Errorf("Expected size 0 after clear, got %d", set.Size())
		}
		if !set.IsEmpty() {
			t.Error("Expected set to be empty after clear")
		}

		// Verify no elements remain
		for i := 1; i <= 5; i++ {
			if set.Contains(i) {
				t.Errorf("Expected set to not contain %d after clear", i)
			}
		}
	})

	t.Run("add after clear", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		set.Clear()
		set.Add(42)

		if set.Size() != 1 {
			t.Errorf("Expected size 1 after clear and add, got %d", set.Size())
		}
		if !set.Contains(42) {
			t.Error("Expected set to contain 42 after clear and add")
		}
	})
}

func TestToSlice(t *testing.T) {
	t.Run("empty hashset to slice", func(t *testing.T) {
		set := NewHashSet[int](nil)
		slice := set.ToSlice()
		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(slice))
		}
	})

	t.Run("hashset to slice", func(t *testing.T) {
		values := []int{3, 1, 4, 1, 5, 9, 2, 6} // includes duplicates
		set := NewHashSet(values)
		slice := set.ToSlice()

		// Sort both slices for comparison since order is not guaranteed
		expectedUnique := []int{1, 2, 3, 4, 5, 6, 9}
		sort.Ints(slice)

		if !reflect.DeepEqual(slice, expectedUnique) {
			t.Errorf("Expected slice %v, got %v", expectedUnique, slice)
		}
	})

	t.Run("string hashset to slice", func(t *testing.T) {
		values := []string{"hello", "world", "test", "hello"}
		set := NewHashSet(values)
		slice := set.ToSlice()

		// Convert to map for easier checking since order is not guaranteed
		sliceMap := make(map[string]bool)
		for _, v := range slice {
			sliceMap[v] = true
		}

		expected := map[string]bool{"hello": true, "world": true, "test": true}
		if !reflect.DeepEqual(sliceMap, expected) {
			t.Errorf("Expected slice to contain %v, got %v", expected, sliceMap)
		}
	})

	t.Run("slice independence", func(t *testing.T) {
		set := NewHashSet([]int{1, 2, 3})
		slice := set.ToSlice()

		// Modify the slice
		slice[0] = 999

		// Original set should be unchanged
		if set.Contains(999) {
			t.Error("Modifying returned slice should not affect original set")
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("empty hashsets equal", func(t *testing.T) {
		set1 := NewHashSet[int](nil)
		set2 := NewHashSet[int](nil)
		if !set1.Equals(set2) {
			t.Error("Expected empty sets to be equal")
		}
	})

	t.Run("identical hashsets equal", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5}
		set1 := NewHashSet(values)
		set2 := NewHashSet(values)
		if !set1.Equals(set2) {
			t.Error("Expected identical sets to be equal")
		}
	})

	t.Run("different order same content equal", func(t *testing.T) {
		set1 := NewHashSet([]int{1, 2, 3})
		set2 := NewHashSet([]int{3, 1, 2})
		if !set1.Equals(set2) {
			t.Error("Expected sets with same content in different order to be equal")
		}
	})

	t.Run("different size not equal", func(t *testing.T) {
		set1 := NewHashSet([]int{1, 2, 3})
		set2 := NewHashSet([]int{1, 2})
		if set1.Equals(set2) {
			t.Error("Expected sets with different sizes to not be equal")
		}
	})

	t.Run("different content not equal", func(t *testing.T) {
		set1 := NewHashSet([]int{1, 2, 3})
		set2 := NewHashSet([]int{4, 5, 6})
		if set1.Equals(set2) {
			t.Error("Expected sets with different content to not be equal")
		}
	})

	t.Run("partially different content not equal", func(t *testing.T) {
		set1 := NewHashSet([]int{1, 2, 3})
		set2 := NewHashSet([]int{1, 2, 4})
		if set1.Equals(set2) {
			t.Error("Expected sets with partially different content to not be equal")
		}
	})

	t.Run("string hashsets equality", func(t *testing.T) {
		set1 := NewHashSet([]string{"hello", "world"})
		set2 := NewHashSet([]string{"world", "hello"})
		set3 := NewHashSet([]string{"hello", "test"})

		if !set1.Equals(set2) {
			t.Error("Expected string sets with same content to be equal")
		}
		if set1.Equals(set3) {
			t.Error("Expected string sets with different content to not be equal")
		}
	})
}

func TestHashSetIntegration(t *testing.T) {
	t.Run("complex workflow", func(t *testing.T) {
		// Create set with initial values
		set := NewHashSet([]int{1, 2, 3, 2, 4}) // duplicates should be ignored

		// Verify initial state
		if set.Size() != 4 {
			t.Errorf("Expected initial size 4, got %d", set.Size())
		}

		// Add new values
		set.Add(5)
		set.Add(6)
		set.Add(1) // duplicate

		if set.Size() != 6 {
			t.Errorf("Expected size 6 after adds, got %d", set.Size())
		}

		// Remove some values
		set.Remove(2)
		set.Remove(6)
		set.Remove(99) // non-existing

		if set.Size() != 4 {
			t.Errorf("Expected size 4 after removes, got %d", set.Size())
		}

		// Verify final content
		expected := []int{1, 3, 4, 5}
		for _, v := range expected {
			if !set.Contains(v) {
				t.Errorf("Expected set to contain %d", v)
			}
		}

		// Convert to slice and verify
		slice := set.ToSlice()
		sort.Ints(slice)
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}

		// Clear and verify
		set.Clear()
		if !set.IsEmpty() {
			t.Error("Expected set to be empty after clear")
		}
	})

	t.Run("set operations with equals", func(t *testing.T) {
		set1 := NewHashSet([]int{1, 2, 3})
		set2 := NewHashSet[int](nil)

		// Build set2 to match set1
		set2.Add(3)
		set2.Add(1)
		set2.Add(2)

		if !set1.Equals(set2) {
			t.Error("Expected manually built sets to be equal")
		}

		// Modify and check inequality
		set2.Add(4)
		if set1.Equals(set2) {
			t.Error("Expected sets to be unequal after modification")
		}
	})
}

// Benchmark tests
func BenchmarkHashSetAdd(b *testing.B) {
	set := NewHashSet[int](nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		set.Add(i)
	}
}

func BenchmarkHashSetContains(b *testing.B) {
	set := NewHashSet[int](nil)
	// Pre-populate the set
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Contains(i % 1000)
	}
}

func BenchmarkHashSetRemove(b *testing.B) {
	set := NewHashSet[int](nil)
	// Pre-populate the set
	for i := 0; i < b.N; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Remove(i)
	}
}

func BenchmarkHashSetToSlice(b *testing.B) {
	set := NewHashSet[int](nil)
	// Pre-populate the set
	for i := 0; i < 1000; i++ {
		set.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.ToSlice()
	}
}
