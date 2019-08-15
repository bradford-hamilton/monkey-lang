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

// LastPoppedStackElement returns last popped element on the top of the stack. We do not explicitly
// set them to nil or remove them when calling pop so it will point to last popped.
func (vm *VM) LastPoppedStackElement() object.Object {
	return vm.stack[vm.sp]
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
		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
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

func (vm *VM) pop() object.Object {
	obj := vm.stack[vm.sp-1]
	vm.sp--
	return obj
}

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	leftType := left.Type()
	rightType := right.Type()

	if leftType == object.IntegerObj && rightType == object.IntegerObj {
		return vm.executeBinaryIntegerOperation(op, left, right)
	}

	return fmt.Errorf("Unsupported types for binary operation: %s %s", leftType, rightType)
}

func (vm *VM) executeBinaryIntegerOperation(op code.Opcode, left, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	var result int64

	switch op {
	case code.OpAdd:
		result = leftValue + rightValue
	case code.OpSub:
		result = leftValue - rightValue
	case code.OpMul:
		result = leftValue * rightValue
	case code.OpDiv:
		result = leftValue / rightValue
	default:
		return fmt.Errorf("Unknown integer operator: %d", op)
	}

	return vm.push(&object.Integer{Value: result})
}
