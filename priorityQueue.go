package common

import "container/heap"

type PriorityQueueItem interface {
	GetPriority() int
	GetIndex() int
	SetIndex(int)
}

type PriorityQueue[T PriorityQueueItem] struct {
	heap priorityQueueHeap[T]
}

type priorityQueueHeap[T PriorityQueueItem] []T

func (h priorityQueueHeap[T]) Len() int {
	return len(h)
}

func (h priorityQueueHeap[T]) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.GetPriority() < b.GetPriority()
}

func (h priorityQueueHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].SetIndex(i)
	h[j].SetIndex(j)
}

func (h *priorityQueueHeap[T]) Push(x any) {
	n := len(*h)
	item := x.(T)
	item.SetIndex(n)
	*h = append(*h, item)
}

func (h *priorityQueueHeap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	item.SetIndex(-1) // for safety
	*h = old[0 : n-1]
	return item
}

func NewPriorityQueue[T PriorityQueueItem]() *PriorityQueue[T] {
	p := &PriorityQueue[T]{}
	heap.Init(&p.heap)
	return p
}

func (p *PriorityQueue[T]) Push(item T) {
	heap.Push(&p.heap, item)
}

func (p *PriorityQueue[T]) Pop() (item T) {
	return heap.Pop(&p.heap).(T)
}

func (p *PriorityQueue[T]) Fix(item T) {
	heap.Fix(&p.heap, item.GetIndex())
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return len(p.heap) == 0
}
