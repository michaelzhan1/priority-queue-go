package priorityqueue

import "container/heap"

// PriorityQueue is a minheap priority queue.
type PriorityQueue[T any] struct {
	heap *heapImpl[T]
}

// NewPriorityQueue returns a new priority queue of generic type T.
func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: &heapImpl[T]{
			items: []T{},
			less:  less,
		},
	}
}

// Push pushes a new item of type T to the priority queue.
func (pq *PriorityQueue[T]) Push(item T) {
	heap.Push(pq.heap, item)
}

// Pop removes the top item from the priority queue, and returns true if successful.
// It will return false with the zero value of T if the priority queue is empty.
func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	item := heap.Pop(pq.heap).(T)
	return item, true
}

// Peek returns the current top value of the priority queue without removing it.
// It will return true if successful, otherwise false with the zero value of T if the priority queue is empty.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.heap.items[0], true
}

// IsEmpty returns a bool representing whether the priority queue has items or not.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return pq.heap.Len() == 0
}
