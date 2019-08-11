package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/ast"
)

type ObjectType string

// Define object types
const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	StringObj      = "STRING"
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

// String type holds the value of the string
type String struct {
	Value string
}

// Type returns our String's ObjectType
func (s *String) Type() ObjectType { return StringObj }

// Inspect returns a string representation of the String's Value
func (s *String) Inspect() string { return s.Value }

// Boolean type holds the value of the boolean as a bool
type Boolean struct {
	Value bool
}

// Type returns our Boolean's ObjectType (BooleanObj)
func (b *Boolean) Type() ObjectType { return BooleanObj }

// Inspect returns a string representation of the Boolean's Value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null type is an empty struct
type Null struct{}

// Type returns Null's ObjectType (NullObj)
func (n *Null) Type() ObjectType { return NullObj }

// Inspect returns a string representation of Null ("null")
func (n *Null) Inspect() string { return "null" }

// ReturnValue type holds a return value
type ReturnValue struct {
	Value Object
}

// Type returns our ReturnValue's ObjectType (ReturnValueObj)
func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }

// Inspect returns a string representation of the ReturnValue's Value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error type holds an error Message
type Error struct {
	Message string
}

// Type returns our Error's ObjectType (ErrorObj)
func (e *Error) Type() ObjectType { return ErrorObj }

// Inspect returns an error message string
func (e *Error) Inspect() string { return "Error: " + e.Message }

// NewEnvironment creates and returns a pointer to an Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new Environment and attaches the outer environment
// that's passed in, to the new environment, as it's enclosing environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Environment holds a store of key value pairs and a pointer to an "outer", enclosing environment
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get retrieves a key from an Environment's store by name. If it does not find it, it recursively looks
// for the key in the enclosing environment(s)
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets a key to an Environment's store by name
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// Function holds Parameters as a slice of *Identifier, a Body which is a *ast.BlockStatement
// and a pointer to it's environment
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns our Function's ObjectType (FunctionObj)
func (f *Function) Type() ObjectType { return FunctionObj }

// Inspect returns a string representation of the function definition
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}
