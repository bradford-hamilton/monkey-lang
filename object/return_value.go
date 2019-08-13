package object

// ReturnValue type holds a return value
type ReturnValue struct {
	Value Object
}

// Type returns our ReturnValue's ObjectType (ReturnValueObj)
func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }

// Inspect returns a string representation of the ReturnValue's Value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
