package object

import (
	"fmt"
	"strings"
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
	{"len", &Builtin{Fn: bLen}},
	{"print", &Builtin{Fn: bPrint}},
	{"first", &Builtin{Fn: bFirst}},
	{"last", &Builtin{Fn: bLast}},
	{"rest", &Builtin{Fn: bRest}},
	{"push", &Builtin{Fn: bPush}},
	{"pop", &Builtin{Fn: bPop}},
	{"split", &Builtin{Fn: bSplit}},
	{"join", &Builtin{Fn: bJoin}},
}

func bLen(args ...Object) Object {
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
}

func bPrint(args ...Object) Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}
	return nil
}

func bFirst(args ...Object) Object {
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
}

func bLast(args ...Object) Object {
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
}

func bRest(args ...Object) Object {
	if len(args) != 1 {
		return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
	}
	if args[0].Type() != ArrayObj {
		return newError("Argument to `rest` must be an Array. Got: %s", args[0].Type())
	}

	array := args[0].(*Array)
	length := len(array.Elements)
	if length > 0 {
		newElements := make([]Object, length-1)
		copy(newElements, array.Elements[1:length])
		return &Array{Elements: newElements}
	}

	return nil
}

func bPush(args ...Object) Object {
	if len(args) != 2 {
		return newError("Wrong number of arguments. Got: %d, Expected: 2", len(args))
	}
	if args[0].Type() != ArrayObj {
		return newError("Argument to `push` must be an Array. Got: %s", args[0].Type())
	}

	array := args[0].(*Array)
	length := len(array.Elements)

	newElements := make([]Object, length+1)
	copy(newElements, array.Elements)
	newElements[length] = args[1]

	return &Array{Elements: newElements}
}

func bPop(args ...Object) Object {
	if len(args) != 1 {
		return newError("Wrong number of arguments. Got: %d, Expected: 1", len(args))
	}
	if args[0].Type() != ArrayObj {
		return newError("Argument to `pop` must be an Array. Got: %s", args[0].Type())
	}

	array := args[0].(*Array)
	length := len(array.Elements)
	if length == 0 {
		return nil
	}

	newElements := make([]Object, length-1)
	copy(newElements, array.Elements[0:length-1])

	return &Array{Elements: newElements}
}

func bSplit(args ...Object) Object {
	if len(args) != 2 {
		return newError("Wrong number of arguments. Got: %d, Expected: 2", len(args))
	}
	if args[0].Type() != StringObj {
		return newError("First argument to `split` must be a String. Got: %s", args[0].Type())
	}
	if args[1].Type() != StringObj {
		return newError("Second argument to `split` must be a String. Got: %s", args[1].Type())
	}

	str := args[0].(*String)
	splitOn := args[1].(*String)
	array := strings.Split(str.Value, splitOn.Value)

	monkeyArray := []Object{}
	for _, v := range array {
		monkeyArray = append(monkeyArray, &String{Value: v})
	}

	return &Array{Elements: monkeyArray}
}

func bJoin(args ...Object) Object {
	if len(args) != 2 {
		return newError("Wrong number of arguments. Got: %d, Expected: 2", len(args))
	}
	if args[0].Type() != ArrayObj {
		return newError("First argument to `join` must be an Array. Got: %s", args[0].Type())
	}
	if args[1].Type() != StringObj {
		return newError("Second argument to `join` must be a String. Got: %s", args[1].Type())
	}

	array := args[0].(*Array)
	length := len(array.Elements)
	if length == 0 {
		return &String{Value: ""}
	}
	joinOn := args[1].(*String)
	goString := joinOn.Value

	var goSlice []string
	for _, e := range array.Elements {
		switch obj := e.(type) {
		case *String:
			goSlice = append(goSlice, obj.Value)
		default:
			return newError("You can only join an array of all strings. Illegal type found: %s", obj.Type())
		}
	}

	return &String{
		Value: strings.Join(goSlice, goString),
	}
}
