package search

import "fmt"

// checkAllVisited is a util function, given a cache,
// it will check against the known existing graph
// in Search to determine whether all vertices
// were visited.
func (s *Search) checkAllVisited(c *cache) bool {
	fmt.Printf("Length of vertex in graph: %v\n", len(s.Graph.Vertices))
	// Confirm that all the vertices in the graph
	// have been visited.
	for _, vertex := range s.Graph.Vertices {
		fmt.Printf("Vertex in graph: %v\n", vertex.ID)

		// Make sure the vertex exists in the
		// graph. This is more of a sanity check.
		// TODO: (ccdle12): should I return an error?
		if ok := c.Visited[vertex]; !ok {
			fmt.Printf("Vertex is not in cache")
			return false
		}

		// The vertex was not visited.
		if !c.Visited[vertex] {
			return false
		}
	}

	return true
}
