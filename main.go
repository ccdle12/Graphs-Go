package main

import (
	"fmt"
	"github.com/ccdle12/djikstras-algorithm/graph"
	"github.com/ccdle12/djikstras-algorithm/search"
)

func main() {
	// Create all vertices.
	vertices := graph.NewVertices()

	// Create a graph using an adjaceny list
	// of vertices.
	graph := graph.NewGraph(vertices)

	// Print out the topology of the graph.
	graph.PrintGraph()

	// Create a DFS.
	fmt.Println("[PERFORMING DEPTH FIRST SEARCH]")
	search := search.New(graph)
	search.DFS(graph.Vertices[0])

	// TESTING
	search.BellManford(graph.Vertices[0])
	search.Dijkstra(graph.Vertices[0], "E")
}
