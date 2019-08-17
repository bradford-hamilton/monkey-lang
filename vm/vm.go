package vm

import (
	"fmt"

	"github.com/bradford-hamilton/monkey-lang/code"
	"github.com/bradford-hamilton/monkey-lang/compiler"
	"github.com/bradford-hamilton/monkey-lang/object"
)

// StackSize is an integer defining the size of our stack
const StackSize = 2048

// True, False, and Null - immutable & unique. No need to create new objects in memory each time we
// need one. Makes comparison easier as well because they always point to same place in memory

// True - Pointer to a Monkey object.Boolean with value true
var True = &object.Boolean{Value: true}

// False - Pointer to a Monkey object.Boolean of value false
var False = &object.Boolean{Value: false}

// Null - Pointer to a Monkey object.Null
var Null = &object.Null{}

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
		case code.OpPop:
			vm.pop()
		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}
		case code.OpTrue:
			err := vm.push(True)
			if err != nil {
				return err
			}
		case code.OpFalse:
			err := vm.push(False)
			if err != nil {
				return err
			}
		case code.OpEqualEqual, code.OpNotEqual, code.OpGreaterThan:
			err := vm.executeComparison(op)
			if err != nil {
				return err
			}
		case code.OpBang:
			err := vm.executeBangOperator()
			if err != nil {
				return err
			}
		case code.OpMinus:
			err := vm.executeMinusOperator()
			if err != nil {
				return err
			}
		case code.OpJump:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip = pos - 1
		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				ip = pos - 1
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

func (vm *VM) executeComparison(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()

	if left.Type() == object.IntegerObj || right.Type() == object.IntegerObj {
		return vm.executeIntegerComparison(op, left, right)
	}

	switch op {
	case code.OpEqualEqual:
		return vm.push(nativeBoolToBooleanObj(right == left))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObj(right != left))
	default:
		return fmt.Errorf("Unknown operator: %d (%s %s)", op, left.Type(), right.Type())
	}
}

func (vm *VM) executeIntegerComparison(op code.Opcode, left, right object.Object) error {
	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	switch op {
	case code.OpEqualEqual:
		return vm.push(nativeBoolToBooleanObj(rightValue == leftValue))
	case code.OpNotEqual:
		return vm.push(nativeBoolToBooleanObj(rightValue != leftValue))
	case code.OpGreaterThan:
		return vm.push(nativeBoolToBooleanObj(leftValue > rightValue))
	default:
		return fmt.Errorf("Unknown operator: %d", op)
	}
}

func nativeBoolToBooleanObj(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}

func (vm *VM) executeBangOperator() error {
	operand := vm.pop()

	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	default:
		return vm.push(False)
	}
}

func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()

	if operand.Type() != object.IntegerObj {
		return fmt.Errorf("Unsupported type for negation: %s", operand.Type())
	}

	value := operand.(*object.Integer).Value

	return vm.push(&object.Integer{Value: -value})
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	default:
		return true
	}
}
