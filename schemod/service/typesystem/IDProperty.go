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
