package typesystem

// T_ID_A implements the type: ID<A>
type T_ID_A struct {
	v ID
}

// ID returns the type identifier
func (o T_ID_A) ID() IDType {
	return ID__A
}

// NewT_ID_A creates a new instance of ID<A>
func NewT_ID_A(v ID) T_ID_A {
	return T_ID_A{
		v: v,
	}
}
