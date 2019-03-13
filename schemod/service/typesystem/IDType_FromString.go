package typesystem

import "fmt"

// FromString initializes the value from a string
func (o *IDType) FromString(b string) error {
	switch b {

	case "root":
		*o = PT__Root

	case "nil":
		*o = PT__nil

	/**** Traits ****/

	case "integer":
		*o = TR__integer

	/**** Primitives ****/
	case "Version":
		*o = PT__Version

	case "Bool":
		*o = PT__Bool

	case "Int32":
		*o = PT__Int32

	case "Uint32":
		*o = PT__Uint32

	case "Int64":
		*o = PT__Int64

	case "Uint64":
		*o = PT__Uint64

	case "String":
		*o = PT__String

	case "Text":
		*o = PT__Text

	/**** IDs ****/

	case "ID<A>":
		*o = ID__A

	/**** Structs ****/

	case "C":
		*o = ST__C

	case "B":
		*o = ST__B

	/**** Enums ****/

	case "E":
		*o = NU__E

	/**** Entities ****/

	case "A":
		*o = EN__A

	/**** Anonymous Unions ****/

	case "union(root.rp.$page)":
		*o = AU__root_rp_page

	/**** Anonymous Structs ****/

	case "struct(root.rp.$page.1)":
		*o = AS__root_rp_page_struct1

	/**** Arrays ****/

	case "Array<ID<A>>":
		*o = AR__ID_A

	case "Array<A>":
		*o = AR__A

	case "Array<C>":
		*o = AR__C

	/**** Optional Primitives ****/

	case "?Bool":
		*o = PTo__Bool

	case "?Uint32":
		*o = PTo__Uint32

	case "?String":
		*o = PTo__String

	/**** Optional Structs ****/

	case "?B":
		*o = STo__B

	/**** Optional Arrays ****/

	case "?Array<C>":
		*o = ARo__C

	}

	return fmt.Errorf("unknown IDType value: %s", b)
}
