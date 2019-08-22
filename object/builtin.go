package object

import (
	"fmt"
)

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

// GetBuiltinByName takes a name, iterates over our builtins slice and returns
// the appropriate builtin
func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}

func newError(msgWithFormatVerbs string, values ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(msgWithFormatVerbs, values...)}
}

// Builtins defines all of Monkey's built in functions
var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{
		"len",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
				}
				switch arg := args[0].(type) {
				case *Array:
					return &Integer{Value: int64(len(arg.Elements))}
				case *String:
					return &Integer{Value: int64(len(arg.Value))}
				default:
					return newError("Argument to `len` not supported. Got: %s", args[0].Type())
				}
			},
		},
	},
	{
		"puts",
		&Builtin{
			Fn: func(args ...Object) Object {
				for _, arg := range args {
					fmt.Println(arg.Inspect())
				}
				return nil
			},
		},
	},
	{
		"first",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
				}
				if args[0].Type() != ArrayObj {
					return newError("Argument to `first` must be an Array. Got: %s", args[0].Type())
				}

				array := args[0].(*Array)
				if len(array.Elements) > 0 {
					return array.Elements[0]
				}

				return nil
			},
		},
	},
	{
		"last",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
				}
				if args[0].Type() != ArrayObj {
					return newError("Argument to `last` must be an Array. Got: %s", args[0].Type())
				}

				array := args[0].(*Array)
				length := len(array.Elements)
				if length > 0 {
					return array.Elements[length-1]
				}

				return nil
			},
		},
	},
	{
		"rest",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 1 {
					return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
				}
				if args[0].Type() != ArrayObj {
					return newError("Argument to `rest` must be an Array. Got: %s", args[0].Type())
				}

				array := args[0].(*Array)
				length := len(array.Elements)
				if length > 0 {
					newElements := make([]Object, length-1, length-1)
					copy(newElements, array.Elements[1:length])
					return &Array{Elements: newElements}
				}

				return nil
			},
		},
	},
	{
		"push",
		&Builtin{
			Fn: func(args ...Object) Object {
				if len(args) != 2 {
					return newError("Wrong number of arguments. Got: %d, Expected: 2", len(args))
				}
				if args[0].Type() != ArrayObj {
					return newError("Argument to `push` must be an Array. Got: %s", args[0].Type())
				}

				array := args[0].(*Array)
				length := len(array.Elements)

				newElements := make([]Object, length+1, length+1)
				copy(newElements, array.Elements)
				newElements[length] = args[1]

				return &Array{Elements: newElements}
			},
		},
	},
}
