package typesystem

// T_Bool implements the type: Bool
type T_Bool struct {
	v bool
}

// ID returns the type identifier
func (o T_Bool) ID() IDType {
	return PT__Bool
}

// NewT_Bool constructs a new instance of Bool
func NewT_Bool(v bool) T_Bool {
	return T_Bool{
		v: v,
	}
}
