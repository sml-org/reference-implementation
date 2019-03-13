package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_Uint32 implements the type: Uint32
type T_Uint32 struct {
	Value uint32
}

// Type returns the type identifier
func (o T_Uint32) Type() IDType {
	return PT__Uint32
}

// Len returns the size of the value in bytes
func (o T_Uint32) Len() uint64 {
	return 4
}

// Serialize serializes the value to the given byte stream
func (o T_Uint32) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	buf := make([]byte, 4)
	byteOrder.PutUint32(buf, uint32(o.Value))
	_, err := stream.Write(buf)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_Uint32) SerializeJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Uint32) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, &o.Value)
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_Uint32) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_Uint32) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
