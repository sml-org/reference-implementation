package typesystem

// Type returns the property identifier the parameter is applicable on
func (i IDParameter) Type() IDType {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return AU__root_rp_page

	/**** struct B ****/

	case PM__B_bo_o1:
		return PTo__Uint32
	case PM__B_bo_o2:
		return PTo__String
	case PM__B_bx_x1:
		return PT__Uint32
	case PM__B_bx_x2:
		return PT__Text
	case PM__B_bx_x3:
		return PT__Bool
	case PM__B_by_y:
		return PT__Bool
	case PM__B_bz_z:
		return PTo__Bool
	}

	return IDType(0)
}
