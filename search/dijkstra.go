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
//
// TODO: (ccdle12
// Should either use source and target as vertices
// or use both as string IDs.
func (s *Search) Dijkstra(sourceVertex *graph.Vertex, targetID string) []string {
	// Create a cache.
	c := s.newCache()

	// Use private implementation to find
	// shortest path.
	return s.dijkstra(sourceVertex, targetID, c)
}

func (s *Search) dijkstra(v *graph.Vertex, targetID string, c *cache) []string {
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

				// Add a Previous pointer to the vertex, when querying the path from a
				// a vertex to the source, it will follow the pointers.
				s.Graph.Previous[edge.To.ID] = poppedVertex

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

	path := s.findShortestPath("A", "E")
	fmt.Printf("%v\n", path)

	return path
}

// findShortestPath is a helper function for djikstra. Since the Graph will have
// an update map of Previous pointers, findShortestPath will follow the Previous
// pointers from the target back to the source and return an array of IDs from
// the perspective of Source -> Target.
func (s *Search) findShortestPath(source, target string) []string {
	// NOTE: (ccdle12)
	// This will return the path given an ID
	path := make([]string, 0, 0)

	// Set current to the target, we will traverse from the
	// the target back to the source.
	current := target

	for current != source {
		// Add current to the Path.
		path = append([]string{current}, path...)
		current = s.Graph.Previous[current].ID
	}

	// Add the source as the first vertex in the path.
	path = append([]string{source}, path...)

	return path

}
