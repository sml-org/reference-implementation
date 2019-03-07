package typesystem

// IDAttribute represents an attribute identifier
type IDAttribute uint16

const (
	_ IDAttribute = iota

	/**** graph root ****/

	// AT__root_rp_version = root.rp:version
	//   type: Version
	AT__root_rp_version

	// AT__root_rp_length = root.rp:length
	//   type: Uint64
	AT__root_rp_length

	// AT__root_ra_version = root.ra:version
	//   type: Version
	AT__root_ra_version

	// AT__root_rb_version = root.rb:version
	//   type: Version
	AT__root_rb_version

	// AT__root_re_version = root.re:version
	//   type: Version
	AT__root_re_version

	// AT__root_rc_version = root.rc:version
	//   type: Version
	AT__root_rc_version

	// AT__root_rc_length = root.rp:length
	//   type: Uint64
	AT__root_rc_length

	/**** entity A ****/

	// AT__A_ab_version = A.ab:version
	//   type: Version
	AT__A_ab_version

	// AT__A_ac_version = A.ac:version
	//   type: Version
	AT__A_ac_version

	/**** struct C ****/

	// AT__C_cb_version = C.cb:version
	//   type: Version
	AT__C_cb_version

	// AT__C_ce_version = C.ce:version
	//   type: Version
	AT__C_ce_version

	// AT__C_cs_version = C.cs:version
	//   type: Version
	AT__C_cs_version

	// AT__C_cs_length = C.cs:length
	//   type: Uint64
	AT__C_cs_length

	// AT__C_ct_version = C.ct:version
	//   type: Version
	AT__C_ct_version

	// AT__C_ct_length = C.ct:length
	//   type: Uint64
	AT__C_ct_length

	// AT__C_cu_version = C.cu:version
	//   type: Version
	AT__C_cu_version

	// AT__C_ci_version = C.ci:version
	//   type: Version
	AT__C_ci_version

	/**** struct B ****/

	// AT__B_bo_length = B.bo:length
	//   type: Uint64
	AT__B_bo_length

	// AT__B_bx_length = B.bx:length
	//   type: Uint64
	AT__B_bx_length
)

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
