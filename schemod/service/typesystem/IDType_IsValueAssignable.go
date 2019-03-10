package typesystem

// IsValueAssignable returns true if v is of an assignable type,
// otherwise returns false
func (i IDType) IsValueAssignable(v interface{}) bool {
	switch i {
	/**** Traits ****/

	case TR__integer:
		switch v.(type) {
		case T_Int32:
			return true
		case T_Int64:
			return true
		case T_Uint32:
			return true
		case T_Uint64:
			return true
		}

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		switch v.(type) {
		case []T_ID_A:
			// is of type Array<ID<A>>
			return true
		case A_2:
			// is of type anonymous struct root.rp.$page:struct1
			return true
		}

	/**** Optional Primitives ****/

	case PTo__Bool:
		switch v.(type) {
		case T_Bool:
			return true
		case nil:
			return true
		}
	case PTo__Uint32:
		switch v.(type) {
		case T_Uint32:
			return true
		case nil:
			return true
		}
	case PTo__String:
		switch v.(type) {
		case T_String:
			return true
		case nil:
			return true
		}

	/**** Optional Structs ****/

	case STo__B:
		switch v.(type) {
		case T_B:
			return true
		case nil:
			return true
		}

	/**** Optional Arrays ****/

	case ARo__C:
		switch v.(type) {
		case T_C:
			return true
		case nil:
			return true
		}
	}

	return false
}
