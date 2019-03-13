package typesystem

import "fmt"

// MarshalJSON serializes the value to a JSON meta-object
func (o IDType) MarshalJSON() ([]byte, error) {
	if o == IDType(0) {
		return nil, fmt.Errorf("null type of IDType can't be marshaled to JSON")
	}
	return []byte(o.String()), nil
}
