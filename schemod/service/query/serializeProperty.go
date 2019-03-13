package query

import (
	"encoding/binary"
	"io"
	"reference-implementation/schemod/service/metainfo"
	"reference-implementation/schemod/service/typesystem"
)

func serializeProperty(
	nodeIndexCounter *uint32,
	parentIndex uint32,
	byteOrder binary.ByteOrder,
	stream io.Writer,
	propID typesystem.IDProperty,
	propSelection PropertySelection,
) error {
	(*nodeIndexCounter)++
	propNodeIdx := *nodeIndexCounter

	buf := make(
		[]byte,
		metainfo.LenNodeIndex+ // Graph Node Index
			metainfo.LenNodeIndex+ // Graph Node Parent Index
			metainfo.LenNodeIDType, // Graph Node ID
	)

	byteOrder.PutUint32(buf[0:4], uint32(propID))
	byteOrder.PutUint32(buf[4:8], propNodeIdx)
	byteOrder.PutUint32(buf[8:12], parentIndex)

	// Write graph node type
	if _, err := stream.Write(buf); err != nil {
		return err
	}

	// Write parameters
	for paramID, arg := range propSelection.Args {
		if err := serializeArg(
			nodeIndexCounter,
			propNodeIdx,
			byteOrder,
			stream,
			paramID,
			arg,
		); err != nil {
			return err
		}
	}

	// Write attributes
	for attrID, attr := range propSelection.Attrs {
		if err := serializeAttribute(
			nodeIndexCounter,
			propNodeIdx,
			byteOrder,
			stream,
			attrID,
			attr,
		); err != nil {
			return err
		}
	}

	return nil
}
