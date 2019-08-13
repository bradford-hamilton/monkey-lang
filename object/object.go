package object

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
