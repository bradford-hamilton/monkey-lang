package object

import (
	"fmt"
	"testing"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/token"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	int1 := &Integer{Value: 10}
	int2 := &Integer{Value: 10}
	diff1 := &Integer{Value: 30}
	diff2 := &Integer{Value: 30}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}

	if int1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}

func TestHash(t *testing.T) {
	h := &Hash{
		Pairs: map[HashKey]HashPair{
			HashKey{
				Type:  StringObj,
				Value: 1,
			}: HashPair{
				Key:   &String{Value: "monkey"},
				Value: &String{Value: "lang"},
			},
		},
	}

	if h.Type() != HashObj {
		t.Errorf("integer.Type() returned wrong type. Expected: HashObj. Got: %s", h.Type())
	}

	if h.Inspect() != "{monkey: lang}" {
		t.Errorf("h.Inspect() returned wrong string representation. Expected: {monkey: lang}. Got: %s", h.Inspect())
	}
}

func TestArray(t *testing.T) {
	elements := []Object{
		&Integer{Value: 1},
		&Integer{Value: 2},
		&Integer{Value: 3},
	}
	arr := &Array{Elements: elements}

	if arr.Type() != ArrayObj {
		t.Errorf("arr.Type() returned wrong type. Expected: ArrayObj. Got: %s", arr.Type())
	}

	if arr.Inspect() != "[1, 2, 3]" {
		t.Errorf("arr.Inspect() returned wrong string representation. Expected: [1, 2, 3]. Got: %s", arr.Inspect())
	}
}

func TestBoolean(t *testing.T) {
	b := &Boolean{}

	if b.Type() != BooleanObj {
		t.Errorf("b.Type() returned wrong type. Expected: BooleanObj. Got: %s", b.Type())
	}

	if b.Inspect() != "false" {
		t.Errorf("arr.Inspect() returned wrong string representation. Expected: false. Got: %s", b.Inspect())
	}
}

func TestClosure(t *testing.T) {
	cl := &Closure{
		Fn:   &CompiledFunction{},
		Free: []Object{},
	}

	if cl.Type() != ClosureObj {
		t.Errorf("cl.Type() returned wrong type. Expected: ClosureObj. Got: %s", cl.Type())
	}

	if cl.Inspect() != fmt.Sprintf("Closure[%p]", cl) {
		t.Errorf("cl.Inspect() returned wrong string representation. Expected: Closure[%p]. Got: %s", cl, cl.Inspect())
	}
}

func TestCompiledFunction(t *testing.T) {
	cf := &CompiledFunction{
		Instructions:  []byte("OpDoesntMatter"),
		NumLocals:     1,
		NumParameters: 1,
	}

	if cf.Type() != CompiledFunctionObj {
		t.Errorf("cf.Type() returned wrong type. Expected: CompiledFunctionObj. Got: %s", cf.Type())
	}

	if cf.Inspect() != fmt.Sprintf("CompiledFunction[%p]", cf) {
		t.Errorf("cf.Inspect() returned wrong string representation. Expected: CompiledFunction[%p]. Got: %s", cf, cf.Inspect())
	}
}

func TestErrors(t *testing.T) {
	e := &Error{
		Message: "Uh oh spaghettio",
	}

	if e.Type() != ErrorObj {
		t.Errorf("e.Type() returned wrong type. Expected: ErrorObj. Got: %s", e.Type())
	}

	if e.Inspect() != "Error: Uh oh spaghettio" {
		t.Errorf("e.Inspect() returned wrong string representation. Expected: Error: Uh oh spaghettio. Got: %s", e.Inspect())
	}
}

func TestFunctions(t *testing.T) {
	f := &Function{
		Parameters: []*ast.Identifier{
			&ast.Identifier{
				Token: token.Token{
					Type: token.String,
				},
				Value: "arg1",
			},
		},
		Body: &ast.BlockStatement{
			Token: token.Token{Type: token.String, Literal: "let"},
			Statements: []ast.Statement{
				&ast.LetStatement{
					Token: token.Token{Type: token.String, Literal: "let"},
					Name:  &ast.Identifier{Value: "waaat"},
					Value: &ast.StringLiteral{Token: token.Token{Literal: "thing"}},
				},
			},
		},
		Env: &Environment{},
	}

	if f.Type() != FunctionObj {
		t.Errorf("f.Type() returned wrong type. Expected: FunctionObj. Got: %s", f.Type())
	}

	if f.Inspect() != "func(arg1) {\nlet waaat = thing;\n" {
		t.Errorf("f.Inspect() returned wrong string representation. Expected:\n func(arg1) {\nlet waaat = thing;\n. Got:\n %s", f.Inspect())
	}
}

