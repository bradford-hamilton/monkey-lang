package object

// Null type is an empty struct
type Null struct{}

// Type returns Null's ObjectType (NullObj)
func (n *Null) Type() ObjectType { return NullObj }

// Inspect returns a string representation of Null ("null")
func (n *Null) Inspect() string { return "null" }
