package orderedhashmap

import "github.com/thefrost13/gollections/node"

type OrderedHashMap[K comparable, V any] struct {
	first *node.LinkedNode[KVPair[K, V]]
	last  *node.LinkedNode[KVPair[K, V]]
	size  int
	m     map[K]*node.LinkedNode[KVPair[K, V]]
}
type KVPair[K comparable, V any] struct {
	Key   K
	Value V
}

func NewOrderedHashMap[K comparable, V any]() *OrderedHashMap[K, V] {
	return &OrderedHashMap[K, V]{
		m: make(map[K]*node.LinkedNode[KVPair[K, V]]),
	}
}
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

func (ohm *OrderedHashMap[K, V]) Get(key K) (V, bool) {
	if node, exists := ohm.m[key]; exists {
		return node.Value.Value, true
	}
	var zeroValue V
	return zeroValue, false
}

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
			ohm.last = nil
		}
		ohm.size--
	}
}

func (ohm *OrderedHashMap[K, V]) Size() int {
	return ohm.size
}

func (ohm *OrderedHashMap[K, V]) IsEmpty() bool {
	return ohm.size == 0
}

func (ohm *OrderedHashMap[K, V]) Keys() []K {
	keys := make([]K, 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		keys = append(keys, node.Value.Key)
	}
	return keys
}

func (ohm *OrderedHashMap[K, V]) Values() []V {
	values := make([]V, 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		values = append(values, node.Value.Value)
	}
	return values
}

func (ohm *OrderedHashMap[K, V]) ToSlice() []KVPair[K, V] {
	slice := make([]KVPair[K, V], 0, ohm.size)
	for node := ohm.first; node != nil; node = node.Next {
		slice = append(slice, node.Value)
	}
	return slice
}
