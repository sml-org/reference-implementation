package query_test

import (
	"reference-implementation/schemod/service/query"
	"reference-implementation/schemod/service/typesystem"
	"reference-implementation/schemod/service/typesystem/id"
	"testing"

	"github.com/stretchr/testify/require"
)

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

func TestWrongArgument(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__B_bo_o1: query.Argument{
						Type:  typesystem.AR__ID_A,
						Value: "asd",
					},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

func TestWrongArgumentType(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: query.Argument{
						Type:  typesystem.PT__Bool,
						Value: true,
					},
				},
			},
		},
	}
	require.Error(t, q.Validate())
}

func TestPolymorphicParameter(t *testing.T) {
	assignableTypes := map[typesystem.IDType]interface{}{
		typesystem.AS__root_rp_page_struct1: true,
		typesystem.AR__ID_A:                 true,
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
