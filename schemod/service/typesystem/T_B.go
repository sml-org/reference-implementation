package typesystem

// T_B implements the type: struct B
type T_B struct {
	bo T_String
	bx T_String
	by bool
	bz bool
}

// ID returns the type identifier
func (o *T_B) ID() IDType {
	return ST__B
}

// NewT_B constructs a new instance of struct B
func NewT_B(
	bo T_String,
	bx T_String,
	by bool,
	bz bool,
) *T_B {
	return &T_B{
		bo: bo,
		bx: bx,
		by: by,
		bz: bz,
	}
}
