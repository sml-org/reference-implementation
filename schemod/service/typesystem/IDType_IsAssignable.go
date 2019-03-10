package typesystem

// IsAssignable returns true if the given type is accepted
// by the identified polymorphic type
func (i IDType) IsAssignable(t IDType) bool {
	switch i {
	/**** Traits ****/

	case TR__integer:
		switch t {
		case PT__Int32:
			return true
		case PT__Int64:
			return true
		case PT__Uint32:
			return true
		case PT__Uint64:
			return true
		}

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		switch t {
		case AR__ID_A:
			return true
		case AS__root_rp_page_struct1:
			return true
		}

	/**** Optional Primitives ****/

	case PTo__Bool:
		switch t {
		case PT__Bool:
			return true
		case PT__nil:
			return true
		}
	case PTo__Uint32:
		switch t {
		case PT__Uint32:
			return true
		case PT__nil:
			return true
		}
	case PTo__String:
		switch t {
		case PT__String:
			return true
		case PT__nil:
			return true
		}

	/**** Optional Structs ****/

	case STo__B:
		switch t {
		case ST__B:
			return true
		case PT__nil:
			return true
		}

	/**** Optional Arrays ****/

	case ARo__C:
		switch t {
		case AR__C:
			return true
		case PT__nil:
			return true
		}
	}

	return false
}
