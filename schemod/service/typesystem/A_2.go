package typesystem

// A_2 implements the anonymous type: struct (root.rp.$page.1)
type A_2 struct {
	Cursor T_ID_A    // cursor ID<A>
	Limit  R_Integer // limit  @(integer)
}

// ID returns the type identifier
func (o A_2) ID() IDType {
	return AS__root_rp_page_struct1
}

// NewA_2 creates a new instance of the anonymous type: struct (root.rp.$page.1)
func NewA_2(
	cursor T_ID_A,
	limit R_Integer,
) *A_2 {
	return &A_2{
		Cursor: cursor,
		Limit:  limit,
	}
}
