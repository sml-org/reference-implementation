package metainfo

const SchemaVersionID = "66d365641d4de50e0d3a91323b93108e"

var SchemaVersionIDBytes = []byte{
	0x66, 0xd3, 0x65, 0x64,
	0x1d, 0x4d, 0xe5, 0x0e,
	0x0d, 0x3a, 0x91, 0x32,
	0x3b, 0x93, 0x10, 0x8e,
}

const ProtocolVersionMajor = 1
const ProtocolVersionMinor = 0

// LenNodeIndex defines the graph node index length in bytes
const LenNodeIndex = 4

// LenIDType defines the type identifier length in bytes
const LenIDType = 4

// LenNodeIDType defines the graph node identifier length in bytes
const LenNodeIDType = 4
