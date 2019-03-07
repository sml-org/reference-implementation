package typesystem

// T_E represents the value type of enum E
type T_E uint32

const (
	_ T_E = iota

	// T_E_e1 reflects E::e1
	T_E_e1

	// T_E_e2 reflects E::e2
	T_E_e2

	// T_E_e3 reflects E::e3
	T_E_e3

	// T_E_e4 reflects E::e4
	T_E_e4
)

// String returns the string
func (o T_E) String() string {
	switch o {
	case T_E_e1:
		return "E::e1"
	case T_E_e2:
		return "E::e2"
	case T_E_e3:
		return "E::e3"
	case T_E_e4:
		return "E::e4"
	}
	return ""
}

// ID returns the type identifier
func (o T_E) ID() IDType {
	return NU__E
}
