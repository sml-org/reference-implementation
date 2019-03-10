package typesystem

// Type returns the type identifier of the property type
func (i IDProperty) Type() IDType {
	switch i {

	/**** root ****/

	case PR__root_rp:
		return AR__A
	case PR__root_ra:
		return EN__A
	case PR__root_rb:
		return ST__B
	case PR__root_re:
		return NU__E
	case PR__root_rc:
		return ARo__C

	/**** entity A ****/

	case PR__A_ab:
		return STo__B
	case PR__A_ac:
		return ST__C

	/**** struct C ****/

	case PR__C_cb:
		return ST__B
	case PR__C_ce:
		return NU__E
	case PR__C_cs:
		return PT__String
	case PR__C_ct:
		return PT__Text
	case PR__C_cu:
		return PT__Uint32
	case PR__C_ci:
		return PT__Int32

	/**** struct B ****/

	case PR__B_bo:
		return PT__String
	case PR__B_bx:
		return PT__String
	case PR__B_by:
		return PT__Bool
	case PR__B_bz:
		return PTo__Bool
	}

	return IDType(0)
}
