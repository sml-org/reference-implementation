package engine

import (
	"reference-implementation/schemod/service/typesystem/entity"
	"reference-implementation/schemod/service/typesystem/id"
	"reference-implementation/schemod/service/typesystem/prim"
)

// FetchA1 fetches A by predicate:
//   $t:id in p
func (ng Engine) FetchA1(
	p []id.A,
	limit prim.Int32,
) []entity.A {
	/* TODO: fetch A by: $t:id in p */
	return nil
}
