package pregel

type Verttex struct {
	ID    string
	Value []byte
}

type Edge struct {
	From  string
	To    string
	Value []byte
}
type VertexOperationType int
type EdgeOperationType int

const (
	_ VertexOperationType = iota
	VertexAdded
	VertexRemoved
	VertexValueChanged
)

const (
	_ EdgeOperationType = iota
	EdgeAdded
	EdgeRemoved
	EdgeValueChanged
)

type VertexMessage struct {
	To        string
	JobId     string
	SuperStep int
	Value     []byte
}

type VertexHalt struct {
	To        string
	JobId     string
	SuperStep int
}

type VertexOperation struct {
	ID        string
	JobId     string
	Superstep int
	Type      VertexOperationType
	Value     []byte
}

type EdgeOperation struct {
	From      string
	JobId     string
	SuperStep int
	To        string
	Value     []byte
	Type EdgeOperationType
}
