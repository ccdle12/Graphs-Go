package graph

// Vertex are the nodes
// in the graph.
type Vertex struct {
	ID          string
	Connections []*Edge
}

// NewVertices will create all the vertices in
// the graph. This is using uni-directional
// edges represented in an adjacency list.
func NewVertices() []*Vertex {
	// Create all Vertices.
	A := &Vertex{ID: "A"}
	B := &Vertex{ID: "B"}
	C := &Vertex{ID: "C"}
	D := &Vertex{ID: "D"}
	E := &Vertex{ID: "E"}

	// Create all the edges.
	// NOTE: This graph is uni-directional.
	A.Connections = append(A.Connections, &Edge{From: A, To: B, Fee: 4})
	B.Connections = append(B.Connections, &Edge{From: B, To: A, Fee: 4})

	A.Connections = append(A.Connections, &Edge{From: A, To: C, Fee: 2})
	C.Connections = append(C.Connections, &Edge{From: C, To: A, Fee: 2})

	B.Connections = append(B.Connections, &Edge{From: B, To: C, Fee: 1})
	C.Connections = append(C.Connections, &Edge{From: C, To: B, Fee: 1})

	B.Connections = append(B.Connections, &Edge{From: B, To: D, Fee: 2})
	D.Connections = append(D.Connections, &Edge{From: D, To: B, Fee: 2})

	C.Connections = append(C.Connections, &Edge{From: C, To: D, Fee: 4})
	D.Connections = append(D.Connections, &Edge{From: D, To: C, Fee: 4})

	C.Connections = append(C.Connections, &Edge{From: C, To: E, Fee: 5})
	E.Connections = append(E.Connections, &Edge{From: E, To: C, Fee: 5})

	D.Connections = append(D.Connections, &Edge{From: D, To: E, Fee: 1})
	E.Connections = append(E.Connections, &Edge{From: E, To: D, Fee: 1})

	// Create a slice of Vertices and return.
	return []*Vertex{A, B, C, D, E}
}
