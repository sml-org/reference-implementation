package typesystem

// IDGraphNode represents a graph node identifer
type IDGraphNode uint32

const (
	_ IDGraphNode = iota

	/******** Properties ********/

	/**** graph root ****/

	// GNPR__root_rp = property root.rp
	//   type: Array<A>
	GNPR__root_rp

	// GNPR__root_ra = property root.ra
	//   type: A
	GNPR__root_ra

	// GNPR__root_rb = property root.rb
	//   type: B
	GNPR__root_rb

	// GNPR__root_re = property root.re
	//   type: E
	GNPR__root_re

	// GNPR__root_rc = property root.rc
	//   type: ?Array<C>
	GNPR__root_rc

	/**** entity A ****/

	// GNPR__A_ab = property A.ab
	//   type: ?B
	GNPR__A_ab

	// GNPR__A_ac = property A.ac
	//   type: C
	GNPR__A_ac

	/**** struct C ****/

	// GNPR__C_cb = property C.cb
	//   type: B
	GNPR__C_cb

	// GNPR__C_ce = property C.ce
	//   type: E
	GNPR__C_ce

	// GNPR__C_cs = property C.cs
	//   type: String
	GNPR__C_cs

	// GNPR__C_ct = property C.ct
	//   type: Text
	GNPR__C_ct

	// GNPR__C_cu = property C.cu
	//   type: Uint32
	GNPR__C_cu

	// GNPR__C_ci = property C.ci
	//   type: Int32
	GNPR__C_ci

	/**** struct B ****/

	// GNPR__B_bo = property B.bo
	//   type: String
	GNPR__B_bo

	// GNPR__B_bx = property B.bx
	//   type: String
	GNPR__B_bx

	// GNPR__B_by = property B.by
	//   type: Bool
	GNPR__B_by

	// GNPR__B_bz = property B.bz
	//   type: ?Bool
	GNPR__B_bz

	/******** Attributes ********/

	/**** graph root ****/

	// GNAT__root_rp_version = attribute root.rp:version
	//   type: Version
	GNAT__root_rp_version

	// GNAT__root_rp_length = attribute root.rp:length
	//   type: Uint64
	GNAT__root_rp_length

	// GNAT__root_ra_version = attribute root.ra:version
	//   type: Version
	GNAT__root_ra_version

	// GNAT__root_rb_version = attribute root.rb:version
	//   type: Version
	GNAT__root_rb_version

	// GNAT__root_re_version = attribute root.re:version
	//   type: Version
	GNAT__root_re_version

	// GNAT__root_rc_version = attribute root.rc:version
	//   type: Version
	GNAT__root_rc_version

	// GNAT__root_rc_length = attribute root.rp:length
	//   type: Uint64
	GNAT__root_rc_length

	/**** entity A ****/

	// GNAT__A_ab_version = attribute A.ab:version
	//   type: Version
	GNAT__A_ab_version

	// GNAT__A_ac_version = attribute A.ac:version
	//   type: Version
	GNAT__A_ac_version

	/**** struct C ****/

	// GNAT__C_cb_version = attribute C.cb:version
	//   type: Version
	GNAT__C_cb_version

	// GNAT__C_ce_version = attribute C.ce:version
	//   type: Version
	GNAT__C_ce_version

	// GNAT__C_cs_version = attribute C.cs:version
	//   type: Version
	GNAT__C_cs_version

	// GNAT__C_cs_length = attribute C.cs:length
	//   type: Uint64
	GNAT__C_cs_length

	// GNAT__C_ct_version = attribute C.ct:version
	//   type: Version
	GNAT__C_ct_version

	// GNAT__C_ct_length = attribute C.ct:length
	//   type: Uint64
	GNAT__C_ct_length

	// GNAT__C_cu_version = attribute C.cu:version
	//   type: Version
	GNAT__C_cu_version

	// GNAT__C_ci_version = attribute C.ci:version
	//   type: Version
	GNAT__C_ci_version

	/**** struct B ****/

	// GNAT__B_bo_length = attribute B.bo:length
	//   type: Uint64
	GNAT__B_bo_length

	// GNAT__B_bx_length = attribute B.bx:length
	//   type: Uint64
	GNAT__B_bx_length

	/******** Parameters ********/

	/**** root ****/

	// GNPM__root_rp_page = parameter root.rp.$page
	//   type: AU__root_rp_page
	GNPM__root_rp_page

	/**** struct B ****/

	// GNPM__B_bo_o1 = parameter B.bo.$o1
	//   type: ?Uint32
	GNPM__B_bo_o1

	// GNPM__B_bo_o2 = parameter B.bo.$o2
	//   type: ?String
	GNPM__B_bo_o2

	// GNPM__B_bx_x1 = parameter B.bx.$x1
	//   type: Uint32
	GNPM__B_bx_x1

	// GNPM__B_bx_x2 = parameter B.bx.$x2
	//   type: Text
	GNPM__B_bx_x2

	// GNPM__B_bx_x3 = parameter B.bx.$x3
	//   type: Bool
	GNPM__B_bx_x3

	// GNPM__B_by_y = parameter B.by.$y
	//   type: Bool
	GNPM__B_by_y

	// GNPM__B_bz_z = parameter B.bz.$z
	//   type: ?Bool
	GNPM__B_bz_z
)

// NodeType returns the graph node type of the node identifier
func (id IDGraphNode) NodeType() GraphNodeType {
	if id >= GNPR__root_rp && id <= GNPR__B_bz {
		return Property
	} else if id >= GNAT__root_rp_version && id <= GNAT__B_bx_length {
		return Attribute
	} else if id >= GNPM__root_rp_page && id <= GNPM__B_bz_z {
		return Parameter
	}
	return GraphNodeType(0)
}
