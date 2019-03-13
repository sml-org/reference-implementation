package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_E represents the value type of enum E
type T_E uint32

const (
	_ T_E = iota

	// T_E_e1 reflects E::e1
	T_E_e1

	// T_E_e2 reflects E::e2
	T_E_e2

	// T_E_e3 reflects E::e3
	T_E_e3

	// T_E_e4 reflects E::e4
	T_E_e4
)

// String returns the string
func (o T_E) String() string {
	switch o {
	case T_E_e1:
		return "E::e1"
	case T_E_e2:
		return "E::e2"
	case T_E_e3:
		return "E::e3"
	case T_E_e4:
		return "E::e4"
	}
	return ""
}

// Type returns the type identifier
func (o T_E) Type() IDType {
	return NU__E
}

// Len returns the size of the value in bytes
func (o T_E) Len() uint64 {
	return 4
}

// Serialize serializes the value to the given byte stream
func (o T_E) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	buf := make([]byte, 4)
	byteOrder.PutUint32(buf, uint32(o))
	_, err := stream.Write(buf)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_E) SerializeJSON() ([]byte, error) {
	return json.Marshal(uint32(o))
}

// DeserializeJSON implements the Serializable interface
func (o *T_E) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

// MarshalJSON serializes the value to a JSON token
func (o T_E) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_E) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
