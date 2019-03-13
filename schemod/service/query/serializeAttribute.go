package query

import (
	"encoding/binary"
	"io"
	"reference-implementation/schemod/service/typesystem"
)

func serializeAttribute(
	nodeIndexCounter *uint32,
	propNodeIdx uint32,
	byteOrder binary.ByteOrder,
	stream io.Writer,
	attrID typesystem.IDAttribute,
	attr AttributeSelection,
) error {
	//TODO: implement
	return nil
}
