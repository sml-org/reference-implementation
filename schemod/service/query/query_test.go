package query_test

import (
	"reference-implementation/schemod/service/query"
	"reference-implementation/schemod/service/typesystem"
	"reference-implementation/schemod/service/typesystem/anonstruct"
	"reference-implementation/schemod/service/typesystem/array"
	"reference-implementation/schemod/service/typesystem/id"
	"reference-implementation/schemod/service/typesystem/prim"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestQueryValidation tests a successful validation
func TestQueryValidation(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: &array.ID_A{
						Items: []id.A{
							id.A{},
						},
					},
				},
			},
		},
	}
	require.NoError(t, q.Validate())
}

// TestWrongArgument tests applying an argument on the wrong property
func TestWrongArgument(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__B_bo_o1: &array.ID_A{
						Items: []id.A{
							id.A(typesystem.NewID()),
						},
					},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

// TestPolymorphicParameterWrongType tests assigning an unassignable type
// to a parameter of polymorphic type
func TestPolymorphicParameterWrongType(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: &prim.Bool{Value: true},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

// TestPolymorphicParameter tests parameters of polymorphic type to accept
// assignable types
func TestPolymorphicParameter(t *testing.T) {
	assignableTypes := []typesystem.Serializable{
		&anonstruct.A2{},
		&array.ID_A{Items: []id.A{
			id.A(typesystem.NewID()),
		}},
	}
	for _, val := range assignableTypes {
		q := query.Query{
			Props: query.Props{
				typesystem.PR__root_rp: query.PropertySelection{
					Args: query.Args{
						typesystem.PM__root_rp_page: val,
					},
				},
			},
		}
		require.NoError(t, q.Validate())
	}
}

// TestParameterMismatchingType tests providing a mismatching type as
// an argument value
func TestParameterMismatchingType(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: &prim.Bool{Value: true},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

// TestInexistentProperty tests selecting inexistent property
func TestInexistentProperty(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__A_ab: query.PropertySelection{},
		},
	}
	require.Error(t, q.Validate())
}
