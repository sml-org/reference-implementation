package typesystem

// T_Int32 implements the type: Int32
type T_Int32 struct {
	v int32
}

// ID returns the type identifier
func (o T_Int32) ID() IDType {
	return PT__Int32
}
