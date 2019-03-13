package query

import (
	"encoding/binary"
	"fmt"
	"io"
	"reference-implementation/schemod/service/metainfo"
	"reference-implementation/schemod/service/typesystem"
)

func serializeArg(
	nodeIndexCounter *uint32,
	parentIndex uint32,
	byteOrder binary.ByteOrder,
	stream io.Writer,
	paramID typesystem.IDParameter,
	arg typesystem.Serializable,
) error {
	(*nodeIndexCounter)++
	argNodeIdx := *nodeIndexCounter

	// Write argument header
	buf := make(
		[]byte,
		metainfo.LenNodeIndex+ // Graph Node Index
			metainfo.LenNodeIndex+ // Graph Node Parent Index
			metainfo.LenNodeIDType+ // Graph Node ID
			metainfo.LenIDType, // Argument Data Type
	)

	byteOrder.PutUint32(buf[0:4], uint32(paramID))
	byteOrder.PutUint32(buf[4:8], argNodeIdx)
	byteOrder.PutUint32(buf[8:12], parentIndex)
	byteOrder.PutUint32(buf[12:16], uint32(arg.Type()))

	if _, err := stream.Write(buf); err != nil {
		return err
	}

	// Write argument data
	argDataLen := arg.Len()
	if argDataLen > 4294967295 {
		// Overflow
		return fmt.Errorf("argument data too long (%d)", argDataLen)
	}
	if argDataLen < 1 {
		return fmt.Errorf("missing argument data")
	}

	return arg.Serialize(byteOrder, stream)
}
