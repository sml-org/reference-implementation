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
