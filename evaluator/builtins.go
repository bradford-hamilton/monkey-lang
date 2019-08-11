package evaluator

import "github.com/bradford-hamilton/monkey-lang/object"

var builtinFunctions = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Argument to `len` not supported. Got: %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("Argument to `first` must be an Array. Got: %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			if len(array.Elements) > 0 {
				return array.Elements[0]
			}

			return Null
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("Argument to `last` must be an Array. Got: %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				return array.Elements[length-1]
			}

			return Null
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("Argument to `rest` must be an Array. Got: %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, array.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return Null
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Wrong number of arguments. Got: %d, Expected: 2", len(args))
			}
			if args[0].Type() != object.ArrayObj {
				return newError("Argument to `push` must be an Array. Got: %s", args[0].Type())
			}

			array := args[0].(*object.Array)
			length := len(array.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, array.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
}
