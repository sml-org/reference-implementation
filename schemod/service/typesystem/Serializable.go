package typesystem

import (
	"encoding/binary"
	"io"
)

// Serializable defines the interface of serializable types
type Serializable interface {
	// Type returns the type identifier
	Type() IDType

	// Len returns the size of the value in bytes
	Len() uint64

	// Serialize serializes the value to the given byte stream
	Serialize(
		byteOrder binary.ByteOrder,
		stream io.Writer,
	) error

	// SerializeJSON serializes the value to JSON
	SerializeJSON() ([]byte, error)

	// DeserializeJSON deserializes the value from JSON
	DeserializeJSON([]byte) error
}
