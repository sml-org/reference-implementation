package typesystem

// String returns the textual name representation
func (i IDAttribute) String() string {
	switch i {
	/**** graph root ****/

	case AT__root_rp_version:
		return "root.rp:version"
	case AT__root_rp_length:
		return "root.rp:length"
	case AT__root_ra_version:
		return "root.ra:version"
	case AT__root_rb_version:
		return "root.rb:version"
	case AT__root_re_version:
		return "root.re:version"
	case AT__root_rc_version:
		return "root.rc:version"
	case AT__root_rc_length:
		return "root.rp:length"

	/**** entity A ****/

	case AT__A_ab_version:
		return "A.ab:version"
	case AT__A_ac_version:
		return "A.ac:version"

	/**** struct C ****/

	case AT__C_cb_version:
		return "C.cb:version"
	case AT__C_ce_version:
		return "C.ce:version"
	case AT__C_cs_version:
		return "C.cs:version"
	case AT__C_cs_length:
		return "C.cs:length"
	case AT__C_ct_version:
		return "C.ct:version"
	case AT__C_ct_length:
		return "C.ct:length"
	case AT__C_cu_version:
		return "C.cu:version"
	case AT__C_ci_version:
		return "C.ci:version"

	/**** struct B ****/

	case AT__B_bo_length:
		return "B.bo:length"
	case AT__B_bx_length:
		return "B.bx:length"
	}

	return "<undefined>"
}
