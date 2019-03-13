package typesystem

// IDProperty represents a property identifier
type IDProperty uint32

const (
	/**** graph root ****/

	// PR__root_rp = root.rp
	//   type: Array<A>
	PR__root_rp IDProperty = IDProperty(GNPR__root_rp)

	// PR__root_ra = root.ra
	//   type: A
	PR__root_ra IDProperty = IDProperty(GNPR__root_ra)

	// PR__root_rb = root.rb
	//   type: B
	PR__root_rb IDProperty = IDProperty(GNPR__root_rb)

	// PR__root_re = root.re
	//   type: E
	PR__root_re IDProperty = IDProperty(GNPR__root_re)

	// PR__root_rc = root.rc
	//   type: ?Array<C>
	PR__root_rc IDProperty = IDProperty(GNPR__root_rc)

	/**** entity A ****/

	// PR__A_ab = A.ab
	//   type: ?B
	PR__A_ab IDProperty = IDProperty(GNPR__A_ab)

	// PR__A_ac = A.ac
	//   type: C
	PR__A_ac IDProperty = IDProperty(GNPR__A_ac)

	/**** struct C ****/

	// PR__C_cb = C.cb
	//   type: B
	PR__C_cb IDProperty = IDProperty(GNPR__C_cb)

	// PR__C_ce = C.ce
	//   type: E
	PR__C_ce IDProperty = IDProperty(GNPR__C_ce)

	// PR__C_cs = C.cs
	//   type: String
	PR__C_cs IDProperty = IDProperty(GNPR__C_cs)

	// PR__C_ct = C.ct
	//   type: Text
	PR__C_ct IDProperty = IDProperty(GNPR__C_ct)

	// PR__C_cu = C.cu
	//   type: Uint32
	PR__C_cu IDProperty = IDProperty(GNPR__C_cu)

	// PR__C_ci = C.ci
	//   type: Int32
	PR__C_ci IDProperty = IDProperty(GNPR__C_ci)

	/**** struct B ****/

	// PR__B_bo = B.bo
	//   type: String
	PR__B_bo IDProperty = IDProperty(GNPR__B_bo)

	// PR__B_bx = B.bx
	//   type: String
	PR__B_bx IDProperty = IDProperty(GNPR__B_bx)

	// PR__B_by = B.by
	//   type: Bool
	PR__B_by IDProperty = IDProperty(GNPR__B_by)

	// PR__B_bz = B.bz
	//   type: ?Bool
	PR__B_bz IDProperty = IDProperty(GNPR__B_bz)
)
