package typesystem

// T_Int64 implements the type: Int64
type T_Int64 struct {
	v int64
}

// ID returns the type identifier
func (o T_Int64) ID() IDType {
	return PT__Int64
}
