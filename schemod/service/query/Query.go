package query

import (
	"io"
	"reference-implementation/schemod/service/typesystem"
)

// Argument represents a query argument
type Argument struct {
	Type  typesystem.IDType
	Value interface{}
}

// Args represents a set of assigned arguments
type Args = map[typesystem.IDParameter]Argument

// Props represents a set selected properties
type Props = map[typesystem.IDProperty]PropertySelection

// Attrs represents a set selected attributes
type Attrs = map[typesystem.IDAttribute]AttributeSelection

// AttributeSelection represents a property selection
type AttributeSelection struct {
	Props map[typesystem.IDProperty]PropertySelection
}

// PropertySelection represents a property selection
type PropertySelection struct {
	Args  map[typesystem.IDParameter]Argument
	Props map[typesystem.IDProperty]PropertySelection
	Attrs map[typesystem.IDAttribute]AttributeSelection
}

// Query represents an API query
type Query struct {
	Props map[typesystem.IDProperty]PropertySelection
}

// SerializeTo serializes the query to a byte stream
func (q Query) SerializeTo(stream io.WriteCloser) error {
	return nil
}
