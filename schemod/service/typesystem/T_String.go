package typesystem

// T_String implements the type: String
type T_String struct {
	v []byte
}

// ID returns the type identifier
func (o T_String) ID() IDType {
	return PT__String
}
