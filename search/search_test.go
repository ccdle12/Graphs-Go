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

	// Create the DFS Struct.
	search := New(g)

	// Search using DFS starting at A.
	if ok := search.BFS(vertices[0]); !ok {
		t.Fatalf("BFS was unable to visit all nodes")
	}
}
