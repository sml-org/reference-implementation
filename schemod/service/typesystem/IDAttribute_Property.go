package typesystem

// Property returns the property identifier the attribute belongs to
func (i IDAttribute) Property() IDProperty {
	switch i {
	/**** graph root ****/

	case AT__root_rp_version:
		return PR__root_rp
	case AT__root_rp_length:
		return PR__root_rp
	case AT__root_ra_version:
		return PR__root_ra
	case AT__root_rb_version:
		return PR__root_rb
	case AT__root_re_version:
		return PR__root_re
	case AT__root_rc_version:
		return PR__root_rc
	case AT__root_rc_length:
		return PR__root_rc

	/**** entity A ****/

	case AT__A_ab_version:
		return PR__A_ab
	case AT__A_ac_version:
		return PR__A_ac

	/**** struct C ****/

	case AT__C_cb_version:
		return PR__C_cb
	case AT__C_ce_version:
		return PR__C_ce
	case AT__C_cs_version:
		return PR__C_cs
	case AT__C_cs_length:
		return PR__C_cs
	case AT__C_ct_version:
		return PR__C_ct
	case AT__C_ct_length:
		return PR__C_ct
	case AT__C_cu_version:
		return PR__C_cu
	case AT__C_ci_version:
		return PR__C_ci

	/**** struct B ****/

	case AT__B_bo_length:
		return PR__B_bo
	case AT__B_bx_length:
		return PR__B_bx
	}

	return IDProperty(0)
}
