package object

// Environment holds a store of key value pairs and a pointer to an "outer", enclosing environment
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get retrieves a key from an Environment's store by name. If it does not find it, it recursively looks
// for the key in the enclosing environment(s)
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets a key to an Environment's store by name
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// NewEnvironment creates and returns a pointer to an Environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new Environment and attaches the outer environment
// that's passed in, to the new environment, as it's enclosing environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}
