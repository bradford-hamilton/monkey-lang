package object

import "fmt"

// Boolean type holds the value of the boolean as a bool
type Boolean struct {
	Value bool
}

// Type returns our Boolean's ObjectType (BooleanObj)
func (b *Boolean) Type() ObjectType { return BooleanObj }

// Inspect returns a string representation of the Boolean's Value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
