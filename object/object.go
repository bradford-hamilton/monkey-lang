package object

import "fmt"

type ObjectType string

// Define object types
const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	NullObj    = "NULL"
)

// Object represents monkey's object system. Every value in monkey-lang
// will be wrapped inside this struct
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer type holds the value of the integer as an int64
type Integer struct {
	Value int64
}

// Type returns our Integer's ObjectType
func (i *Integer) Type() ObjectType { return IntegerObj }

// Inspect returns a string representation of the Integer's Value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Boolean type holds the value of the boolean as a bool
type Boolean struct {
	Value bool
}

// Type returns our Boolean's ObjectType
func (b *Boolean) Type() ObjectType { return BooleanObj }

// Inspect returns a string representation of the Boolean's Value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null type is an empty struct
type Null struct{}

// Type returns Null's ObjectType
func (n *Null) Type() ObjectType { return NullObj }

// Inspect returns a string representation of Null ("null")
func (n *Null) Inspect() string { return "null" }
