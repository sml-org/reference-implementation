package typesystem

// Property returns the property identifier the parameter is applicable on
func (i IDParameter) Property() IDProperty {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return PR__root_rp

	/**** struct B ****/

	case PM__B_bo_o1:
		return PR__B_bo
	case PM__B_bo_o2:
		return PR__B_bo
	case PM__B_bx_x1:
		return PR__B_bx
	case PM__B_bx_x2:
		return PR__B_bx
	case PM__B_bx_x3:
		return PR__B_bx
	case PM__B_by_y:
		return PR__B_by
	case PM__B_bz_z:
		return PR__B_bz
	}

	return IDProperty(0)
}
