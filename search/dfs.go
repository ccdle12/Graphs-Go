package search

import (
	// "fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
)

// DFS is the open function to call. This will call dfs to
// perform a recursive search using Depth Fist.
//
// Time Complexity:
// n := number of vertices
// m := number of edges
//
// O(n + m)
func (s *Search) DFS(v *graph.Vertex) bool {
	// Create a cache for the DFS.
	cache := s.newCache()

	// Call the closed recursive implementation.
	c := s.dfs(v, cache)

	// Return a bool after checking the cache against
	// the graph.
	return s.checkAllVisited(c)
}

// Closed implementation of dfs that will use recursion.
func (s *Search) dfs(v *graph.Vertex, c *cache) *cache {
	// fmt.Printf("At vertex: %v\n", v.ID)

	// If the vertex has already been visited,
	// return the up the recursion stack.
	if c.Visited[v] {
		return c
	}

	// Set visited to true.
	c.Visited[v] = true

	// Visit each connection from this vertex.
	for _, edge := range v.Connections {
		s.dfs(edge.To, c)
	}

	return c
}
