package orderedhashmap

import (
	"reflect"
	"testing"
)

func TestNewOrderedHashMap(t *testing.T) {
	t.Run("create empty ordered hash map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		if ohm == nil {
			t.Fatal("NewOrderedHashMap should not return nil")
		}
		if ohm.Size() != 0 {
			t.Errorf("Expected size 0, got %d", ohm.Size())
		}
		if !ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return true for new map")
		}
		if ohm.first != nil {
			t.Error("Expected first to be nil for empty map")
		}
		if ohm.last != nil {
			t.Error("Expected last to be nil for empty map")
		}
	})
}

func TestOrderedHashMap_Set(t *testing.T) {
	t.Run("set single key-value pair", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)

		if ohm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", ohm.Size())
		}
		if ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return false after setting value")
		}
		if ohm.first == nil {
			t.Error("Expected first to be set")
		}
		if ohm.last == nil {
			t.Error("Expected last to be set")
		}
		if ohm.first != ohm.last {
			t.Error("Expected first and last to be same for single element")
		}
	})

	t.Run("set multiple key-value pairs", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key3", 300)

		if ohm.Size() != 3 {
			t.Errorf("Expected size 3, got %d", ohm.Size())
		}
		if ohm.first == ohm.last {
			t.Error("Expected first and last to be different for multiple elements")
		}
	})

	t.Run("update existing key", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key1", 150) // Update existing key

		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}
		value, exists := ohm.Get("key1")
		if !exists {
			t.Error("Expected key1 to exist")
		}
		if value != 150 {
			t.Errorf("Expected value 150, got %d", value)
		}
	})

	t.Run("set with different types", func(t *testing.T) {
		ohm := NewOrderedHashMap[int, string]()
		ohm.Set(1, "one")
		ohm.Set(2, "two")

		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}
		value, exists := ohm.Get(1)
		if !exists || value != "one" {
			t.Errorf("Expected 'one', got %s", value)
		}
	})
}

func TestOrderedHashMap_Get(t *testing.T) {
	t.Run("get existing key", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)

		value, exists := ohm.Get("key1")
		if !exists {
			t.Error("Expected key1 to exist")
		}
		if value != 100 {
			t.Errorf("Expected value 100, got %d", value)
		}

		value, exists = ohm.Get("key2")
		if !exists {
			t.Error("Expected key2 to exist")
		}
		if value != 200 {
			t.Errorf("Expected value 200, got %d", value)
		}
	})

	t.Run("get non-existing key", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)

		value, exists := ohm.Get("nonexistent")
		if exists {
			t.Error("Expected key to not exist")
		}
		if value != 0 {
			t.Errorf("Expected zero value 0, got %d", value)
		}
	})

	t.Run("get from empty map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()

		value, exists := ohm.Get("key1")
		if exists {
			t.Error("Expected key to not exist in empty map")
		}
		if value != 0 {
			t.Errorf("Expected zero value 0, got %d", value)
		}
	})
}

