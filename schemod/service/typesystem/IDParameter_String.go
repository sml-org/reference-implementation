package typesystem

// String returns the textual name representation
func (i IDParameter) String() string {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return "root.rp.$page"

	/**** struct B ****/

	case PM__B_bo_o1:
		return "B.bo.$o1"
	case PM__B_bo_o2:
		return "B.bo.$o2"
	case PM__B_bx_x1:
		return "B.bx.$x1"
	case PM__B_bx_x2:
		return "B.bx.$x2"
	case PM__B_bx_x3:
		return "B.bx.$x3"
	case PM__B_by_y:
		return "B.by.$y"
	case PM__B_bz_z:
		return "B.bz.$z"
	}

	return "<undefined>"
}
