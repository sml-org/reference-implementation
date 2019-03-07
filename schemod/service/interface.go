package service

import (
	"reference-implementation/schemod/service/typesystem/entity"
	"reference-implementation/schemod/service/typesystem/id"
	"reference-implementation/schemod/service/typesystem/prim"
	tp "reference-implementation/schemod/service/typesystem/type"
)

// Loader defines the loader interface
type Loader interface {
	// FetchA1 fetches A by predicate:
	//   $t:id in p
	FetchA1(
		p []id.A,
		limit prim.Int32,
	) []entity.A

	// FetchA2 fetches A by predicate:
	//   $t:id > cursor
	FetchA2(
		cursor id.A,
		limit prim.Int32,
	) []entity.A

	// FetchA3 fetches A by predicate:
	//   $t:id < cursor
	FetchA3(
		cursor id.A,
		limit prim.Int32,
	) []entity.A
}

// Engine defines the interface of the service engine
type Engine interface {
	Loader
	Initialize(
		ra entity.A,
		rb tp.B,
		re tp.E,
		rc []tp.C, // optional
	) error
}
