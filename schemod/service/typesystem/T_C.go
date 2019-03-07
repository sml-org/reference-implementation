package typesystem

// T_C implements the type: struct C
type T_C struct {
	cb T_B
	ce T_E
	cs T_String
	ct T_Text
	cu T_Uint32
	ci T_Int32
}

// NewT_C constructs a new instance of struct B
func NewT_C(
	cb T_B,
	ce T_E,
	cs T_String,
	ct T_Text,
	cu T_Uint32,
	ci T_Int32,
) *T_C {
	return &T_C{
		cb: cb,
		ce: ce,
		cs: cs,
		ct: ct,
		cu: cu,
		ci: ci,
	}
}

// ID returns the type identifier
func (o *T_C) ID() IDType {
	return ST__C
}
