package object

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
