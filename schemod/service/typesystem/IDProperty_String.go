package typesystem

// String returns the textual name representation
func (i IDProperty) String() string {
	switch i {
	/**** graph root ****/

	case PR__root_rp:
		return "root.rp"
	case PR__root_ra:
		return "root.ra"
	case PR__root_rb:
		return "root.rb"
	case PR__root_re:
		return "root.re"
	case PR__root_rc:
		return "root.rc"

	/**** entity A ****/

	case PR__A_ab:
		return "A.ab"
	case PR__A_ac:
		return "A.ac"

	/**** struct C ****/

	case PR__C_cb:
		return "C.cb"
	case PR__C_ce:
		return "C.ce"
	case PR__C_cs:
		return "C.cs"
	case PR__C_ct:
		return "C.ct"
	case PR__C_cu:
		return "C.cu"
	case PR__C_ci:
		return "C.ci"

	/**** struct B ****/

	case PR__B_bo:
		return "B.bo"
	case PR__B_bx:
		return "B.bx"
	case PR__B_by:
		return "B.by"
	case PR__B_bz:
		return "B.bz"
	}

	return "<undefined>"
}
