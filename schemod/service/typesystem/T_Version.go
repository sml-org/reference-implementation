package typesystem

import "encoding/json"

// T_Version implements the type: Version
type T_Version ID

// Type returns the type identifier
func (o T_Version) Type() IDType {
	return PT__Version
}

// Len returns the size of the value in bytes
func (o T_Version) Len() uint64 {
	return 16
}

// SerializeJSON implements the Serializable interface
func (o T_Version) SerializeJSON() ([]byte, error) {
	return json.Marshal(o)
}

// DeserializeJSON implements the Serializable interface
func (o *T_Version) DeserializeJSON(b []byte) error {
	return json.Unmarshal(b, o)
}
