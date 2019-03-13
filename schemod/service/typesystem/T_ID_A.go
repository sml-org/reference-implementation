package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_ID_A implements the type: ID<A>
type T_ID_A ID

// Type returns the type identifier
func (o T_ID_A) Type() IDType {
	return ID__A
}

// Len returns the size of the value in bytes
func (o T_ID_A) Len() uint64 {
	return 16
}

// Serialize serializes the value to the given byte stream
func (o T_ID_A) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	return o.Serialize(byteOrder, stream)
}

// SerializeJSON implements the Serializable interface
func (o T_ID_A) SerializeJSON() ([]byte, error) {
	return json.Marshal(ID(o))
}

// DeserializeJSON implements the Serializable interface
func (o *T_ID_A) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, (*ID)(o))
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_ID_A) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_ID_A) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
