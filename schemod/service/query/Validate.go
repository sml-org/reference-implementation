package query

import (
	"fmt"
	"reference-implementation/schemod/service/typesystem"
	"reflect"
)

// ValidateAttribute recursively validates an attribute selection
func ValidateAttribute(
	attrID typesystem.IDAttribute,
	attrSelection AttributeSelection,
) error {
	attrTypeID := attrID.Type()

	if !attrTypeID.HasQueryableProperties() {
		// Ensure no selections are made on attributes of non-composite type
		if len(attrSelection.Props) > 0 {
			return fmt.Errorf(
				"selection on non-composite leaf attribute of type %s",
				attrTypeID.String(),
			)
		}

		// Scalar attributes don't need to be validated
		return nil
	}

	// Verify attribute properties
	for propID, selection := range attrSelection.Props {
		if err := ValidateProperty(
			attrTypeID,
			propID,
			selection,
		); err != nil {
			return err
		}
	}

	return nil
}

// ValidateProperty recursively validates a property query node
func ValidateProperty(
	expectedParentTypeID typesystem.IDType,
	propID typesystem.IDProperty,
	selection PropertySelection,
) error {
	// Ensure the property is expected in the given node type
	propParentTypeID := propID.HostType()
	if propParentTypeID != expectedParentTypeID {
		return fmt.Errorf(
			"property %s is a property of type %s and not of type %s",
			propID.String(),
			propParentTypeID.String(),
			expectedParentTypeID.String(),
		)
	}

	// Verify arguments
	for argID, arg := range selection.Args {
		argPropertyID := argID.Property()
		// Ensure the argument is applicable to this property
		if argPropertyID != propID {
			return fmt.Errorf(
				"argument %s is applicable to property %s "+
					"and not applicable to property %s",
				argID.String(),
				argPropertyID.String(),
				propID.String(),
			)
		}

		// Ensure the provided argument type is assignable
		paramTypeID := argID.Type()
		if !paramTypeID.IsAssignable(arg.Type) {
			return fmt.Errorf(
				"mismatching types: %s can't be used as a value for "+
					"argument %s of type %s",
				arg.Type.String(),
				argID.String(),
				paramTypeID.String(),
			)
		}

		// Validate argument value
		if !paramTypeID.IsValueAssignable(arg.Value) {
			return fmt.Errorf(
				"mismatching types: can't assign value of Go type %s "+
					"to argument of type %s",
				reflect.TypeOf(arg.Value).String(),
				arg.Type.String(),
			)
		}
	}

	// Verify attributes
	for attrID, attrSelection := range selection.Attrs {
		attrPropertyID := attrID.Property()
		// Ensure the attribute belongs to this property
		if attrPropertyID != propID {
			return fmt.Errorf(
				"attribute %s is an attribute of property %s "+
					"and not an attribute of %s",
				attrID.String(),
				attrPropertyID.String(),
				propID.String(),
			)
		}

		// Recursively verify the attribute selection
		if err := ValidateAttribute(attrID, attrSelection); err != nil {
			return err
		}
	}

	if !propParentTypeID.HasQueryableProperties() {
		// Ensure no selections are made on properties of non-composite type
		if len(selection.Props) > 0 {
			return fmt.Errorf(
				"selection on non-composite leaf property of type %s",
				propParentTypeID.String(),
			)
		}

		// Scalar properties don't need to be verified further
		return nil
	}

	// Verify sub-properties
	for subPropID, subPropSelection := range selection.Props {
		if err := ValidateProperty(
			propParentTypeID,
			subPropID,
			subPropSelection,
		); err != nil {
			return err
		}
	}

	return nil
}

// Validate validates the query object and returns an error if there's a
// violation, otherwise returns false
func (q Query) Validate() error {
	for propID, prop := range q.Props {
		if err := ValidateProperty(
			typesystem.PT__Root,
			propID,
			prop,
		); err != nil {
			return err
		}
	}
	return nil
}
