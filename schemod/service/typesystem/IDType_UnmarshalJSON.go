package typesystem

// UnmarshalJSON unmarshals the value from a JSON meta-object
func (o IDType) UnmarshalJSON(b []byte) error {
	return o.FromString(string(b))
}
