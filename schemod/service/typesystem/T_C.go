package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"io"
)

// T_C implements the type: struct C
type T_C struct {
	Cb T_B      `json:"cb"`
	Ce T_E      `json:"ce"`
	Cs T_String `json:"cs"`
	Ct T_Text   `json:"ct"`
	Cu T_Uint32 `json:"cu"`
	Ci T_Int32  `json:"ci"`
}

// Type returns the type identifier
func (o *T_C) Type() IDType {
	return ST__C
}

// Len returns the size of the value in bytes
func (o *T_C) Len() uint64 {
	return 0 +
		o.Cb.Len() +
		o.Ce.Len() +
		o.Cs.Len() +
		o.Ct.Len() +
		o.Cu.Len() +
		o.Ci.Len() +
		0
}

// Serialize serializes the value to the given byte stream
func (o T_C) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	if err := o.Cb.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Ce.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Cs.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Ct.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Cu.Serialize(byteOrder, stream); err != nil {
		return err
	}
	if err := o.Ci.Serialize(byteOrder, stream); err != nil {
		return err
	}
	return nil
}

// SerializeJSON implements the Serializable interface
func (o *T_C) SerializeJSON() ([]byte, error) {
	return json.Marshal(o)
}

// DeserializeJSON implements the Serializable interface
func (o *T_C) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

// NewT_C constructs a new instance of type B
func NewT_C(
	cb T_B,
	ce T_E,
	cs T_String,
	ct T_Text,
	cu T_Uint32,
	ci T_Int32,
) *T_C {
	return &T_C{
		Cb: cb,
		Ce: ce,
		Cs: cs,
		Ct: ct,
		Cu: cu,
		Ci: ci,
	}
}
