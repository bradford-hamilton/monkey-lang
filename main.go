package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/compiler"
	"github.com/bradford-hamilton/monkey-lang/evaluator"
	"github.com/bradford-hamilton/monkey-lang/lexer"
	"github.com/bradford-hamilton/monkey-lang/object"
	"github.com/bradford-hamilton/monkey-lang/parser"
	"github.com/bradford-hamilton/monkey-lang/repl"
	"github.com/bradford-hamilton/monkey-lang/vm"
)

func main() {
	engine := flag.String("engine", "vm", "Engine options are \"vm\" or \"eval\"")
	console := flag.Bool("console", false, "Provide console flag to enter interactive repl")
	flag.Parse()

	if *engine != "vm" && *engine != "eval" {
		fmt.Printf("Engine must be either 'vm' or 'eval', got %s\n", *engine)
		return
	}

	if *console {
		repl.Start(os.Stdin, os.Stdout, engine)
	} else {
		if len(flag.Args()) != 1 {
			fmt.Println("Incorrect usage. Usage: `monkey [option...] filePath`")
			return
		}

		filePath := flag.Args()[0]

		contents, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Failure to read file '%s'. Err: %s", string(contents), err)
		}

		l := lexer.New(string(contents))
		p := parser.New(l)
		program := p.ParseProgram()

		if *engine == "vm" {
			compileBytecodeAndRun(program)
		} else {
			evaluateAst(program)
		}
	}
}

// Evaluate the AST with evaluator and print result
func evaluateAst(program *ast.RootNode) {
	env := object.NewEnvironment()
	result := evaluator.Eval(program, env)
	fmt.Println(result.Inspect())
}

// Compile program to bytecode, pass to VM, and run. Prints the last popped stack element (result)
func compileBytecodeAndRun(program *ast.RootNode) {
	comp := compiler.New()
	err := comp.Compile(program)
	if err != nil {
		fmt.Printf("compiler error: %s", err)
	}

	vm := vm.New(comp.Bytecode())
	err = vm.Run()
	if err != nil {
		fmt.Printf("vm error: %s", err)
	}

	stackElem := vm.LastPoppedStackElement()
	fmt.Println(stackElem.Inspect())
}
