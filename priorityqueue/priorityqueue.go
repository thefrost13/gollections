package priorityqueue

import "container/heap"

type priorityQueueItem[T any] struct {
	value    T
	priority int
}

type PriorityQueue[T any] []priorityQueueItem[T]

func (pq PriorityQueue[T]) Len() int           { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue[T]) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue[T]) Push(x any) {
	*pq = append(*pq, x.(priorityQueueItem[T]))
}
func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	last := old[n-1]
	*pq = old[0 : n-1]
	return last
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}

func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	item := priorityQueueItem[T]{value: value, priority: priority}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) Dequeue() T {
	var value T
	if pq.Len() != 0 {
		item := heap.Pop(pq)
		return item.(priorityQueueItem[T]).value
	}
	return value
}

func (pq *PriorityQueue[T]) Peek() T {
	var value T
	if pq.Len() != 0 {
		value = (*pq)[0].value
	}
	return value
}

func (pq *PriorityQueue[T]) Size() int {
	return pq.Len()
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq *PriorityQueue[T]) Clear() {
	*pq = (*pq)[:0]
}

func (pq *PriorityQueue[T]) ToSlice() []T {
	slice := make([]T, pq.Len())
	for i, item := range *pq {
		slice[i] = item.value
	}
	return slice
}
