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