func TestIntegers(t *testing.T) {
	integer := &Integer{Value: 666}

	if integer.Type() != IntegerObj {
		t.Errorf("integer.Type() returned wrong type. Expected: IntegerObj. Got: %s", integer.Type())
	}

	if integer.Inspect() != "666" {
		t.Errorf("integer.Inspect() returned wrong string representation. Expected: 666. Got: %s", integer.Inspect())
	}
}

func TestNull(t *testing.T) {
	n := &Null{}

	if n.Type() != NullObj {
		t.Errorf("n.Type() returned wrong type. Expected: NullObj. Got: %s", n.Type())
	}

	if n.Inspect() != "null" {
		t.Errorf("n.Inspect() returned wrong string representation. Expected: null. Got: %s", n.Inspect())
	}
}

func TestReturnValues(t *testing.T) {
	rv := &ReturnValue{Value: &String{Value: "im a returned string"}}

	if rv.Type() != ReturnValueObj {
		t.Errorf("n.Type() returned wrong type. Expected: ReturnValueObj. Got: %s", rv.Type())
	}

	if rv.Inspect() != "im a returned string" {
		t.Errorf("n.Inspect() returned wrong string representation. Expected: im a returned string. Got: %s", rv.Inspect())
	}
}

func TestStrings(t *testing.T) {
	s := &String{Value: "thurman merman"}

	if s.Type() != StringObj {
		t.Errorf("n.Type() returned wrong type. Expected: StringObj. Got: %s", s.Type())
	}

	if s.Inspect() != "thurman merman" {
		t.Errorf("n.Inspect() returned wrong string representation. Expected: thurman merman. Got: %s", s.Inspect())
	}
}

func TestEnvironments(t *testing.T) {
	env := NewEnclosedEnvironment(NewEnvironment())

	env.Set("innerKey", &String{Value: "innerValue"})
	env.outer.Set("outerKey", &String{Value: "outerValue"})

	obj, ok := env.Get("innerKey")
	if !ok {
		t.Errorf("Failure to retrieve key in env")
	}
	if obj.Inspect() != "innerValue" {
		t.Errorf("Expected 'innerValue'. Got: %s", obj.Inspect())
	}

	obj, ok = env.Get("outerKey")
	if !ok {
		t.Errorf("Failure to retrieve key in env")
	}
	if obj.Inspect() != "outerValue" {
		t.Errorf("Expected 'outerValue'. Got: %s", obj.Inspect())
	}
}

func TestBuiltins(t *testing.T) {
	b := &Builtin{}

	if b.Type() != BuiltinObj {
		t.Errorf("b.Type() returned wrong type. Expected: BuiltinObj. Got: %s", b.Type())
	}

	if b.Inspect() != "builtin function" {
		t.Errorf("b.Inspect() returned wrong string representation. Expected: builtin function. Got: %s", b.Inspect())
	}

	notABuiltin := GetBuiltinByName("notABuiltin")
	if notABuiltin != nil {
		t.Errorf("GetBuiltinByName(\"notABuiltin\") should have return nil")
	}

	err := newError("Message with %s %s", "format", "verbs")
	if err.Message != "Message with format verbs" {
		t.Errorf("newError returned wrong error string. Expected: 'Message with format verbs'. Got: %s", err.Message)
	}
}

func TestLen(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 1},
		&Integer{Value: 2},
		&Integer{Value: 3},
	}}
	str := &String{Value: "neat string"}
	null := &Null{}

	lenBuiltin := GetBuiltinByName("len")
	if lenBuiltin.Fn(str, null).Inspect() != "Error: Wrong number of arguments. Got: 2, Expected: 1" {
		t.Errorf("len builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", lenBuiltin.Fn(null).Inspect())
	}
	if lenBuiltin.Fn(arr).Inspect() != "3" {
		t.Errorf("len builtin returned wrong result. Expected: 3. Got: %s", lenBuiltin.Fn(arr).Inspect())
	}
	if lenBuiltin.Fn(str).Inspect() != "11" {
		t.Errorf("len builtin returned wrong result. Expected: 11. Got: %s", lenBuiltin.Fn(str).Inspect())
	}
	if lenBuiltin.Fn(null).Inspect() != "Error: Argument to `len` not supported. Got: NULL" {
		t.Errorf("len builtin returned wrong result. Expected: Error: Argument to `len` not supported. Got: NULL. Got: %s", lenBuiltin.Fn(null).Inspect())
	}
}

