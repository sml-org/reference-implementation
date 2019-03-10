package typesystem

// Type returns the type identifier of the attribute type
func (i IDAttribute) Type() IDType {
	switch i {
	/**** graph root ****/

	case AT__root_rp_version:
		return PT__Version
	case AT__root_rp_length:
		return PT__Uint32
	case AT__root_ra_version:
		return PT__Version
	case AT__root_rb_version:
		return PT__Version
	case AT__root_re_version:
		return PT__Version
	case AT__root_rc_version:
		return PT__Version
	case AT__root_rc_length:
		return PT__Uint32

	/**** entity A ****/

	case AT__A_ab_version:
		return PT__Version
	case AT__A_ac_version:
		return PT__Version

	/**** struct C ****/

	case AT__C_cb_version:
		return PT__Version
	case AT__C_ce_version:
		return PT__Version
	case AT__C_cs_version:
		return PT__Version
	case AT__C_cs_length:
		return PT__Uint32
	case AT__C_ct_version:
		return PT__Version
	case AT__C_ct_length:
		return PT__Uint32
	case AT__C_cu_version:
		return PT__Version
	case AT__C_ci_version:
		return PT__Version

	/**** struct B ****/

	case AT__B_bo_length:
		return PT__Uint32
	case AT__B_bx_length:
		return PT__Uint32
	}

	return IDType(0)
}