func TestOrderedHashMap_Delete(t *testing.T) {
	t.Run("delete existing key from single element map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)

		ohm.Delete("key1")

		if ohm.Size() != 0 {
			t.Errorf("Expected size 0, got %d", ohm.Size())
		}
		if !ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return true after deletion")
		}
		if ohm.first != nil {
			t.Error("Expected first to be nil after deletion")
		}
		if ohm.last != nil {
			t.Error("Expected last to be nil after deletion")
		}

		_, exists := ohm.Get("key1")
		if exists {
			t.Error("Expected key1 to not exist after deletion")
		}
	})

	t.Run("delete first element from multiple element map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key3", 300)

		ohm.Delete("key1")

		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}

		_, exists := ohm.Get("key1")
		if exists {
			t.Error("Expected key1 to not exist after deletion")
		}

		// Check that order is preserved
		keys := ohm.Keys()
		expected := []string{"key2", "key3"}
		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})

	t.Run("delete middle element from multiple element map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key3", 300)

		ohm.Delete("key2")

		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}

		_, exists := ohm.Get("key2")
		if exists {
			t.Error("Expected key2 to not exist after deletion")
		}

		// Check that order is preserved
		keys := ohm.Keys()
		expected := []string{"key1", "key3"}
		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})

	t.Run("delete last element from multiple element map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key3", 300)

		ohm.Delete("key3")

		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}

		_, exists := ohm.Get("key3")
		if exists {
			t.Error("Expected key3 to not exist after deletion")
		}

		// Check that order is preserved
		keys := ohm.Keys()
		expected := []string{"key1", "key2"}
		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})

	t.Run("delete last element updates last pointer correctly", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)
		ohm.Set("key2", 200)
		ohm.Set("key3", 300)

		// Delete the last element
		ohm.Delete("key3")

		// Verify last pointer is updated correctly
		if ohm.last == nil {
			t.Error("Expected last to not be nil after deleting last element")
		}
		if ohm.last.Value.Key != "key2" {
			t.Errorf("Expected last element to be key2, got %s", ohm.last.Value.Key)
		}

		// Verify we can still add elements correctly
		ohm.Set("key4", 400)
		keys := ohm.Keys()
		expected := []string{"key1", "key2", "key4"}
		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})

	t.Run("delete non-existing key", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("key1", 100)

		originalSize := ohm.Size()
		ohm.Delete("nonexistent")

		if ohm.Size() != originalSize {
			t.Errorf("Expected size to remain %d, got %d", originalSize, ohm.Size())
		}
	})

	t.Run("delete from empty map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()

		ohm.Delete("key1")

		if ohm.Size() != 0 {
			t.Errorf("Expected size 0, got %d", ohm.Size())
		}
		if !ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return true")
		}
	})
}

func TestOrderedHashMap_Size(t *testing.T) {
	t.Run("size operations", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()

		if ohm.Size() != 0 {
			t.Errorf("Expected size 0, got %d", ohm.Size())
		}

		ohm.Set("key1", 100)
		if ohm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", ohm.Size())
		}

		ohm.Set("key2", 200)
		if ohm.Size() != 2 {
			t.Errorf("Expected size 2, got %d", ohm.Size())
		}

		ohm.Delete("key1")
		if ohm.Size() != 1 {
			t.Errorf("Expected size 1, got %d", ohm.Size())
		}

		ohm.Delete("key2")
		if ohm.Size() != 0 {
			t.Errorf("Expected size 0, got %d", ohm.Size())
		}
	})
}

func TestOrderedHashMap_IsEmpty(t *testing.T) {
	t.Run("isEmpty operations", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()

		if !ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return true for new map")
		}

		ohm.Set("key1", 100)
		if ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return false after adding element")
		}

		ohm.Delete("key1")
		if !ohm.IsEmpty() {
			t.Error("Expected IsEmpty to return true after removing all elements")
		}
	})
}

func TestOrderedHashMap_Keys(t *testing.T) {
	t.Run("keys maintain insertion order", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 1)
		ohm.Set("second", 2)
		ohm.Set("third", 3)

		keys := ohm.Keys()
		expected := []string{"first", "second", "third"}

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})

	t.Run("keys from empty map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		keys := ohm.Keys()

		if len(keys) != 0 {
			t.Errorf("Expected empty keys slice, got %v", keys)
		}
	})

	t.Run("keys with updates maintain order", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 1)
		ohm.Set("second", 2)
		ohm.Set("third", 3)
		ohm.Set("second", 22) // Update existing key

		keys := ohm.Keys()
		expected := []string{"first", "second", "third"}

		if !reflect.DeepEqual(keys, expected) {
			t.Errorf("Expected keys %v, got %v", expected, keys)
		}
	})
}

func TestOrderedHashMap_Values(t *testing.T) {
	t.Run("values maintain insertion order", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 10)
		ohm.Set("second", 20)
		ohm.Set("third", 30)

		values := ohm.Values()
		expected := []int{10, 20, 30}

		if !reflect.DeepEqual(values, expected) {
			t.Errorf("Expected values %v, got %v", expected, values)
		}
	})

	t.Run("values from empty map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		values := ohm.Values()

		if len(values) != 0 {
			t.Errorf("Expected empty values slice, got %v", values)
		}
	})

	t.Run("values with updates", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 10)
		ohm.Set("second", 20)
		ohm.Set("third", 30)
		ohm.Set("second", 25) // Update existing key

		values := ohm.Values()
		expected := []int{10, 25, 30}

		if !reflect.DeepEqual(values, expected) {
			t.Errorf("Expected values %v, got %v", expected, values)
		}
	})
}

