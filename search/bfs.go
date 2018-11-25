package search

import (
	"container/heap"
	// "fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
)

// BFS is the open function to call. This will call bfs to
// perform the breadth first search.
func (s *Search) BFS(v *graph.Vertex) bool {
	// Create a cache for the DFS.
	cache := s.newCache()

	// Call the closed recursive implementation.
	c := s.bfs(v, cache)

	// Return a bool after checking the cache against
	// the graph.
	return s.checkAllVisited(c)
}

// Closed implementation of bfs.
func (s *Search) bfs(v *graph.Vertex, c *cache) *cache {
	// Set the source vertex "v" as visited.
	c.Visited[v] = true

	// Push the vertex on to the queue.
	item := &Item{
		value:    v,
		priority: 0,
		index:    0,
	}
	heap.Push(&c.Queue, item)

	// Loop while the queue is not empty.
	for c.Queue.Len() > 0 {
		// Pop an item off the queue.
		poppedItem := heap.Pop(&c.Queue).(*Item)

		// Cast the value to a Vertex.
		vertex := poppedItem.value.(*graph.Vertex)

		// Loop through all the edges from this vertex.
		for _, edge := range vertex.Connections {
			// If the "To" vertex has been visited,
			// skip it.
			if ok := c.Visited[edge.To]; ok {
				continue
			}

			// Set the "To" vertex as visited.
			c.Visited[edge.To] = true

			// Push the visited node to the queue.
            // In this case, priority will act as 
            // distance.
			item := &Item{
				value:    edge.To,
				priority: poppedItem.priority + 1,
				index:    len(c.Queue),
			}
			heap.Push(&c.Queue, item)
		}
	}

	return c
}
