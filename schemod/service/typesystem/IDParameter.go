package typesystem

// IDParameter represents a parameter identifier
type IDParameter uint32

const (
	_ IDParameter = iota

	/**** root ****/

	// PM__root_rp_page = root.rp.$page
	//   type: AU__root_rp_page
	PM__root_rp_page

	/**** struct B ****/

	// PM__B_bo_o1 = B.bo.$o1
	//   type: ?Uint32
	PM__B_bo_o1

	// PM__B_bo_o2 = B.bo.$o2
	//   type: ?String
	PM__B_bo_o2

	// PM__B_bx_x1 = B.bx.$x1
	//   type: Uint32
	PM__B_bx_x1

	// PM__B_bx_x2 = B.bx.$x2
	//   type: Text
	PM__B_bx_x2

	// PM__B_bx_x3 = B.bx.$x3
	//   type: Bool
	PM__B_bx_x3

	// PM__B_by_y  = B.by.$y
	//   type: Bool
	PM__B_by_y

	// PM__B_bz_z  = B.bz.$z
	//   type: ?Bool
	PM__B_bz_z
)
