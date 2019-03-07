package typesystem

// IDProperty represents a property identifier
type IDProperty uint32

const (
	_ IDProperty = iota

	/**** graph root ****/

	// PR__root_rp = root.rp
	//   type: Array<A>
	PR__root_rp

	// PR__root_ra = root.ra
	//   type: A
	PR__root_ra

	// PR__root_rb = root.rb
	//   type: B
	PR__root_rb

	// PR__root_re = root.re
	//   type: E
	PR__root_re

	// PR__root_rc = root.rc
	//   type: ?Array<C>
	PR__root_rc

	/**** entity A ****/

	// PR__A_ab = A.ab
	//   type: ?B
	PR__A_ab

	// PR__A_ac = A.ac
	//   type: C
	PR__A_ac

	/**** struct C ****/

	// PR__C_cb = C.cb
	//   type: B
	PR__C_cb

	// PR__C_ce = C.ce
	//   type: E
	PR__C_ce

	// PR__C_cs = C.cs
	//   type: String
	PR__C_cs

	// PR__C_ct = C.ct
	//   type: Text
	PR__C_ct

	// PR__C_cu = C.cu
	//   type: Uint32
	PR__C_cu

	// PR__C_ci = C.ci
	//   type: Int32
	PR__C_ci

	/**** struct B ****/

	// PR__B_bo = B.bo
	//   type: String
	PR__B_bo

	// PR__B_bx = B.bx
	//   type: String
	PR__B_bx

	// PR__B_by = B.by
	//   type: Bool
	PR__B_by

	// PR__B_bz = B.bz
	//   type: ?Bool
	PR__B_bz
)

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
