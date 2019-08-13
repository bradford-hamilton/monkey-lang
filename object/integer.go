package object

import "fmt"

// Integer type holds the value of the integer as an int64
type Integer struct {
	Value int64
}

// Type returns our Integer's ObjectType
func (i *Integer) Type() ObjectType { return IntegerObj }

// Inspect returns a string representation of the Integer's Value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
