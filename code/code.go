package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Instructions - Type alias for a slice of byte.
type Instructions []byte

// String is a sort of mini dissassembler to print our bytecode in human readable format
func (ins Instructions) String() string {
	var out bytes.Buffer

	i := 0
	for i < len(ins) {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])

		fmt.Fprintf(&out, "%04d %s\n", i, ins.fmtInstruction(def, operands))

		i += 1 + read
	}

	return out.String()
}

func (ins Instructions) fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)

	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: operand len %d does not match defined %d\n", len(operands), operandCount)
	}

	switch operandCount {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	case 2:
		return fmt.Sprintf("%s %d %d", def.Name, operands[0], operands[1])
	}

	return fmt.Sprintf("ERROR: unhandled operandCount for %s\n", def.Name)
}

// Opcode is merely a byte
type Opcode byte

// Define our opcode types
const (
	// Constants
	OpConstant Opcode = iota

	// Aritmatic & stack pop
	OpAdd
	OpPop
	OpSub
	OpMul
	OpDiv

	// Boolean
	OpTrue
	OpFalse

	// Comparison
	OpEqualEqual
	OpNotEqual
	OpGreaterThan

	// Prefix/unary
	OpMinus
	OpBang

	// Jump for conditionals
	OpJumpNotTruthy // Jump to alternative if consequence is not truthy
	OpJump          // Jump no matter what (if we evaluate consequence and dont want the alternative)

	// Null
	OpNull

	// Get and Set global variables
	OpGetGlobal
	OpSetGlobal

	// Data structures and index access
	OpArray
	OpHash
	OpIndex

	// Call expression, return with value, return without value
	OpCall
	OpReturnValue // Tells VM to leave value on top of stack
	OpReturn      // Tells VM implicit return of Null

	// Get and Set local bindings
	OpGetLocal
	OpSetLocal

	// Builtin functions
	OpGetBuiltin

	// Closures and it's variables
	OpClosure
	OpGetFree
)

// Definition for an opcode. Name helps to make an Opcode readable and OperandWidths
// contains the number of bytes each operand takes up.
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant:      {"OpConstant", []int{2}}, // OpConstant is 2 bytes wide which makes it a uint16 (limits value to 65535)
	OpAdd:           {"OpAdd", []int{}},
	OpPop:           {"OpPop", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpMul", []int{}},
	OpDiv:           {"OpDiv", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEqualEqual:    {"OpEqualEqual", []int{}},
	OpNotEqual:      {"OpNotEqual", []int{}},
	OpGreaterThan:   {"OpGreaterThan", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},
	OpNull:          {"OpNull", []int{}},
	OpGetGlobal:     {"OpGetGlobal", []int{2}},
	OpSetGlobal:     {"OpSetGlobal", []int{2}},
	OpArray:         {"OpArray", []int{2}},
	OpHash:          {"OpHash", []int{2}},
	OpIndex:         {"OpIndex", []int{}},
	OpCall:          {"OpCall", []int{1}}, // OpCall is 1 byte wide which makes it a uint8 (limits value to 255)
	OpReturnValue:   {"OpReturnValue", []int{}},
	OpReturn:        {"OpReturn", []int{}},
	OpGetLocal:      {"OpGetLocal", []int{1}},
	OpSetLocal:      {"OpSetLocal", []int{1}},
	OpGetBuiltin:    {"OpGetBuiltin", []int{1}},

	// Has two operands, first is two bytes wide - the constant index. Specifies where in the constant pool we
	// can find the *object.CompiledFunction that's to be converted into a closure. It's two bytes wide because
	// the operand of OpConstant is also two bytes wide. The second operand, one byte wide, specifies how many
	// free variables sit on the stack and need to be transferred to the about-to-be-created closure.
	OpClosure: {"OpClosure", []int{2, 1}},
	OpGetFree: {"OpGetFree", []int{1}},
}

// Lookup finds a definition by opcode. It returns it if it is found otherwise returns an error
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("Opcode %d undefined", op)
	}

	return def, nil
}

// Make creates our bytecode. First we find out how long the resulting instruction is going
// to be. That allows us to allocate a byte slice with the proper length. We then add the
// Opcode as its first byte by casting it into one. Then we iterate over the defined OperandWidths,
// take the matching element from operands, and put it in the instruction. After encoding the two
// byte operand in big endian, we increment offset by its width.
func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1

	for i, o := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		case 1:
			instruction[offset] = byte(o)
		}
		offset += width
	}

	return instruction
}

// ReadOperands is Make's counterpart. It decodes the operands of a bytecode instruction.
func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

// ReadUint16 turns a byte sequence (Instructions) into a uint16
func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

// ReadUint8 turns a byte sequence (Instructions) into a uint16
func ReadUint8(ins Instructions) uint8 { return uint8(ins[0]) }
