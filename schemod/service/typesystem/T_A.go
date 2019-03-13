package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_A implements the type: entity A
type T_A struct {
	Ab *T_B `json:"ab"`
	Ac T_C  `json:"ac"`
}

// Type returns the type identifier
func (o T_A) Type() IDType {
	return EN__A
}

// Serialize serializes the value to the given byte stream
func (o T_A) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	panic("not yet implemented")
}

// SerializeJSON implements the Serializable interface
func (o *T_A) SerializeJSON() ([]byte, error) {
	return json.Marshal(o)
}

// DeserializeJSON implements the Serializable interface
func (o *T_A) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

// NewT_A constructs a new instance of entity A
func NewT_A(
	ab *T_B,
	ac T_C,
) *T_A {
	return &T_A{
		Ab: ab,
		Ac: ac,
	}
}
