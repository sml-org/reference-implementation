package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_B implements the type: struct B
type T_B struct {
	Bo T_String `json:"bo"`
	Bx T_String `json:"bx"`
	By T_Bool   `json:"by"`
	Bz T_Bool   `json:"bz"`
}

// Type returns the type identifier
func (o *T_B) Type() IDType {
	return ST__B
}

// Len returns the size of the value in bytes
func (o T_B) Len() uint64 {
	return 0 +
		o.Bo.Len() +
		o.Bx.Len() +
		o.By.Len() +
		o.Bz.Len() +
		0
}

// Serialize serializes the value to the given byte stream
func (o T_B) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	if err := o.Bo.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Bx.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.By.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Bz.Serialize(byteOrder, stream); err != nil {
		return err
	}
	return nil
}

// SerializeJSON implements the Serializable interface
func (o *T_B) SerializeJSON() ([]byte, error) {
	return json.Marshal(o)
}

// DeserializeJSON implements the Serializable interface
func (o *T_B) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

// NewT_B constructs a new instance of type B
func NewT_B(
	bo T_String,
	bx T_String,
	by T_Bool,
	bz T_Bool,
) *T_B {
	return &T_B{
		Bo: bo,
		Bx: bx,
		By: by,
		Bz: bz,
	}
}
