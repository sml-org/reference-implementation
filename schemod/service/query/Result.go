package query

import (
	"reference-implementation/schemod/service/typesystem"
)

// ResultValue represents a result value
type ResultValue struct {
	ValueType typesystem.IDType
	Value     interface{}
}

// Result represents a query result object
type Result struct {
	Data   []ResultValue
	Errors []error
}
