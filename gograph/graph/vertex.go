package graph

// Vertex represents a vertex in a graph.
type Vertex interface {
	// Update updates the vertex state for a given step and sends messages to other vertices.
	// It returns a boolean indicating whether the vertex is still active after the update.
	Update(step int, engineChan chan VertexMessage) bool

	// GetID returns the ID of the vertex.
	GetID() int

	// GetValue returns the value of the vertex.
	GetValue() float64

	// GetOutVertices returns the IDs of the vertices that the current vertex has outgoing edges to.
	GetOutVertices() []int

	// GetActive returns a boolean indicating whether the vertex is active.
	GetActive() bool

	// GetSuperStep returns the current superstep of the vertex.
	GetSuperStep() int

	// GetMessages returns the messages received by the vertex.
	GetMessages() []VertexMessage

	// ReceiveMessages receives messages sent to the vertex.
	ReceiveMessages(msg []VertexMessage)
}

// BaseVertex is a basic implementation of the Vertex interface.
type BaseVertex struct {
	ID          int
	Value       float64
	OutVertices []int
	IncMsgs     []VertexMessage
	Active      bool
	SuperStep   int
}

// GetID returns the ID of the vertex.
func (v *BaseVertex) GetID() int {
	return v.ID
}

// GetValue returns the value of the vertex.
func (v *BaseVertex) GetValue() float64 {
	return v.Value
}

// GetOutVertices returns the IDs of the vertices that the current vertex has outgoing edges to.
func (v *BaseVertex) GetOutVertices() []int {
	return v.OutVertices
}

// GetSuperStep returns the current superstep of the vertex.
func (v *BaseVertex) GetSuperStep() int {
	return v.SuperStep
}

// GetActive returns a boolean indicating whether the vertex is active.
func (v *BaseVertex) GetActive() bool {
	return v.Active
}

// GetMessages returns the messages received by the vertex.
func (v *BaseVertex) GetMessages() []VertexMessage {
	return v.IncMsgs
}
