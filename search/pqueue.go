package search

import (
	"container/heap"
)

// Item is a struct managed in the priority queue.
type Item struct {
	value    interface{}
	priority int
	index    int
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Returns the length of the priority queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less is an implementation of the heap.Interface.
//
// This determines whether the heap will order the
// itmes from low priority first or high priority.
//
// In the case below it will pop the lowest
// priority and subsequently the items  will
// range from smallest to largest.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push a new Item into the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	// Get the length of the queue.
	n := len(*pq)

	// Cast the input to an Item.
	item := x.(*Item)

	// Assign an index to the item (last index of the queue).
	item.index = n

	// Append the item to the queue.
	*pq = append(*pq, item)
}

// Pop removes an Item from the queue.
func (pq *PriorityQueue) Pop() interface{} {
	// Create a reference to the priority queue.
	old := *pq

	// Get the length of the queue.
	n := len(old)

	// Reference the last item in the queue.
	item := old[n-1]

	// Set the index to -1 for safety, meaning
	// it's being removed.
	item.index = -1 // for safety

	// Slice the old queue by removing the last item.
	*pq = old[0 : n-1]

	// Return the retrieved item.
	return item
}

// Update an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
