package typesystem

// String returns the textual name representation
func (i IDType) String() string {
	switch i {

	case PT__Root:
		return "root"
	case PT__nil:
		return "nil"

	/**** Traits ****/

	case TR__integer:
		return "integer"

	/**** Primitives ****/
	case PT__Version:
		return "Version"
	case PT__Bool:
		return "Bool"
	case PT__Int32:
		return "Int32"
	case PT__Uint32:
		return "Uint32"
	case PT__Int64:
		return "Int64"
	case PT__Uint64:
		return "Uint64"
	case PT__String:
		return "String"
	case PT__Text:
		return "Text"

	/**** IDs ****/

	case ID__A:
		return "ID<A>"

	/**** Structs ****/

	case ST__C:
		return "C"
	case ST__B:
		return "B"

	/**** Enums ****/

	case NU__E:
		return "E"

	/**** Entities ****/

	case EN__A:
		return "A"

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		return "union(root.rp.$page)"

	/**** Anonymous Structs ****/

	case AS__root_rp_page_struct1:
		return "struct(root.rp.$page.1)"

	/**** Arrays ****/

	case AR__ID_A:
		return "Array<ID<A>>"
	case AR__A:
		return "Array<A>"
	case AR__C:
		return "Array<C>"

	/**** Optional Primitives ****/

	case PTo__Bool:
		return "?Bool"
	case PTo__Uint32:
		return "?Uint32"
	case PTo__String:
		return "?String"

	/**** Optional Structs ****/

	case STo__B:
		return "?B"

	/**** Optional Arrays ****/

	case ARo__C:
		return "?Array<C>"

	}

	return "<undefined>"
}
