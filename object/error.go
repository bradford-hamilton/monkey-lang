package object

// Error type holds an error Message
type Error struct {
	Message string
}

// Type returns our Error's ObjectType (ErrorObj)
func (e *Error) Type() ObjectType { return ErrorObj }

// Inspect returns an error message string
func (e *Error) Inspect() string { return "Error: " + e.Message }
