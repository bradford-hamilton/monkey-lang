package vm

import (
	"fmt"

	"github.com/bradford-hamilton/monkey-lang/code"
	"github.com/bradford-hamilton/monkey-lang/compiler"
	"github.com/bradford-hamilton/monkey-lang/object"
)

// StackSize is an integer defining the size of our stack
const StackSize = 2048

// VM defines our Virtual Machine. It holds our constant pool, instructions, a stack, and an integer (index)
// that points to the next free slot in the stack
type VM struct {
	constants    []object.Object
	instructions code.Instructions
	stack        []object.Object
	sp           int // Stack pointer: always points to the next free slot in the stack. Top of stack is stack[ip-1]
}

// New initializers and returns a pointer to a VM. It takes bytecode and sets the bytecode's instructions
// and constants to the VM, creates a new stack with StackSize number of elements, and initalizes the ip to 0
func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

// StackTop returns the element currently at the top of the stack
func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

// Run runs our VM and starts the fetch-decode-execute cycle
func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (vm *VM) push(obj object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = obj
	vm.sp++

	return nil
}
