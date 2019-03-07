package typesystem

// T_A implements the type: entity A
type T_A struct {
	ab *T_B
	ac T_C
}

// ID returns the type identifier
func (o T_A) ID() IDType {
	return EN__A
}

// NewT_A constructs a new instance of entity A
func NewT_A(
	ab *T_B,
	ac T_C,
) *T_A {
	return &T_A{
		ab: ab,
		ac: ac,
	}
}
