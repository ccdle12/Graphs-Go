package search

import (
	"container/heap"
	"fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
	"testing"
)

func TestCreatingPQueue(t *testing.T) {
	// Get all vertices.
	vertices := graph.NewVertices()

	// Add the vertices to the items map.
	items := make(map[interface{}]int)
	for i, vertex := range vertices {
		items[vertex] = i
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    &graph.Vertex{ID: "Z"},
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		// Pop an item from the queue.
		item := heap.Pop(&pq).(*Item)

		// Cast the item to a vertex.
		v := item.value.(*graph.Vertex)

		fmt.Printf("%.2d:%s\n", item.priority, v.ID)
	}
}