func TestPrint(t *testing.T) {
	str := &String{Value: "neat string"}

	printBuiltin := GetBuiltinByName("print")
	if printBuiltin.Fn(str) != nil {
		t.Errorf("print builtin should print its arguments and return nil. Returned: %s", printBuiltin.Fn(str))
	}
}

func TestFirst(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 99},
		&Integer{Value: 7},
		&Integer{Value: 356},
	}}
	emptyArr := &Array{}
	str := &String{Value: "neat string"}

	firstBuiltin := GetBuiltinByName("first")
	if firstBuiltin.Fn(arr, arr).Inspect() != "Error: Wrong number of arguments. Got: 2, Expected: 1" {
		t.Errorf("first builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", firstBuiltin.Fn(str).Inspect())
	}
	if firstBuiltin.Fn(str).Inspect() != "Error: Argument to `first` must be an Array. Got: STRING" {
		t.Errorf("first builtin returned wrong result. Expected: Error: Argument to `first` must be an Array. Got: STRING. Got: %s", firstBuiltin.Fn(str).Inspect())
	}
	if firstBuiltin.Fn(arr).Inspect() != "99" {
		t.Errorf("first builtin returned wrong result. Expected: 99. Got: %s", firstBuiltin.Fn(arr).Inspect())
	}
	if firstBuiltin.Fn(emptyArr) != nil {
		t.Errorf("first builtin returned wrong result. Expected: nil. Got: %s", firstBuiltin.Fn(arr).Inspect())
	}
}

func TestLast(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 99},
		&Integer{Value: 7},
		&Integer{Value: 356},
	}}
	emptyArr := &Array{}
	str := &String{Value: "neat string"}

	lastBuiltin := GetBuiltinByName("last")
	if lastBuiltin.Fn(arr, arr).Inspect() != "Error: Wrong number of arguments. Got: 2, Expected: 1" {
		t.Errorf("last builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", lastBuiltin.Fn(str).Inspect())
	}
	if lastBuiltin.Fn(str).Inspect() != "Error: Argument to `last` must be an Array. Got: STRING" {
		t.Errorf("last builtin returned wrong result. Expected: Error: Argument to `last` must be an Array. Got: STRING. Got: %s", lastBuiltin.Fn(str).Inspect())
	}
	if lastBuiltin.Fn(arr).Inspect() != "356" {
		t.Errorf("last builtin returned wrong result. Expected: 356. Got: %s", lastBuiltin.Fn(arr).Inspect())
	}
	if lastBuiltin.Fn(emptyArr) != nil {
		t.Errorf("last builtin returned wrong result. Expected: nil. Got: %s", lastBuiltin.Fn(arr).Inspect())
	}
}

func TestRest(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 99},
		&Integer{Value: 7},
		&Integer{Value: 356},
	}}
	emptyArr := &Array{}
	str := &String{Value: "neat string"}

	restBuiltin := GetBuiltinByName("rest")
	if restBuiltin.Fn(arr, arr).Inspect() != "Error: Wrong number of arguments. Got: 2, Expected: 1" {
		t.Errorf("rest builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", restBuiltin.Fn(str).Inspect())
	}
	if restBuiltin.Fn(str).Inspect() != "Error: Argument to `rest` must be an Array. Got: STRING" {
		t.Errorf("rest builtin returned wrong result. Expected: Error: Argument to `rest` must be an Array. Got: STRING. Got: %s", restBuiltin.Fn(str).Inspect())
	}
	if restBuiltin.Fn(arr).Inspect() != "[7, 356]" {
		t.Errorf("rest builtin returned wrong result. Expected: [7, 356]. Got: %s", restBuiltin.Fn(arr).Inspect())
	}
	if restBuiltin.Fn(emptyArr) != nil {
		t.Errorf("rest builtin returned wrong result. Expected: nil. Got: %s", restBuiltin.Fn(arr).Inspect())
	}
}

func TestPush(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 99},
		&Integer{Value: 7},
		&Integer{Value: 356},
	}}
	newEl := &Integer{Value: 666}
	str := &String{Value: "neat string"}

	pushBuiltin := GetBuiltinByName("push")
	if pushBuiltin.Fn(arr).Inspect() != "Error: Wrong number of arguments. Got: 1, Expected: 2" {
		t.Errorf("push builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", pushBuiltin.Fn(str).Inspect())
	}
	if pushBuiltin.Fn(str, str).Inspect() != "Error: Argument to `push` must be an Array. Got: STRING" {
		t.Errorf("push builtin returned wrong result. Expected: Error: Argument to `push` must be an Array. Got: STRING. Got: %s", pushBuiltin.Fn(str).Inspect())
	}
	if pushBuiltin.Fn(arr, newEl).Inspect() != "[99, 7, 356, 666]" {
		t.Errorf("push builtin returned wrong result. Expected: [99, 7, 356, 666]. Got: %s", pushBuiltin.Fn(arr, newEl).Inspect())
	}
}

