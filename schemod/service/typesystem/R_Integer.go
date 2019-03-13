package typesystem

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// R_Integer implements the type: trait integer
type R_Integer struct {
	v interface{}
}

// Type returns the type identifier
func (o R_Integer) Type() IDType {
	return TR__integer
}

// Len returns the size of the value in bytes
func (o R_Integer) Len() uint64 {
	switch v := o.v.(type) {
	case T_Int32:
		return v.Len()
	case T_Int64:
		return v.Len()
	case T_Uint32:
		return v.Len()
	case T_Uint64:
		return v.Len()
	}
	return 0
}

// ActualType returns the type identifier of the actual type of the value
func (o R_Integer) ActualType() IDType {
	switch o.v.(type) {
	case T_Int32:
		// is of type Int32
		return PT__Int32
	case T_Int64:
		// is of type Int64
		return PT__Int64
	case T_Uint32:
		// is of type Uint32
		return PT__Uint32
	case T_Uint64:
		// is of type Uint64
		return PT__Uint64
	}
	return IDType(0)
}

// SerializeJSON implements the Serializable interface
func (o R_Integer) SerializeJSON() ([]byte, error) {
	typeID := o.ActualType()

	switch typeID {
	case PT__Int32:
		return json.Marshal(struct {
			Type  IDType  `json:"@type"`
			Value T_Int32 `json:"@value"`
		}{
			Type:  AR__ID_A,
			Value: o.v.(T_Int32),
		})
	case PT__Int64:
		return json.Marshal(struct {
			Type  IDType  `json:"@type"`
			Value T_Int64 `json:"@value"`
		}{
			Type:  AS__root_rp_page_struct1,
			Value: o.v.(T_Int64),
		})
	case PT__Uint32:
		return json.Marshal(struct {
			Type  IDType   `json:"@type"`
			Value T_Uint32 `json:"@value"`
		}{
			Type:  AS__root_rp_page_struct1,
			Value: o.v.(T_Uint32),
		})
	case PT__Uint64:
		return json.Marshal(struct {
			Type  IDType   `json:"@type"`
			Value T_Uint64 `json:"@value"`
		}{
			Type:  AS__root_rp_page_struct1,
			Value: o.v.(T_Uint64),
		})
	}

	return nil, fmt.Errorf("unexpected type idenfier: %s", typeID.String())
}

// DeserializeJSON implements the Serializable interface
func (o *R_Integer) DeserializeJSON(b []byte) error {
	// Determine value type
	var metaObject struct {
		Type IDType `json:"@type"`
	}
	if err := json.Unmarshal(b, &metaObject); err != nil {
		return err
	}

	switch metaObject.Type {
	case PT__Int32:
		// Unmarshal value as Int32
		var value struct {
			Value T_Int32 `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	case PT__Int64:
		// Unmarshal value as Int64
		var value struct {
			Value T_Int64 `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	case PT__Uint32:
		// Unmarshal value as Uint32
		var value struct {
			Value T_Uint32 `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	case PT__Uint64:
		// Unmarshal value as Uint64
		var value struct {
			Value T_Uint64 `json:"@value"`
		}
		if err := json.Unmarshal(b, &value); err != nil {
			return err
		}
		o.v = value.Value
	}

	return nil
}

// MarshalJSON serializes the value to a JSON meta-object
func (o R_Integer) MarshalJSON() ([]byte, error) {
	return o.SerializeJSON()
}

// UnmarshalJSON unmarshals the value from a JSON meta-object
func (o *R_Integer) UnmarshalJSON(b []byte) error {
	return o.DeserializeJSON(b)
}

// NewR_Integer constructs a new instance of anonymous union (root.rp.$page)
//
// panics if an unacceptable type is passed to v
func NewR_Integer(v interface{}) R_Integer {
	if AU__root_rp_page.IsValueAssignable(v) {
		panic(fmt.Errorf(
			"value of type %s is not assignable to "+
				"anonymous union (root.rp.$page)",
			reflect.TypeOf(v).String(),
		))
	}
	return R_Integer{
		v: v,
	}
}
