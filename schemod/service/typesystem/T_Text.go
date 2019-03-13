package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_Text implements the type: Text
type T_Text struct {
	Length uint32
	Value  []byte
}

// Type returns the type identifier
func (o T_Text) Type() IDType {
	return PT__Text
}

// Len returns the size of the value in bytes
func (o T_Text) Len() uint64 {
	return uint64(len(o.Value))
}

// Serialize serializes the value to the given byte stream
func (o T_Text) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	_, err := stream.Write(o.Value)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_Text) SerializeJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Text) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, &o.Value)
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_Text) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_Text) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}

// NewT_Text constructs a new instance of type Text
func NewT_Text(s []byte) T_Text {
	v := make([]byte, len(s))
	copy(v, s)
	return T_Text{Value: v}
}
