package query

import (
	"reference-implementation/schemod/service/metainfo"
)

var queryHeader []byte

func init() {
	queryHeader = make([]byte, 16)
	copy(queryHeader[0:16], metainfo.SchemaVersionIDBytes)
}
