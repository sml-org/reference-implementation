package typesystem

// IDAttribute represents an attribute identifier
type IDAttribute uint16

const (
	/**** graph root ****/

	// AT__root_rp_version = root.rp:version
	//   type: Version
	AT__root_rp_version IDAttribute = IDAttribute(GNAT__root_rp_version)

	// AT__root_rp_length = root.rp:length
	//   type: Uint64
	AT__root_rp_length IDAttribute = IDAttribute(GNAT__root_rp_length)

	// AT__root_ra_version = root.ra:version
	//   type: Version
	AT__root_ra_version IDAttribute = IDAttribute(GNAT__root_ra_version)

	// AT__root_rb_version = root.rb:version
	//   type: Version
	AT__root_rb_version IDAttribute = IDAttribute(GNAT__root_rb_version)

	// AT__root_re_version = root.re:version
	//   type: Version
	AT__root_re_version IDAttribute = IDAttribute(GNAT__root_re_version)

	// AT__root_rc_version = root.rc:version
	//   type: Version
	AT__root_rc_version IDAttribute = IDAttribute(GNAT__root_rc_version)

	// AT__root_rc_length = root.rp:length
	//   type: Uint64
	AT__root_rc_length IDAttribute = IDAttribute(GNAT__root_rc_length)

	/**** entity A ****/

	// AT__A_ab_version = A.ab:version
	//   type: Version
	AT__A_ab_version IDAttribute = IDAttribute(GNAT__A_ab_version)

	// AT__A_ac_version = A.ac:version
	//   type: Version
	AT__A_ac_version IDAttribute = IDAttribute(GNAT__A_ac_version)

	/**** struct C ****/

	// AT__C_cb_version = C.cb:version
	//   type: Version
	AT__C_cb_version IDAttribute = IDAttribute(GNAT__C_cb_version)

	// AT__C_ce_version = C.ce:version
	//   type: Version
	AT__C_ce_version IDAttribute = IDAttribute(GNAT__C_ce_version)

	// AT__C_cs_version = C.cs:version
	//   type: Version
	AT__C_cs_version IDAttribute = IDAttribute(GNAT__C_cs_version)

	// AT__C_cs_length = C.cs:length
	//   type: Uint64
	AT__C_cs_length IDAttribute = IDAttribute(GNAT__C_cs_length)

	// AT__C_ct_version = C.ct:version
	//   type: Version
	AT__C_ct_version IDAttribute = IDAttribute(GNAT__C_ct_version)

	// AT__C_ct_length = C.ct:length
	//   type: Uint64
	AT__C_ct_length IDAttribute = IDAttribute(GNAT__C_ct_length)

	// AT__C_cu_version = C.cu:version
	//   type: Version
	AT__C_cu_version IDAttribute = IDAttribute(GNAT__C_cu_version)

	// AT__C_ci_version = C.ci:version
	//   type: Version
	AT__C_ci_version IDAttribute = IDAttribute(GNAT__C_ci_version)

	/**** struct B ****/

	// AT__B_bo_length = B.bo:length
	//   type: Uint64
	AT__B_bo_length IDAttribute = IDAttribute(GNAT__B_bo_length)

	// AT__B_bx_length = B.bx:length
	//   type: Uint64
	AT__B_bx_length IDAttribute = IDAttribute(GNAT__B_bx_length)
)
