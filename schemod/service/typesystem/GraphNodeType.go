package typesystem

// GraphNodeType represents the identifier of a category of graph nodes
type GraphNodeType byte

const (
	_ GraphNodeType = iota

	// Property represents the identifier of the graph property nodes category
	Property

	// Parameter represents the identifier of the graph parameter nodes category
	Parameter

	// Attribute represents the identifier of the graph attribute nodes category
	Attribute
)
