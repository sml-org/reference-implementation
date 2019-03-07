package typesystem

// T_Uint64 implements the type: Uint64
type T_Uint64 struct {
	v uint64
}

// ID returns the type identifier
func (o T_Uint64) ID() IDType {
	return PT__Uint64
}
