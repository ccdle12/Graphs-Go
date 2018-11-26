package search

import (
	"github.com/ccdle12/djikstras-algorithm/graph"
	"testing"
)

// TestAllVisitedDFS will test that given a graph,
// all the vertices in the graph would have been
// visited using DFS.
func TestAllVisitedDFS(t *testing.T) {
	// Create the default vertices.
	vertices := graph.NewVertices()

	// Create the a graph using the vertices.
	g := graph.NewGraph(vertices)
	g.PrintGraph()

	// Create the DFS Struct.
	search := New(g)

	// Search using DFS starting at A.
	if ok := search.DFS(vertices[0]); !ok {
		t.Fatalf("DFS was unable to visit all nodes")
	}
}

// TestAllVisitedBFS will test that given a graph,
// all the vertices in the graph would have been
// visited using BFS.
func TestAllVisitedBFS(t *testing.T) {
	// Create the default vertices.
	vertices := graph.NewVertices()

	// Create the a graph using the vertices.
	g := graph.NewGraph(vertices)

	// Create the Search Struct.
	search := New(g)

	// Search using DFS starting at A.
	if ok := search.BFS(vertices[0]); !ok {
		t.Fatalf("BFS was unable to visit all nodes")
	}
}

// TestDijkstra will test that given a graph,
// source vertex and target ID. It will find
// the shortest path.
func TestDijkstra(t *testing.T) {
	// Create the default vertices.
	vertices := graph.NewVertices()

	// Create the graph using the vertices.
	g := graph.NewGraph(vertices)

	// Creat the search struct.
	search := New(g)

	// Search using Dijkstras algorithm for
	// the shortest path.
	expected := []string{"A", "C", "D", "E"}

	// Find the shortest path from the first vertex
	// "A" to the target "E".
	path := search.Dijkstra(g.Vertices[0], "E")

	// Check that the Path returned is the one expected.
	for i, ID := range path {
		if ID != expected[i] {
			t.Fatalf("paths do not match at index: %v"+
				"\n Expecting: %v | Received: %v", i, expected[i], ID)
		}
	}
}
