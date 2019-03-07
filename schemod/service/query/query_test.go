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
	require.NoError(t, q.Validate())
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
	require.NoError(t, q.Validate())
}

func TestWrongAcceptedUnionType(t *testing.T) {
	q := query.Query{
		Props: query.Props{
			typesystem.PR__root_rp: query.PropertySelection{
				Args: query.Args{
					typesystem.PM__root_rp_page: query.Argument{
						Type:  typesystem.AR__ID_A,
						Value: true,
					},
				},
			},
		},
	}
	require.NoError(t, q.Validate())
}
