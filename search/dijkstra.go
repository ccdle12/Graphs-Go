package search

import (
	"container/heap"
	"fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
)

// Dijkstra is the open function call that
// will cal the private implementation of
// Dijkstras algorithm to find the shortest
// path in a graph.
func (s *Search) Dijkstra(v *graph.Vertex) {
	// Create a cache.
	c := s.newCache()

	// Use private implementation to find
	// shortest path.
	s.dijkstra(v, c)
}

func (s *Search) dijkstra(v *graph.Vertex, c *cache) {
	fmt.Printf("[Starting Dijkstra]\n")

	// Loop through each vertex in the graph.
	// Add each vertex to the map and set as
	// infinite distance.
	for _, vertex := range s.Graph.Vertices {
		c.Distance[vertex] = MaxInt
	}

	// Set the source vertex's distance as 0.
	c.Distance[v] = 0

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
		poppedVertex := poppedItem.value.(*graph.Vertex)
		fmt.Printf("Popped vertex ID: %v\n", poppedVertex.ID)

		// If the vertex has already been visited
		// continue.
		if c.Visited[poppedVertex] {
			fmt.Printf("Vertex: %v has already been visited\n", poppedVertex.ID)
			continue
		}

		// Set the vertex to visited.
		c.Visited[poppedVertex] = true

		// Loop over each edge in the popped vertex.
		for _, edge := range poppedVertex.Connections {
			// If the distance of the popped vertex + the fee to the edge.To (vertex)
			// is greater than the cached distance of the edge.To (vertex), then update
			// the edge.To (vertex), this represents the shortest path so far to the
			// edge.To (vertex).
			if c.Distance[poppedVertex]+edge.Fee < c.Distance[edge.To] {
				c.Distance[edge.To] = c.Distance[poppedVertex] + edge.Fee

				// Push the visited node to the queue.
				item := &Item{
					value:    edge.To,
					priority: c.Distance[edge.To],
					index:    len(c.Queue),
				}
				heap.Push(&c.Queue, item)
			}
		}
	}

	// Loop over the graph vertices and print it's distance
	// from the map.
	for _, vertex := range s.Graph.Vertices {
		// Print the distance.
		fmt.Printf("Vertex: %v | Distance: %v\n", vertex.ID, c.Distance[vertex])
	}
}
