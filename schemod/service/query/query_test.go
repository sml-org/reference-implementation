package query_test

import (
	"reference-implementation/schemod/service/query"
	"reference-implementation/schemod/service/typesystem"
	"reference-implementation/schemod/service/typesystem/anonstruct"
	"reference-implementation/schemod/service/typesystem/id"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestQueryValidation tests a successful validation
func TestQueryValidation(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: query.Argument{
						Type: typesystem.AR__ID_A,
						Value: []typesystem.T_ID_A{
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
					typesystem.PM__B_bo_o1: query.Argument{
						Type: typesystem.AR__ID_A,
						Value: []id.A{
							typesystem.NewT_ID_A(typesystem.NewID()),
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
					typesystem.PM__root_rp_page: query.Argument{
						Type:  typesystem.PT__Bool,
						Value: typesystem.NewT_Bool(true),
					},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

// TestPolymorphicParameter tests parameters of polymorphic type to accept
// assignable types
func TestPolymorphicParameter(t *testing.T) {
	assignableTypes := map[typesystem.IDType]interface{}{
		typesystem.AS__root_rp_page_struct1: anonstruct.A2{},
		typesystem.AR__ID_A: []id.A{
			typesystem.NewT_ID_A(typesystem.NewID()),
		},
	}
	for tp, val := range assignableTypes {
		q := query.Query{
			Props: query.Props{
				typesystem.PR__root_rp: query.PropertySelection{
					Args: query.Args{
						typesystem.PM__root_rp_page: query.Argument{
							Type:  tp,
							Value: val,
						},
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
					typesystem.PM__root_rp_page: query.Argument{
						Type:  typesystem.AR__ID_A,
						Value: typesystem.NewT_Bool(true),
					},
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
