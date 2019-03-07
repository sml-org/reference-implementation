package engine

import (
	"reference-implementation/schemod/service/typesystem/entity"
	"reference-implementation/schemod/service/typesystem/id"
	"reference-implementation/schemod/service/typesystem/prim"
)

// FetchA2 fetches A by predicate:
//   $t:id > cursor
func (ng Engine) FetchA2(
	cursor id.A,
	limit prim.Int32,
) []entity.A {
	/* TODO: fetch A by: $t:id > cursor */
	return nil
}
