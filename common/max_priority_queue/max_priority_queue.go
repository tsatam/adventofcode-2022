package max_priority_queue

import (
	"container/heap"
)

//
// Taken from Go's examples in container/heap
//

// An Item is something we manage in a priority queue.
type Item[T any] struct {
	value    T   // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A MaxPriorityQueue implements heap.Interface and holds Items.
type MaxPriorityQueue[T any] []*Item[T]

func (pq MaxPriorityQueue[T]) Len() int { return len(pq) }

func (pq MaxPriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq MaxPriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MaxPriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MaxPriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func New[T any](length, capacity int) MaxPriorityQueue[T] {
	queue := make(MaxPriorityQueue[T], length, capacity)
	heap.Init(&queue)
	return queue
}

func (pq *MaxPriorityQueue[T]) PopMax() T {
	return heap.Pop(pq).(*Item[T]).value
}

func (pq *MaxPriorityQueue[T]) Empty() bool {
	return len(*pq) == 0
}

func (pq *MaxPriorityQueue[T]) AddAtPriority(value T, priority int) {
	heap.Push(pq, &Item[T]{
		value:    value,
		priority: priority,
		index:    -1,
	})
}

// func (pq *MaxPriorityQueue[T]) SetPriority(value T, priority int) {
// 	var index int
// 	for i, item := range *pq {
// 		if item.value == value {
// 			index = i
// 			item.priority = priority
// 			break
// 		}
// 	}

// 	heap.Fix(pq, index)
// }
