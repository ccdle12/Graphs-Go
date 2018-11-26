package graph

// Edge represent a channel.
// between the vertexes.
type Edge struct {
	From *Vertex
	To   *Vertex
	Fee  int
}