func TestOrderedHashMap_ToSlice(t *testing.T) {
	t.Run("toSlice maintains insertion order", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 10)
		ohm.Set("second", 20)
		ohm.Set("third", 30)

		slice := ohm.ToSlice()
		expected := []KVPair[string, int]{
			{Key: "first", Value: 10},
			{Key: "second", Value: 20},
			{Key: "third", Value: 30},
		}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})

	t.Run("toSlice from empty map", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		slice := ohm.ToSlice()

		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got %v", slice)
		}
	})

	t.Run("toSlice with updates", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("first", 10)
		ohm.Set("second", 20)
		ohm.Set("first", 15) // Update existing key

		slice := ohm.ToSlice()
		expected := []KVPair[string, int]{
			{Key: "first", Value: 15},
			{Key: "second", Value: 20},
		}

		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected slice %v, got %v", expected, slice)
		}
	})
}

func TestOrderedHashMap_ComplexOperations(t *testing.T) {
	t.Run("mixed operations maintain order", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()

		// Add elements
		ohm.Set("a", 1)
		ohm.Set("b", 2)
		ohm.Set("c", 3)
		ohm.Set("d", 4)

		// Delete middle element
		ohm.Delete("b")

		// Add new element
		ohm.Set("e", 5)

		// Update existing element
		ohm.Set("c", 33)

		// Verify order and values
		keys := ohm.Keys()
		expectedKeys := []string{"a", "c", "d", "e"}
		if !reflect.DeepEqual(keys, expectedKeys) {
			t.Errorf("Expected keys %v, got %v", expectedKeys, keys)
		}

		values := ohm.Values()
		expectedValues := []int{1, 33, 4, 5}
		if !reflect.DeepEqual(values, expectedValues) {
			t.Errorf("Expected values %v, got %v", expectedValues, values)
		}

		if ohm.Size() != 4 {
			t.Errorf("Expected size 4, got %d", ohm.Size())
		}
	})

	t.Run("stress test with many operations", func(t *testing.T) {
		ohm := NewOrderedHashMap[int, string]()

		// Add many elements
		for i := 0; i < 100; i++ {
			ohm.Set(i, "value"+string(rune('0'+i%10)))
		}

		if ohm.Size() != 100 {
			t.Errorf("Expected size 100, got %d", ohm.Size())
		}

		// Delete every other element
		for i := 0; i < 100; i += 2 {
			ohm.Delete(i)
		}

		if ohm.Size() != 50 {
			t.Errorf("Expected size 50, got %d", ohm.Size())
		}

		// Verify remaining elements are correct
		keys := ohm.Keys()
		for i, key := range keys {
			expected := 1 + i*2
			if key != expected {
				t.Errorf("Expected key %d at position %d, got %d", expected, i, key)
			}
		}
	})
}

func TestOrderedHashMap_EdgeCases(t *testing.T) {
	t.Run("zero values", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("zero", 0)

		value, exists := ohm.Get("zero")
		if !exists {
			t.Error("Expected key with zero value to exist")
		}
		if value != 0 {
			t.Errorf("Expected value 0, got %d", value)
		}
	})

	t.Run("empty string keys", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, int]()
		ohm.Set("", 42)

		value, exists := ohm.Get("")
		if !exists {
			t.Error("Expected empty string key to exist")
		}
		if value != 42 {
			t.Errorf("Expected value 42, got %d", value)
		}
	})

	t.Run("nil pointer values", func(t *testing.T) {
		ohm := NewOrderedHashMap[string, *int]()
		ohm.Set("nil", nil)

		value, exists := ohm.Get("nil")
		if !exists {
			t.Error("Expected key with nil value to exist")
		}
		if value != nil {
			t.Errorf("Expected value nil, got %v", value)
		}
	})
}
