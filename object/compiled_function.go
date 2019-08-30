package object

import (
	"fmt"

	"github.com/bradford-hamilton/monkey-lang/code"
)

// CompiledFunction holds the instructions we get from the compilation of a function
// literal and is an object.Object, which means we can add it as a constant to our
// compiler.Bytecode and load it in the VM. It also holds the NumLocals which we pass
// to the VM to allocate the correct amount of stack space ("hole") to save the local
// bindings
type CompiledFunction struct {
	Instructions  code.Instructions
	NumLocals     int
	NumParameters int
}

// Type returns our CompiledFunction's ObjectType (CompiledFunctionObj)
func (cf *CompiledFunction) Type() ObjectType { return CompiledFunctionObj }

// Inspect returns the string "CompiledFunction[address]" - Address of 0th element in base 16 notation, with leading 0x
func (cf *CompiledFunction) Inspect() string { return fmt.Sprintf("CompiledFunction[%p]", cf) }
