package object

// String type holds the value of the string
type String struct {
	Value string
}

// Type returns our String's ObjectType
func (s *String) Type() ObjectType { return StringObj }

// Inspect returns a string representation of the String's Value
func (s *String) Inspect() string { return s.Value }
