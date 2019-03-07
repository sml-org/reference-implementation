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

// String returns the textual name representation
func (i IDParameter) String() string {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return "root.rp.$page"

	/**** struct B ****/

	case PM__B_bo_o1:
		return "B.bo.$o1"
	case PM__B_bo_o2:
		return "B.bo.$o2"
	case PM__B_bx_x1:
		return "B.bx.$x1"
	case PM__B_bx_x2:
		return "B.bx.$x2"
	case PM__B_bx_x3:
		return "B.bx.$x3"
	case PM__B_by_y:
		return "B.by.$y"
	case PM__B_bz_z:
		return "B.bz.$z"
	}

	return "<undefined>"
}

// Property returns the property identifier the parameter is applicable on
func (i IDParameter) Property() IDProperty {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return PR__root_rp

	/**** struct B ****/

	case PM__B_bo_o1:
		return PR__B_bo
	case PM__B_bo_o2:
		return PR__B_bo
	case PM__B_bx_x1:
		return PR__B_bx
	case PM__B_bx_x2:
		return PR__B_bx
	case PM__B_bx_x3:
		return PR__B_bx
	case PM__B_by_y:
		return PR__B_by
	case PM__B_bz_z:
		return PR__B_bz
	}

	return IDProperty(0)
}

// Type returns the property identifier the parameter is applicable on
func (i IDParameter) Type() IDType {
	switch i {
	/**** root ****/

	case PM__root_rp_page:
		return AU__root_rp_page

	/**** struct B ****/

	case PM__B_bo_o1:
		return PTo__Uint32
	case PM__B_bo_o2:
		return PTo__String
	case PM__B_bx_x1:
		return PT__Uint32
	case PM__B_bx_x2:
		return PT__Text
	case PM__B_bx_x3:
		return PT__Bool
	case PM__B_by_y:
		return PT__Bool
	case PM__B_bz_z:
		return PTo__Bool
	}

	return IDType(0)
}
