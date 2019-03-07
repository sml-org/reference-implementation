package typesystem

// T_Uint32 implements the type: Uint32
type T_Uint32 struct {
	v uint32
}

// ID returns the type identifier
func (o T_Uint32) ID() IDType {
	return PT__Uint32
}
