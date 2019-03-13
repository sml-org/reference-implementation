package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_String implements the type: String
type T_String struct {
	Value []byte
}

// Type returns the type identifier
func (o T_String) Type() IDType {
	return PT__String
}

// Len returns the size of the value in bytes
func (o T_String) Len() uint64 {
	return uint64(len(o.Value))
}

// Serialize serializes the value to the given byte stream
func (o T_String) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	_, err := stream.Write(o.Value)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_String) SerializeJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

// DeserializeJSON implements the Serializable interface
func (o *T_String) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, &o.Value)
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_String) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_String) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}

// NewT_String constructs a new instance of type String
func NewT_String(v []byte) T_String {
	return T_String{
		Value: v,
	}
}
