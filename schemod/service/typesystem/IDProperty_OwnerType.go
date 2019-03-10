package typesystem

// OwnerType returns the type identifier of the type that owns the property
func (i IDProperty) OwnerType() IDType {
	switch i {

	/**** root ****/

	case PR__root_rp:
		return PT__Root
	case PR__root_ra:
		return PT__Root
	case PR__root_rb:
		return PT__Root
	case PR__root_re:
		return PT__Root
	case PR__root_rc:
		return PT__Root

	/**** entity A ****/

	case PR__A_ab:
		return EN__A
	case PR__A_ac:
		return EN__A

	/**** struct C ****/

	case PR__C_cb:
		return ST__C
	case PR__C_ce:
		return ST__C
	case PR__C_cs:
		return ST__C
	case PR__C_ct:
		return ST__C
	case PR__C_cu:
		return ST__C
	case PR__C_ci:
		return ST__C

	/**** struct B ****/

	case PR__B_bo:
		return ST__B
	case PR__B_bx:
		return ST__B
	case PR__B_by:
		return ST__B
	case PR__B_bz:
		return ST__B
	}

	return IDType(0)
}
