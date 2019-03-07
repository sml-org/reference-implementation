package typesystem

import (
	"fmt"
	"reflect"
)

// A_1 implements the type: anonymous union (root.rp.$page)
type A_1 struct {
	v interface{}
}

// ID returns the type identifier
func (o A_1) ID() IDType {
	return AU__root_rp_page
}

// NewA_1 constructs a new instance of anonymous union (root.rp.$page)
//
// panics if an unacceptable type is passed to v
func NewA_1(v interface{}) A_1 {
	if !A_1_AcceptsType(v) {
		panic(fmt.Errorf(
			"anonymous union (root.rp.$page) doesn't accept type: %s",
			reflect.TypeOf(v).String(),
		))
	}
	return A_1{
		v: v,
	}
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

// A_1_AcceptsType returns true if v is of an accepted type, otherwise returns
// false
func A_1_AcceptsType(v interface{}) bool {
	if _, ok := v.([]T_ID_A); ok {
		// is of type Array<ID<A>>
		return true
	}
	if _, ok := v.(A_2); ok {
		// is of type anonymous struct root.rp.$page:struct1
		return true
	}
	return false
}
