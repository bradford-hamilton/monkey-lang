package vm

import (
	"fmt"
	"testing"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/compiler"
	"github.com/bradford-hamilton/monkey-lang/lexer"
	"github.com/bradford-hamilton/monkey-lang/object"
	"github.com/bradford-hamilton/monkey-lang/parser"
)

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"50 / 2 * 2 + 10 - 5", 55},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * (2 + 10)", 60},
		{"-5", -5},
		{"-10", -10},
		{"-50 + 100 + -50", 0},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	runVMTests(t, tests)
}

func TestBooleanExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!(if (false) { 5; })", true},
	}

	runVMTests(t, tests)
}

func TestConditionals(t *testing.T) {
	tests := []vmTestCase{
		{"if (true) { 10 }", 10},
		{"if (true) { 10 } else { 20 }", 10},
		{"if (false) { 10 } else { 20 } ", 20},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 > 2) { 10 }", Null},
		{"if (false) { 10 }", Null},
		{"if ((if (false) { 10 })) { 10 } else { 20 }", 20},
	}

	runVMTests(t, tests)
}

func TestGlobalLetStatements(t *testing.T) {
	tests := []vmTestCase{
		{"let one = 1; one", 1},
		{"let one = 1; let two = 2; one + two", 3},
		{"let one = 1; let two = one + one; one + two", 3},
	}

	runVMTests(t, tests)
}

func TestStringExpressions(t *testing.T) {
	tests := []vmTestCase{
		{`"monkey"`, "monkey"},
		{`"mon" + "key"`, "monkey"},
		{`"mon" + "key" + "banana"`, "monkeybanana"},
	}

	runVMTests(t, tests)
}

func TestArrayLiterals(t *testing.T) {
	tests := []vmTestCase{
		{"[]", []int{}},
		{"[1, 2, 3]", []int{1, 2, 3}},
		{"[1 + 2, 3 * 4, 5 + 6]", []int{3, 12, 11}},
	}

	runVMTests(t, tests)
}

func TestHashLiterals(t *testing.T) {
	tests := []vmTestCase{
		{
			"{}", map[object.HashKey]int64{},
		},
		{
			"{1: 2, 2: 3}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 1}).HashKey(): 2,
				(&object.Integer{Value: 2}).HashKey(): 3,
			},
		},
		{
			"{1 + 1: 2 * 2, 3 + 3: 4 * 4}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 2}).HashKey(): 4,
				(&object.Integer{Value: 6}).HashKey(): 16,
			},
		},
	}

	runVMTests(t, tests)
}

func TestIndexExpressions(t *testing.T) {
	tests := []vmTestCase{
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][0 + 2]", 3},
		{"[[1, 1, 1]][0][0]", 1},
		{"[][0]", Null},
		{"[1, 2, 3][99]", Null},
		{"[1][-1]", Null},
		{"{1: 1, 2: 2}[1]", 1},
		{"{1: 1, 2: 2}[2]", 2},
		{"{1: 1}[0]", Null},
		{"{}[0]", Null},
	}

	runVMTests(t, tests)
}

func TestCallingFunctionsWithoutArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let fivePlusTen = func() { 5 + 10; };
				fivePlusTen();
		`,
			expected: 15,
		},
		{
			input: `
				let one = func() { 1; };
				let two = func() { 2; };
				one() + two()
		`,
			expected: 3,
		},
		{
			input: `
				let a = func() { 1 };
				let b = func() { a() + 1 };
				let c = func() { b() + 1 };
				c();
		`,
			expected: 3,
		},
	}

	runVMTests(t, tests)
}

func TestFunctionsWithReturnStatement(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let earlyExit = func() { return 99; 100; };
				earlyExit();
		`,
			expected: 99,
		},
		{
			input: `
				let earlyExit = func() { return 99; return 100; };
				earlyExit();
		`,
			expected: 99,
		},
	}

	runVMTests(t, tests)
}

func TestFunctionsWithoutReturnValue(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let noReturn = func() { };
				noReturn();
		`,
			expected: Null,
		},
		{
			input: `
				let noReturn = func() { };
				let noReturnTwo = func() { noReturn(); };
				noReturn();
				noReturnTwo();
		`,
			expected: Null,
		},
	}

	runVMTests(t, tests)
}

func TestFirstClassFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let returnsOne = func() { 1; };
				let returnsOneReturner = func() { returnsOne; };
				returnsOneReturner()();
		`,
			expected: 1,
		},
		{
			input: `
				let returnsOneReturner = func() {
					let returnsOne = func() { 1; };
					returnsOne;
				};
				returnsOneReturner()();
		`,
			expected: 1,
		},
	}

	runVMTests(t, tests)
}

func TestCallingFunctionsWithBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let one = func() { let one = 1; one };
				one();
		`,
			expected: 1,
		},
		{
			input: `
				let oneAndTwo = func() { let one = 1; let two = 2; one + two; };
				oneAndTwo();
		`,
			expected: 3,
		},
		{
			input: `
				let oneAndTwo = func() { let one = 1; let two = 2; one + two; };
				let threeAndFour = func() { let three = 3; let four = 4; three + four; };
				oneAndTwo() + threeAndFour();
		`,
			expected: 10,
		},
		{
			input: `
				let firstFoobar = func() { let foobar = 50; foobar; };
				let secondFoobar = func() { let foobar = 100; foobar; };
				firstFoobar() + secondFoobar();
		`,
			expected: 150,
		},
		{
			input: `
				let globalSeed = 50;
				let minusOne = func() {
					let num = 1;
					globalSeed - num;
				}
				let minusTwo = func() {
					let num = 2;
					globalSeed - num;
				}
				minusOne() + minusTwo();
		`,
			expected: 97,
		},
	}

	runVMTests(t, tests)
}

func TestCallingFunctionsWithArgumentsAndBindings(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let identity = func(a) { a; };
				identity(4);
		`,
			expected: 4,
		},
		{
			input: `
				let sum = func(a, b) { a + b; };
				sum(1, 2);
		`,
			expected: 3,
		},
		{
			input: `
				let sum = func(a, b) {
					let c = a + b;
					c;
				};
				sum(1, 2);
		`,
			expected: 3,
		},
		{
			input: `
				let sum = func(a, b) {
					let c = a + b;
					c;
				};
				sum(1, 2) + sum(3, 4);`,
			expected: 10,
		},
		{
			input: `
				let sum = func(a, b) {
					let c = a + b;
					c;
				};
				let outer = func() {
					sum(1, 2) + sum(3, 4);
				};
				outer();
		`,
			expected: 10,
		},
		{
			input: `
				let globalNum = 10;

				let sum = func(a, b) {
					let c = a + b;
					c + globalNum;
				};

				let outer = func() {
					sum(1, 2) + sum(3, 4) + globalNum;
				};

				outer() + globalNum;
		`,
			expected: 50,
		},
	}

	runVMTests(t, tests)
}

func TestCallingFunctionsWithWrongArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			input:    `func() { 1; }(1);`,
			expected: `Wrong number of arguments. Expected: 0. Got: 1`,
		},
		{
			input:    `func(a) { a; }();`,
			expected: `Wrong number of arguments. Expected: 1. Got: 0`,
		},
		{
			input:    `func(a, b) { a + b; }(1);`,
			expected: `Wrong number of arguments. Expected: 2. Got: 1`,
		},
	}

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err == nil {
			t.Fatalf("expected VM error but resulted in none.")
		}

		if err.Error() != tt.expected {
			t.Fatalf("wrong VM error. Want: %q. Got: %q", tt.expected, err)
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []vmTestCase{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{
			`len(1)`,
			&object.Error{
				Message: "Argument to `len` not supported. Got: INTEGER",
			},
		},
		{`len("one", "two")`,
			&object.Error{
				Message: "Wrong number of arguments. Got: 2, Expected: 1",
			},
		},
		{`len([1, 2, 3])`, 3},
		{`len([])`, 0},
		{`puts("hello", "world!")`, Null},
		{`first([1, 2, 3])`, 1},
		{`first([])`, Null},
		{`first(1)`,
			&object.Error{
				Message: "Argument to `first` must be an Array. Got: INTEGER",
			},
		},
		{`last([1, 2, 3])`, 3},
		{`last([])`, Null},
		{`last(1)`,
			&object.Error{
				Message: "Argument to `last` must be an Array. Got: INTEGER",
			},
		},
		{`rest([1, 2, 3])`, []int{2, 3}},
		{`rest([])`, Null},
		{`push([], 1)`, []int{1}},
		{`push(1, 1)`,
			&object.Error{
				Message: "Argument to `push` must be an Array. Got: INTEGER",
			},
		},
	}

	runVMTests(t, tests)
}

