package typesystem

// R_Integer implements the type: trait integer
type R_Integer interface {
	FromInt32(T_Int32)
	FromUint32(T_Uint32)
	FromInt64(T_Int64)
	FromUint64(T_Uint64)
	IDType() IDType
}
