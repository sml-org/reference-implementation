package service

// SchemaSubjectCategory reflects the category of schema subject
type SchemaSubjectCategory uint32

const (
	_ SchemaSubjectCategory = iota

	// Type represents the type category
	Type

	// Property represents the property category
	Property

	// PropertyParameter represents the property parameter category
	PropertyParameter

	// EnumValue represents the enum value category
	EnumValue
)
