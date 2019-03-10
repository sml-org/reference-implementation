package typesystem

// HasQueryableProperties returns true if the identified type has queryable
// properties
func (i IDType) HasQueryableProperties() bool {
	switch i {
	/**** Structs ****/

	case ST__C:
		return true
	case ST__B:
		return true

	/**** Entities ****/

	case EN__A:
		return true

	/**** Anonymous Structs ****/

	case AS__root_rp_page_struct1:
		return true

	/**** Arrays ****/

	case AR__A:
		return true
	case AR__C:
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
