package vm

import (
	"github.com/bradford-hamilton/monkey-lang/code"
	"github.com/bradford-hamilton/monkey-lang/object"
)

// Frame - Data structure that holds execution-relevant information. Short for "call frame" or "stack frame"
// and sometimes "activation record". On real machines a frame is not something separate from, but a designated
// part of "the stack". It's where the return address, the arguments to the current function, and it's local
// variables are stored. In VM land we don't have to use the stack. We're not constrained by standardized
// calling conventions and other much too real things, like real memory addresses and locations. Since we can
// store frames anywhere we like. What's kept on the stack and what's not differs from VM to VM. Some keep
// everything on the stack, others only the return address, some only the local variables, some the local
// variables and the arguments of the function call. The implementation depends on the language being
// implemented, the requirements in regards to concurrency and performance, the host language, and more.
// We are choosing the way that is easiest to build, understand, extend, etc.
type Frame struct {
	closure     *object.Closure
	ip          int
	basePointer int // Keeps track of the stacks pointer's value before we execute a function so we can restore stack to this value after executing
}

// Instructions returns the frame's function's instructions
func (f *Frame) Instructions() code.Instructions {
	return f.closure.Fn.Instructions
}

// NewFrame takes a pointer to a compiled function, creates a frame with it, sets the instruction
// pointer to -1, and returns a pointer to the frame.
func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		closure:     cl,
		ip:          -1,
		basePointer: basePointer, // The pointer that points to the bottom of the stack of the current call frame. (Sometimes called "FramePointer")
	}
}
