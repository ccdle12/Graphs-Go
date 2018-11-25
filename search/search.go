package search

import (
	"github.com/ccdle12/djikstras-algorithm/graph"
)

// Search struct is the struct for performing different,
// types of searches in a graph.
// It contains a graph.
type Search struct {
	Graph *graph.Graph
}

// cache is a closed struct used when calling different
// types of searches. It will cache the results in this
// struct.
type cache struct {
	// A map of visited vertices.
	Visited map[*graph.Vertex]bool

	// A map of each vertex distance.
	Distance map[*graph.Vertex]int

	// Queue is for vertices to be processed.
	Queue PriorityQueue
}

// newCache initialises all the fields for the cache
// and returns it.
func (s *Search) newCache() *cache {
	return &cache{
		Visited:  make(map[*graph.Vertex]bool),
		Distance: make(map[*graph.Vertex]int),
		Queue:    make(PriorityQueue, 0),
	}
}

// New creates a Search struct initialised with a
// graph.
func New(g *graph.Graph) *Search {
	return &Search{Graph: g}
}
