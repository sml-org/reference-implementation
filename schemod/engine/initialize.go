package engine

import (
	"reference-implementation/schemod/service/typesystem/entity"
	tp "reference-implementation/schemod/service/typesystem/type"
)

// Initialize
func (ng Engine) Initialize(
	ra entity.A,
	rb tp.B,
	re tp.E,
	rc []tp.C, // optional
) error {
	/* TODO: initialize database */
	return nil
}
