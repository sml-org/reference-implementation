package typesystem

// IDType represents a type identifier
type IDType uint32

const (
	_ IDType = iota

	// PT__Root = service graph root
	PT__Root

	// PT__nil = nil type
	PT__nil

	/**************** TYPES ****************/

	/**** Traits ****/

	// TR__integer = integer
	// implemented as: T_Integer
	TR__integer

	/**** Primitives ****/

	// PT__Version = Version
	// implemented as: T_Version
	PT__Version

	// PT__Bool = Bool
	// implemented as: T_Bool
	PT__Bool

	// PTo__Bool = ?Bool
	// implemented as: *T_Bool
	PTo__Bool

	// PT__Int32 = Int32
	// implemented as: T_Int32
	PT__Int32

	// PT__Uint32 = ?Uint32
	// implemented as: T_Uint32
	PT__Uint32

	// PTo__Uint32 = ?Uint32
	// implemented as: *T_Uint32
	PTo__Uint32

	// PT__Int64 = Int64
	// implemented as: T_Int64
	PT__Int64

	// PT__Uint64 = Uint64
	// implemented as: T_Uint64
	PT__Uint64

	// PT__String = String
	// implemented as: T_String
	PT__String

	// PTo__String = ?String
	// implemented as: *T_String
	PTo__String

	// PT__Text = Text
	// implemented as: T_Text
	PT__Text

	/**** Arrays ****/

	// AR__ID_A = Array<ID<A>>
	// implemented as: []T_ID_A
	AR__ID_A

	// AR__A = Array<A>
	// implemented as: []T_A
	AR__A

	// ARo__C = ?Array<C>
	// implemented as: []T_C (nullable)
	ARo__C

	/**** IDs ****/

	// ID__A = ID<A>
	// implemented as: T_ID_A
	ID__A

	/**** Structs ****/

	// ST__C = struct C
	// implemented as: T_C
	ST__C

	// ST__B = struct B
	// implemented as: T_B
	ST__B

	// STo__B = ?struct B
	// implemented as: *T_B
	STo__B

	/**** Enums ****/

	// NU__E = enum E
	// implemented as: T_E
	NU__E

	/**** Entities ****/

	// EN__A = entity A
	// implemented as: T_A
	EN__A

	/**** Unions ****/

	// none

	/**** Anonymous Unions ****/

	// AU__root_rp_page = anonymous union (root.rp.$page)
	// implemented as: A_1
	AU__root_rp_page

	/**** Anonymous Structs ****/

	// AS__root_rp_page_struct1 = anonymous struct (root.rp.$page):struct1
	// implemented as: A_2
	AS__root_rp_page_struct1
)

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
	case PTo__Bool:
		return "?Bool"
	case PT__Int32:
		return "Int32"
	case PT__Uint32:
		return "Uint32"
	case PTo__Uint32:
		return "Uint32"
	case PT__Int64:
		return "Int64"
	case PT__Uint64:
		return "Uint64"
	case PT__String:
		return "String"
	case PTo__String:
		return "?String"
	case PT__Text:
		return "Text"

	/**** Arrays ****/

	case AR__ID_A:
		return "Array<ID<A>>"
	case AR__A:
		return "Array<A>"
	case ARo__C:
		return "?Array<C>"

	/**** IDs ****/

	case ID__A:
		return "ID<A>"

	/**** Structs ****/

	case ST__C:
		return "C"
	case ST__B:
		return "B"
	case STo__B:
		return "?B"

	/**** Enums ****/

	case NU__E:
		return "E"

	/**** Entities ****/

	case EN__A:
		return "A"

	/**** Unions ****/

	// none

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		return "union(root.rp.$page)"

	/**** Anonymous Structs ****/

	case AS__root_rp_page_struct1:
		return "struct(root.rp.$page.1)"
	}

	return "<undefined>"
}

// IsComposite returns true if the identified type is a composite type
func (i IDType) IsComposite() bool {
	switch i {
	/**** Structs ****/

	case ST__C:
		return true
	case ST__B:
		return true
	case STo__B:
		return true

	/**** Entities ****/

	case EN__A:
		return true

	/**** Anonymous Structs ****/

	case AS__root_rp_page_struct1:
		return true
	}

	return false
}

// IsPolymorphic returns true if the identified type is a polymorphic type
func (i IDType) IsPolymorphic() bool {
	switch i {
	/**** Traits ****/

	case TR__integer:
		return true

	/**** Primitives ****/
	case PTo__Bool:
		return true
	case PTo__Uint32:
		return true
	case PTo__String:
		return true

	/**** Arrays ****/

	case ARo__C:
		return true

	/**** IDs ****/

	// no polymorphic types

	/**** Structs ****/

	case STo__B:
		return true

	/**** Enums ****/

	// no polymorphic types

	/**** Entities ****/

	// no polymorphic types

	/**** Unions ****/

	// none

	/**** Anonymous Unions ****/

	case AU__root_rp_page:
		return true

		/**** Anonymous Structs ****/

		// no polymorphic types
	}

	return false
}
