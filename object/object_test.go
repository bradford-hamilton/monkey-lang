package object

import (
	"fmt"
	"testing"
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

func TestArray(t *testing.T) {
	elements := []Object{
		&Integer{
			Value: 1,
		},
		&Integer{
			Value: 2,
		},
		&Integer{
			Value: 3,
		},
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
	b := &Boolean{Value: true}

	if b.Type() != BooleanObj {
		t.Errorf("b.Type() returned wrong type. Expected: BooleanObj. Got: %s", b.Type())
	}

	if b.Inspect() != "true" {
		t.Errorf("arr.Inspect() returned wrong string representation. Expected: true. Got: %s", b.Inspect())
	}
}

func TestClosure(t *testing.T) {
	cl := &Closure{
		Fn: &CompiledFunction{
			Instructions:  []byte("OpDoesntMatter"),
			NumLocals:     1,
			NumParameters: 1,
		},
		Free: []Object{&Integer{Value: 8}},
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
