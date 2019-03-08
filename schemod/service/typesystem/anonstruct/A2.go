package anonstruct

import "reference-implementation/schemod/service/typesystem"

// A2 reflects the type anonymous struct (root.rp.$page.1)
type A2 = typesystem.A_2

// IDA2 aliases the type identifier of anonymous struct (root.rp.$page.1)
const IDA2 = typesystem.AS__root_rp_page_struct1

// NewA2 creates a new instance of A2
func NewA2(
	cursor typesystem.T_ID_A,
	limit typesystem.R_Integer,
) *A2 {
	return typesystem.NewA_2(
		cursor,
		limit,
	)
}
