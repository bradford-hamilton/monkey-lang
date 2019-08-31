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
