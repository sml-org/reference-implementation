package query

import (
	"encoding/binary"
	"io"
	"reference-implementation/schemod/service/typesystem"
)

// Args represents a set of assigned arguments
type Args = map[typesystem.IDParameter]typesystem.Serializable

// Props represents a set selected properties
type Props = map[typesystem.IDProperty]PropertySelection

// Attrs represents a set selected attributes
type Attrs = map[typesystem.IDAttribute]AttributeSelection

// AttributeSelection represents a property selection
type AttributeSelection struct {
	Props map[typesystem.IDProperty]PropertySelection `json:"prop"`
}

// PropertySelection represents a property selection
type PropertySelection struct {
	Args  map[typesystem.IDParameter]typesystem.Serializable `json:"arg"`
	Props map[typesystem.IDProperty]PropertySelection        `json:"prop"`
	Attrs map[typesystem.IDAttribute]AttributeSelection      `json:"attr"`
}

// Query represents an API query
type Query struct {
	Props map[typesystem.IDProperty]PropertySelection `json:"p"`
}

func (q Query) SerializeJSON() ([]byte, error) {
	return nil, nil
}

// Serialize serializes the query to a byte stream
func (q Query) Serialize(
	byteOrder binary.ByteOrder,
	stream io.WriteCloser,
) error {
	panic("not yet implemented")
	/*
		// Write header
		if _, err := stream.Write(queryHeader); err != nil {
			return err
		}

		nodeIndexCounter := uint32(0)

		// Write root properties
		for propID, propSelection := range q.Props {
			if err := serializeProperty(
				&nodeIndexCounter,
				0,
				byteOrder,
				stream,
				propID,
				propSelection,
			); err != nil {
				return err
			}
		}

		return stream.Close()
	*/
}
