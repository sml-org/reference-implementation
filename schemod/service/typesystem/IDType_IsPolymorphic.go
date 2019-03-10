package typesystem

// IsPolymorphic returns true if the identified type is a polymorphic type
func (i IDType) IsPolymorphic() bool {
	switch i {
	/**** Traits ****/

	case TR__integer:
		return true

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		return true

	/**** Optional Primitives ****/

	case PTo__Bool:
		return true
	case PTo__Uint32:
		return true
	case PTo__String:
		return true

	/**** Optional Structs ****/

	case STo__B:
		return true

	/**** Optional Arrays ****/

	case ARo__C:
		return true
	}

	return false
}
