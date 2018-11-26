package search

import (
	"fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
)

// BellManford is a search algorithm to find
// the shortest path in a graph.
//
// This implementation assumes the edges to
// be kept as an adejaceny list.
//
// Time Complexity: O(nm)
// n := each vertex
// m := each edge from the vertex
func (s *Search) BellManford(v *graph.Vertex) {
	fmt.Printf("[Starting BellManford]\n")

	// Create the cache.
	cache := s.newCache()

	// Loop through each vertex in the graph.
	// Add each vertex to the map and set as
	// infinite distance.
	for _, vertex := range s.Graph.Vertices {
		cache.Distance[vertex] = MaxInt
	}

	// Set the source vertex's distance as 0.
	cache.Distance[v] = 0

	// Loop over the length of the graph.
	for _, vertex := range s.Graph.Vertices {
		// Loop through the edges of each
		// vertex.
		for _, edge := range vertex.Connections {
			// Calculate the min distance.
			var min int
			if cache.Distance[edge.To] < cache.Distance[edge.From]+edge.Fee {
				min = cache.Distance[edge.To]
			} else {
				min = cache.Distance[edge.From] + edge.Fee
			}

			fmt.Printf("Vertex From: %v\n", edge.From.ID)
			fmt.Printf("Vertex To: %v\n", edge.To.ID)
			fmt.Printf("Min Distance: %v\n", min)
			fmt.Printf("--------------------------------\n")

			cache.Distance[edge.To] = min
		}
	}
}
