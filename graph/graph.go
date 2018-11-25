package graph

import (
	"fmt"
)

// Graph is a struct representing
// the whole graph structure, holding
// Vertices.
type Graph struct {
	Vertices []*Vertex
}

// NewGraph returns an instance of a Graph,
// given a list of vertices.
func NewGraph(vertices []*Vertex) *Graph {
	return &Graph{Vertices: vertices}
}

// PrintGraph is a helper function,
// to print out the topology of the graph.
// Time Complexity:
// O(nm)
// n := number of vertices
// m := number of connections/edges
func (g *Graph) PrintGraph() {
	// Loop over each vertex in the graph.
	for _, vertex := range g.Vertices {
		fmt.Printf("Vertex ID: %v\n", vertex.ID)

		// Loop over each connection/edge in the graph.
		for _, edge := range vertex.Connections {
			fmt.Printf("Edge From: %v\n", edge.From.ID)
			fmt.Printf("Edge To: %v\n", edge.To.ID)
			fmt.Printf("Edge Fee: %v\n", edge.Fee)
			fmt.Printf("-------------------------\n")
		}
		fmt.Printf("-------------------------\n")
	}
}
