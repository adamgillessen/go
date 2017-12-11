package apq

import (
	"container/heap"
	"fmt"
)

// item is used to store a value with some key (priority)
type item struct {
	value      interface{}
	index, key int
}

// itemHeap is used internally to store the APQ items
type itemHeap []*item

// Push pushes item x onto the heap
func (h *itemHeap) Push(x interface{}) {
	x.(*item).index = len(*h)
	*h = append(*h, x.(*item))
}

// Pop removes and returns the element at position (*h).Len()
func (h *itemHeap) Pop() interface{} {
	l := len(*h)
	i := (*h)[l-1]
	*h = (*h)[:l-1]
	return i
}

func (h itemHeap) Len() int           { return len(h) }
func (h itemHeap) Less(i, j int) bool { return h[i].key < h[j].key }
func (h itemHeap) Swap(i, j int) {
	h[i].index = j
	h[j].index = i
	h[i], h[j] = h[j], h[i]
}

// APQ is an Adaptive Priority Queue which allows you to queue elements
// priority which can be updated with a reference to the value.
// Note: the APQ assumes that all values are distinct.
type APQ struct {
	h         *itemHeap
	locations map[interface{}]*item
}

// New returns a new empty queue
func New() *APQ {
	return &APQ{
		h:         &itemHeap{},
		locations: make(map[interface{}]*item),
	}
}

// Enqueue puts a value into the queue with the given priority
func (q APQ) Enqueue(x interface{}, priority int) {
	i := &item{
		value: x,
		key:   priority,
	}
	heap.Push(q.h, i)
	q.locations[x] = i
}

// Dequeue removes and returns the value which has *lowest* priority
func (q APQ) Dequeue() interface{} {
	i := heap.Pop(q.h).(*item)
	return i.value
}

// GetPriority returns the current priority of value x.
// An error is returned if the value is not in the queue.
func (q APQ) GetPriority(x interface{}) (int, error) {
	i, ok := q.locations[x]
	if !ok {
		return 0, fmt.Errorf("GetPriority(): value %v was not in the queue", x)
	}
	return i.key, nil
}

// UpdatePriority updates the priority of the value x with a new priority.
// An error is returned if the value is not in the queue.
func (q APQ) UpdatePriority(x interface{}, priority int) error {
	i, ok := q.locations[x]
	if !ok {
		return fmt.Errorf("UpdatePriority(): value %v was not in the queue", x)
	}
	i.key = priority
	heap.Fix(q.h, i.index)
	return nil
}

// Len returns the number of items in the queue
func (q APQ) Len() int {
	return len(*q.h)
}
