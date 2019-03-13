package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

// A_2 implements the anonymous type: struct (root.rp.$page.1)
type A_2 struct {
	Cursor T_ID_A    `json:"cursor"`
	Limit  R_Integer `json:"limit"`
}

// Type returns the type identifier
func (o *A_2) Type() IDType {
	return AS__root_rp_page_struct1
}

// Len returns the size of the value in bytes
func (o *A_2) Len() uint64 {
	return 0 +
		o.Cursor.Len() +
		o.Limit.Len() +
		0
}

// Serialize serializes the value to the given byte stream
func (o *A_2) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	//TODO: implement
	/*
		if err := o.cursor.Serialize(byteOrder, stream); err != nil {
			return err
		}
			if err := o.limit.Serialize(byteOrder, stream); err != nil {
				return err
			}
	*/
	return fmt.Errorf("not yet implemented")
}

// SerializeJSON implements the Serializable interface
func (o *A_2) SerializeJSON() ([]byte, error) {
	return json.Marshal(o)
}

// DeserializeJSON implements the Serializable interface
func (o *A_2) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

// NewA_2 creates a new instance of the anonymous type: struct (root.rp.$page.1)
func NewA_2(
	cursor T_ID_A,
	limit R_Integer,
) *A_2 {
	return &A_2{
		Cursor: cursor,
		Limit:  limit,
	}
}
