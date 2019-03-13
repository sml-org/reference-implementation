package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

// T_Uint64 implements the type: Uint64
type T_Uint64 struct {
	Value uint64
}

// Type returns the type identifier
func (o T_Uint64) Type() IDType {
	return PT__Uint64
}

// Len returns the size of the value in bytes
func (o T_Uint64) Len() uint64 {
	return 8
}

// Serialize serializes the value to the given byte stream
func (o T_Uint64) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	buf := make([]byte, 4)
	byteOrder.PutUint64(buf, uint64(o.Value))
	_, err := stream.Write(buf)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_Uint64) SerializeJSON() ([]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(o.Value))
	return json.Marshal(buf)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Uint64) DeserializeJSON(b []byte) error {
	if len(b) != 8 {
		return fmt.Errorf("Uint64 expected size: 8; actual: %d", len(b))
	}
	o.Value = binary.LittleEndian.Uint64(b)
	return nil
}

// MarshalJSON serializes the value to a JSON token
func (o T_Uint64) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the Go JSON interface
func (o *T_Uint64) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
