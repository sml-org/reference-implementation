package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_Bool implements the type: Bool
type T_Bool struct {
	Value bool
}

// Type returns the type identifier
func (o T_Bool) Type() IDType {
	return PT__Bool
}

// Len returns the size of the value in bytes
func (o T_Bool) Len() uint64 {
	return 1
}

// Serialize serializes the value to the given byte stream
func (o T_Bool) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	buf := make([]byte, 1)
	if o.Value {
		// true
		buf[0] = byte(1)
	} else {
		// false
		buf[0] = byte(0)
	}
	_, err := stream.Write(buf)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_Bool) SerializeJSON() ([]byte, error) {
	return json.Marshal(o.Value)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Bool) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, &o.Value)
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_Bool) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_Bool) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
