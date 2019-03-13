package typesystem

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

// T_Int64 implements the type: Int64
type T_Int64 struct {
	Value int64
}

// ID returns the type identifier
func (o T_Int64) ID() IDType {
	return PT__Int64
}

// Len returns the size of the value in bytes
func (o T_Int64) Len() uint64 {
	return 8
}

// Serialize serializes the value to the given byte stream
func (o T_Int64) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	buf := make([]byte, 8)
	byteOrder.PutUint64(buf, uint64(o.Value))
	_, err := stream.Write(buf)
	return err
}

// SerializeJSON implements the Serializable interface
func (o T_Int64) SerializeJSON() ([]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(o.Value))
	return json.Marshal(buf)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Int64) DeserializeJSON(b []byte) error {
	if len(b) != 8 {
		return fmt.Errorf("Int64 expected size: 8; actual: %d", len(b))
	}
	o.Value = int64(binary.LittleEndian.Uint64(b))
	return nil
}

// MarshalJSON implements the encoding/json.Marshaler interface
func (o T_Int64) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON implements the encoding/json.Unmarshaler interface
func (o *T_Int64) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}
