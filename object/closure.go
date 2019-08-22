package object

import "fmt"

// Closure holds a pointer to its compiled function and a slice of its free objects (variables it
// has access to that are not in either global or local scope)
type Closure struct {
	Fn   *CompiledFunction
	Free []Object
}

// Type returns our Closure's ObjectType (ClosureObj)
func (c *Closure) Type() ObjectType { return ClosureObj }

// Inspect returns a string representation of the Closure with its address
func (c *Closure) Inspect() string {
	return fmt.Sprintf("Closure[%p]", c)
}
