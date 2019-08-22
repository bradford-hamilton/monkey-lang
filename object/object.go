package object

// ObjectType - Type alias for a string.
type ObjectType string

// Define object types
const (
	IntegerObj          = "INTEGER"
	BooleanObj          = "BOOLEAN"
	NullObj             = "NULL"
	ReturnValueObj      = "RETURN_VALUE"
	ErrorObj            = "ERROR"
	FunctionObj         = "FUNCTION"
	StringObj           = "STRING"
	BuiltinObj          = "BUILTIN"
	ArrayObj            = "ARRAY"
	HashObj             = "HASH"
	CompiledFunctionObj = "COMPILED_FUNCTION_OBJ"
	ClosureObj          = "CLOSURE"
)

// Object represents monkey's object system. Every value in monkey-lang
// must implement this interface
type Object interface {
	Type() ObjectType
	Inspect() string
}
