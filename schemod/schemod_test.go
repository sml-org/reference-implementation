package schemod_test

import (
	"reference-implementation/schemod/engine"
	"reference-implementation/schemod/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type str = []byte

func TestQL(t *testing.T) {
	/*
		query($cursor ID<A>, $limit Int32) {
			rp(page: {
				cursor: $cursor
				limit:  $limit
			}) {

			}
		}
		args: {

		}
	*/
	var transport interface{}
	server := service.NewServer(engine.Engine{})
	go func() {
		assert.NoError(t, server.Serve(transport))
	}()

	raw, err := server.Query(Query)
	require.NoError(t, err)
}
