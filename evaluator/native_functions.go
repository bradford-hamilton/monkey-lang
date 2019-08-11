package evaluator

import "github.com/bradford-hamilton/monkey-lang/object"

var nativeFunctions = map[string]*object.Native{
	"len": &object.Native{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Argument to `len` not supported. Got: %s", args[0].Type())
			}
		},
	},
}