func TestClosures(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let newClosure = func(a) {
					func() { a; };
				};
				let closure = newClosure(99);
				closure();
		`,
			expected: 99,
		},
		{
			input: `
				let newAdder = func(a, b) {
					func(c) { a + b + c };
				};
				let adder = newAdder(1, 2);
				adder(8);
		`,
			expected: 11,
		},
		{
			input: `
				let newAdder = func(a, b) {
					let c = a + b;
					func(d) { c + d };
				};
				let adder = newAdder(1, 2);
				adder(8);
		`,
			expected: 11,
		},
		{
			input: `
				let newAdderOuter = func(a, b) {
					let c = a + b;
					func(d) {
						let e = d + c;
						func(f) { e + f; };
					};
				};
				let newAdderInner = newAdderOuter(1, 2)
				let adder = newAdderInner(3);
				adder(8);
		`,
			expected: 14,
		},
		{
			input: `
				let a = 1;
				let newAdderOuter = func(b) {
					func(c) {
						func(d) { a + b + c + d };
					};
				};
				let newAdderInner = newAdderOuter(2)
				let adder = newAdderInner(3);
				adder(8);
		`,
			expected: 14,
		},
		{
			input: `
				let newClosure = func(a, b) {
					let one = func() { a; };
					let two = func() { b; };
					func() { one() + two(); };
				};
				let closure = newClosure(9, 90);
				closure();
		`,
			expected: 99,
		},
	}

	runVMTests(t, tests)
}

func TestRecursiveFunctions(t *testing.T) {
	tests := []vmTestCase{
		{
			input: `
				let countDown = func(x) {
					if (x == 0) {
						return 0;
					} else {
						countDown(x - 1);
					}
				};
				countDown(1);
		`,
			expected: 0,
		},
		{
			input: `
				let countDown = func(x) {
					if (x == 0) {
						return 0;
					} else {
						countDown(x - 1);
					}
				};
				let wrapper = func() {
					countDown(1);
				};
				wrapper();
		`,
			expected: 0,
		},
		{
			input: `
				let wrapper = func() {
					let countDown = func(x) {
						if (x == 0) {
							return 0;
						} else {
							countDown(x - 1);
						}
					};
					countDown(1);
				};
				wrapper();
		`,
			expected: 0,
		},
	}

	runVMTests(t, tests)
}

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVMTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.Bytecode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.LastPoppedStackElement()

		testExpectedObject(t, tt.expected, stackElem)
	}
}

func parse(input string) *ast.RootNode {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	case bool:
		err := testBooleanObject(bool(expected), actual)
		if err != nil {
			t.Errorf("testBooleanObject failed: %s", err)
		}
	case *object.Null:
		if actual != Null {
			t.Errorf("object is not Null : %T (%+v)", actual, actual)
		}
	case string:
		err := testStringObject(expected, actual)
		if err != nil {
			t.Errorf("testStringObject failed: %s", err)
		}
	case []int:
		array, ok := actual.(*object.Array)
		if !ok {
			t.Errorf("Object is not an Array: %T (%+v)", actual, actual)
			return
		}

		if len(array.Elements) != len(expected) {
			t.Errorf("Wrong num of elements. Want: %d. Got: %d", len(expected), len(array.Elements))
			return
		}

		for i, expectedElem := range expected {
			err := testIntegerObject(int64(expectedElem), array.Elements[i])
			if err != nil {
				t.Errorf("testIntegerObject failed: %s", err)
			}
		}
	case map[object.HashKey]int64:
		hash, ok := actual.(*object.Hash)
		if !ok {
			t.Errorf("object is not Hash. got=%T (%+v)", actual, actual)
			return
		}

		if len(hash.Pairs) != len(expected) {
			t.Errorf("hash has wrong number of Pairs. want=%d, got=%d",
				len(expected), len(hash.Pairs))
			return
		}

		for expectedKey, expectedValue := range expected {
			pair, ok := hash.Pairs[expectedKey]
			if !ok {
				t.Errorf("no pair for given key in Pairs")
			}

			err := testIntegerObject(expectedValue, pair.Value)
			if err != nil {
				t.Errorf("testIntegerObject failed: %s", err)
			}
		}
	case *object.Error:
		errObj, ok := actual.(*object.Error)
		if !ok {
			t.Errorf("object is not Error: %T (%+v)", actual, actual)
			return
		}
		if errObj.Message != expected.Message {
			t.Errorf("Wrong error message. Expected: %q. Got: %q", expected.Message, errObj.Message)
		}
	}
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not an Integer. Got: %T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. Expected: %d. Got: %d", result.Value, expected)
	}

	return nil
}

func testBooleanObject(expected bool, actual object.Object) error {
	result, ok := actual.(*object.Boolean)
	if !ok {
		return fmt.Errorf("Object is not Boolean. Got: %T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("Object has wrong value. Expected: %t, Got: %t", result.Value, expected)
	}

	return nil
}

func testStringObject(expected string, actual object.Object) error {
	result, ok := actual.(*object.String)
	if !ok {
		return fmt.Errorf("Object is not a String. Got: %T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("Object has wrong value. Got: %q. Expected: %q", result.Value, expected)
	}

	return nil
}
