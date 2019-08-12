package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/ast"
)

// ObjectType is merely a string. We define different object types as strings
// and can then reference them as such rather than "string"
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
	BuiltinObj     = "BUILTIN"
	ArrayObj       = "ARRAY"
	HashObj        = "HASH"
)

// Object represents monkey's object system. Every value in monkey-lang
// must implement this interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Hashable is one method called HashKey. Any object that that can be used as a HashKey
// must implement this interface (*object.String, *object.boolean, *object.integer)
type Hashable interface {
	HashKey() HashKey
}

// BuiltinFunction is a type representing functions we write in Go and
// expose to our users inside monkey-lang
type BuiltinFunction func(args ...Object) Object

// Builtin is our object wrapper holding a builtin function
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns our Builtin's ObjectType
func (n *Builtin) Type() ObjectType { return BuiltinObj }

// Inspect simply returns "builtin function"
func (n *Builtin) Inspect() string { return "builtin function" }

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

// Array type wraps the array's elements in a slice of Objects
type Array struct {
	Elements []Object
}

// Type returns our Array's ObjectType (ArrayObj)
func (a *Array) Type() ObjectType { return ArrayObj }

// Inspect returns a string representation of the Array's elements: [1, 2, 3]
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey type wraps the key's type and holds its value
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey returns a HashKey with a Value of 1 or 0 (true or false) and a Type of BooleanObj
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{
		Type:  b.Type(),
		Value: value,
	}
}

// HashKey returns a HashKey with a Value of the Integer and a Type of IntegerObj
func (i *Integer) HashKey() HashKey {
	return HashKey{
		Type:  i.Type(),
		Value: uint64(i.Value),
	}
}

// HashKey returns a HashKey with a Value of a 64-bit FNV-1a hash of the String and a Type of StringObj
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{
		Type:  s.Type(),
		Value: h.Sum64(),
	}
}

// HashPair holds key value pairs
type HashPair struct {
	Key   Object
	Value Object
}

// Hash hold Pairs which are a map of HashKey -> HashPair. We map the keys to HashPairs instead
// of Objects for a better ability to Inspect() and see the key and value
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type returns Hash's ObjectType (HashObj)
func (h *Hash) Type() ObjectType { return HashObj }

// Inspect returns a string representation of the Hash
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(
			pairs,
			fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()),
		)
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

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
