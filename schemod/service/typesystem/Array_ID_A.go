package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

// Array_ID_A implements the type: Array<ID<A>>
type Array_ID_A struct {
	Items []T_ID_A
}

// Type returns the type identifier
func (o Array_ID_A) Type() IDType {
	return AR__ID_A
}

// Len returns the size of the value in bytes
func (o Array_ID_A) Len() uint64 {
	return uint64(len(o.Items)) * 16
}

// Serialize serializes the value to the given byte stream
func (o Array_ID_A) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	return fmt.Errorf("not yet implemented")
}

// SerializeJSON implements the Serializable interface
func (o Array_ID_A) SerializeJSON() ([]byte, error) {
	return json.Marshal(o.Items)
}

// DeserializeJSON implements the Serializable interface
func (o *Array_ID_A) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, &o.Items)
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o Array_ID_A) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *Array_ID_A) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
