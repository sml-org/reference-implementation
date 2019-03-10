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
	// implementation: T_Integer
	// categories:     polymorphic
	// length (bytes): not serializable
	TR__integer

	/**** Primitives ****/

	// PT__Version = Version
	// implementation: T_Version
	// categories:     scalar
	// length (bytes): 16
	PT__Version

	// PT__Bool = Bool
	// implementation: T_Bool
	// categories:     scalar, instanciable
	// length (bytes): 1
	PT__Bool

	// PT__Int32 = Int32
	// implementation: T_Int32
	// categories:     scalar, instanciable
	// length (bytes): 4
	PT__Int32

	// PT__Uint32 = ?Uint32
	// implementation: T_Uint32
	// categories:     scalar, instanciable
	// length (bytes): 4
	PT__Uint32

	// PT__Int64 = Int64
	// implementation: T_Int64
	// categories:     scalar, instanciable
	// length (bytes): 4
	PT__Int64

	// PT__Uint64 = Uint64
	// implementation: T_Uint64
	// categories:     scalar, instanciable
	// length (bytes): 4
	PT__Uint64

	// PT__String = String
	// implementation: T_String
	// length (bytes): dynamic
	PT__String

	// PT__Text = Text
	// implementation: T_Text
	// categories:     scalar, instanciable
	// length (bytes): dynamic
	PT__Text

	/**** IDs ****/

	// ID__A = ID<A>
	// implementation: T_ID_A
	// categories:     scalar
	// length (bytes): 16
	ID__A

	/**** Structs ****/

	// ST__C = struct C
	// implementation: T_C
	// categories:     composite, instanciable, queryable
	// length (bytes): dynamic
	ST__C

	// ST__B = struct B
	// implementation: T_B
	// categories:     composite, instanciable, queryable
	// length (bytes): dynamic
	ST__B

	/**** Enums ****/

	// NU__E = enum E
	// implementation: T_E
	// categories:     scalar, instanciable
	// length (bytes): 4
	NU__E

	/**** Entities ****/

	// EN__A = entity A
	// implementation: T_A
	// categories:     composite, queryable
	// length (bytes): dynamic
	EN__A

	/**** Anonymous Unions ****/

	// AU__root_rp_page = anonymous union (root.rp.$page)
	// implementation: A_1
	// categories:     polymorphic
	// length (bytes): not serializable
	AU__root_rp_page

	/**** Anonymous Structs ****/

	// AS__root_rp_page_struct1 = anonymous struct (root.rp.$page):struct1
	// implementation: A_2
	// categories:     composite, instanciable, queryable
	// length (bytes): 20
	AS__root_rp_page_struct1

	/**** Arrays ****/

	// AR__ID_A = Array<ID<A>>
	// implementation: []T_ID_A
	// categories:     instanciable
	// length (bytes): dynamic(n * 16)
	AR__ID_A

	// AR__A = Array<A>
	// implementation: []T_A
	// categories:     instanciable, queryable
	// length (bytes): dynamic
	AR__A

	// AR__C = Array<C>
	// implementation: []T_C
	// categories:     instanciable, queryable
	// length (bytes): dynamic
	AR__C

	/**** Optional Primitives ****/

	// PTo__Bool = ?Bool
	// implementation: *T_Bool
	// categories:     polymorphic
	// length (bytes): not serializable
	PTo__Bool

	// PTo__Uint32 = ?Uint32
	// implementation: *T_Uint32
	// categories:     polymorphic
	// length (bytes): not serializable
	PTo__Uint32

	// PTo__String = ?String
	// implementation: *T_String
	// categories:     polymorphic
	// length (bytes): not serializable
	PTo__String

	/**** Optional Structs ****/

	// STo__B = ?struct B
	// implementation: *T_B
	// categories:     polymorphic, queryable
	// length (bytes): not serializable
	STo__B

	/**** Optional Arrays ****/

	// ARo__C = ?Array<C>
	// implementation: []T_C (nullable)
	// categories:     polymorphic, queryable
	// length (bytes): not serializable
	ARo__C
)