func TestPop(t *testing.T) {
	arr := &Array{Elements: []Object{
		&Integer{Value: 99},
		&Integer{Value: 7},
		&Integer{Value: 356},
	}}
	emptyArr := &Array{Elements: []Object{}}
	str := &String{Value: "neat string"}

	popBuiltin := GetBuiltinByName("pop")
	if popBuiltin.Fn(arr, str).Inspect() != "Error: Wrong number of arguments. Got: 2, Expected: 1" {
		t.Errorf("pop builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 2, Expected: 1. Got: %s", popBuiltin.Fn(arr, str).Inspect())
	}
	if popBuiltin.Fn(str).Inspect() != "Error: Argument to `pop` must be an Array. Got: STRING" {
		t.Errorf("pop builtin returned wrong result. Expected: Error: Argument to `pop` must be an Array. Got: STRING. Got: %s", popBuiltin.Fn(str).Inspect())
	}
	if popBuiltin.Fn(arr).Inspect() != "[99, 7]" {
		t.Errorf("pop builtin returned wrong result. Expected: [99, 7]. Got: %s", popBuiltin.Fn(arr).Inspect())
	}
	if popBuiltin.Fn(emptyArr) != nil {
		t.Errorf("pop builtin returned wrong result. Expected: null. Got: %s", popBuiltin.Fn(emptyArr))
	}
}

func TestSplit(t *testing.T) {
	str := &String{Value: "My name is brad"}
	splitOn := &String{Value: " "}
	array := &Array{Elements: []Object{}}

	splitBuiltin := GetBuiltinByName("split")
	if splitBuiltin.Fn(str).Inspect() != "Error: Wrong number of arguments. Got: 1, Expected: 2" {
		t.Errorf("pop builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 1, Expected: 2. Got: %s", splitBuiltin.Fn(str).Inspect())
	}
	if splitBuiltin.Fn(array, str).Inspect() != "Error: First argument to `split` must be a String. Got: ARRAY" {
		t.Errorf("split builtin returned wrong result. Expected: Error: First argument to `split` must be a String. Got:. Got: %s", splitBuiltin.Fn(array, str).Inspect())
	}
	if splitBuiltin.Fn(str, splitOn).Inspect() != "[My, name, is, brad]" {
		t.Errorf("split builtin returned wrong result. Expected: [My, name, is, brad]. Got: %s", splitBuiltin.Fn(str, splitOn).Inspect())
	}
}

func TestJoin(t *testing.T) {
	array := &Array{
		Elements: []Object{
			&String{Value: "My"},
			&String{Value: "name"},
			&String{Value: "is"},
			&String{Value: "brad"},
		},
	}
	mixedArray := &Array{
		Elements: []Object{
			&String{Value: "My"},
			&String{Value: "name"},
			&Boolean{Value: true},
		},
	}
	joinOn := &String{Value: " "}
	notAnArray := &String{Value: "not an array"}

	joinBuiltin := GetBuiltinByName("join")
	if joinBuiltin.Fn(array).Inspect() != "Error: Wrong number of arguments. Got: 1, Expected: 2" {
		t.Errorf("join builtin returned wrong result. Expected: Error: Wrong number of arguments. Got: 1, Expected: 2. Got: %s", joinBuiltin.Fn(array).Inspect())
	}
	if joinBuiltin.Fn(notAnArray, joinOn).Inspect() != "Error: First argument to `join` must be an Array. Got: STRING" {
		t.Errorf("join builtin returned wrong result. Expected: Error: First argument to `join` must be an Array. Got:. Got: %s", joinBuiltin.Fn(notAnArray, joinOn).Inspect())
	}
	if joinBuiltin.Fn(mixedArray, joinOn).Inspect() != "Error: You can only join an array of all strings. Illegal type found: BOOLEAN" {
		t.Errorf("join builtin returned wrong result. Expected: Error: You can only join an array of all strings. Illegal type found: BOOLEAN. Got: %s", joinBuiltin.Fn(mixedArray, joinOn).Inspect())
	}
	if joinBuiltin.Fn(array, joinOn).Inspect() != "My name is brad" {
		t.Errorf("split builtin returned wrong result. Expected: My name is brad. Got: %s", joinBuiltin.Fn(array, joinOn).Inspect())
	}
}
