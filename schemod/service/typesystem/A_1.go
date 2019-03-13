package typesystem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// A_1 implements the type: anonymous union (root.rp.$page)
type A_1 struct {
	v interface{}
}

// Type returns the type identifier
func (o A_1) Type() IDType {
	return AU__root_rp_page
}

// ActualType returns the type identifier of the actual type of the union
func (o A_1) ActualType() IDType {
	switch o.v.(type) {
	case []T_ID_A:
		// is of type Array<ID<A>>
		return AR__ID_A
	case A_2:
		// is of type anonymous struct root.rp.$page:struct1
		return AS__root_rp_page_struct1
	}
	return IDType(0)
}

// SerializeJSON implements the Serializable interface
func (o A_1) SerializeJSON() ([]byte, error) {
	typeID := o.ActualType()

	switch typeID {
	case AR__ID_A:
		return json.Marshal(struct {
			Type  IDType   `json:"@type"`
			Value []T_ID_A `json:"@value"`
		}{
			Type:  AR__ID_A,
			Value: o.v.([]T_ID_A),
		})
	case AS__root_rp_page_struct1:
		return json.Marshal(struct {
			Type  IDType `json:"@type"`
			Value A_2    `json:"@value"`
		}{
			Type:  AS__root_rp_page_struct1,
			Value: o.v.(A_2),
		})
	}

	return nil, fmt.Errorf("unexpected type idenfier: %s", typeID.String())
}

// DeserializeJSON implements the Serializable interface
func (o *A_1) DeserializeJSON(b []byte) error {
	// Determine value type
	var metaObject struct {
		Type IDType `json:"@type"`
	}
	if err := json.Unmarshal(b, &metaObject); err != nil {
		return err
	}

	switch metaObject.Type {
	case AR__ID_A:
		// Unmarshal value as Array<ID<A>>
		var value struct {
			Value []T_ID_A `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	case AS__root_rp_page_struct1:
		// Unmarshal value as struct (root.rp.$page.1)
		var value struct {
			Value A_2 `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	}

	return nil
}

// MarshalJSON serializes the value to a JSON meta-object
func (o A_1) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON unmarshals the value from a JSON meta-object
func (o *A_1) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}

// NewA_1 constructs a new instance of anonymous union (root.rp.$page)
//
// panics if an unacceptable type is passed to v
func NewA_1(v interface{}) A_1 {
	if AU__root_rp_page.IsValueAssignable(v) {
		panic(fmt.Errorf(
			"value of type %s is not assignable to "+
				"anonymous union (root.rp.$page)",
			reflect.TypeOf(v).String(),
		))
	}
	return A_1{
		v: v,
	}
}
