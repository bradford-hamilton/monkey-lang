package compiler

// A symbol table is a data structure used in interpreters & compilers to associate identifiers
// with information. It can be used in every phase, from lexing to code generation, to store and
// retrieve information about a given identifier (which can be called a symbol). Information such
// as its location, its scope, whether it was previously declared or not, of which type the
// associated value is, and anything else that useful while interpreting or compiling.

// SymbolScope - Type alias for a string. Used in a Symbol to define it's scope.
type SymbolScope string

// Define all our scopes
const (
	GlobalScope   SymbolScope = "GLOBAL"
	LocalScope    SymbolScope = "LOCAL"
	BuiltinScope  SymbolScope = "BUILTIN"
	FreeScope     SymbolScope = "FREE"
	FunctionScope SymbolScope = "FUNCTION"
)

// Symbol - Holds all the necessary info about a symbol - Name, Scope, and Index
type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

// SymbolTable holds a "store" which is a map of strings to Symbols, an int of number of definitions,
// and "Outer" which defines it's parent scope
type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
	Outer          *SymbolTable
	FreeSymbols    []Symbol
}

// NewSymbolTable creates and returns a pointer to a symbol table initialized with a "store"
func NewSymbolTable() *SymbolTable {
	s := make(map[string]Symbol)
	freeSymbols := []Symbol{}

	return &SymbolTable{
		store:       s,
		FreeSymbols: freeSymbols,
	}
}

// NewEnclosedSymbolTable takes an outer parent symbol table and returns a pointer to the new
// enclosed SymbolTable after attaching the outer parent to the new one
func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

// Define takes a name and creates a symbol with the name, an index, and assigns scope. It then assigns
// the symbol to the SymbolTable's store, increases numDefinitions and returns the symbol.
func (s *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: s.numDefinitions,
	}

	if s.Outer == nil {
		symbol.Scope = GlobalScope
	} else {
		symbol.Scope = LocalScope
	}

	s.store[name] = symbol
	s.numDefinitions++

	return symbol
}

// DefineBuiltin creates and returns a symbol within builtin scope
func (s *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: index,
		Scope: BuiltinScope,
	}
	s.store[name] = symbol

	return symbol
}

func (s *SymbolTable) defineFree(original Symbol) Symbol {
	s.FreeSymbols = append(s.FreeSymbols, original)

	symbol := Symbol{
		Name:  original.Name,
		Index: len(s.FreeSymbols) - 1,
	}
	symbol.Scope = FreeScope
	s.store[original.Name] = symbol

	return symbol
}

// DefineFunctionName creates a new Symbol with FunctionScope and adds it to the s.store
func (s *SymbolTable) DefineFunctionName(name string) Symbol {
	symbol := Symbol{
		Name:  name,
		Index: 0,
		Scope: FunctionScope,
	}
	s.store[name] = symbol

	return symbol
}

// Resolve takes a name, looks for it in the SymbolTable's store, and returns it if found along
// with a boolean representing whether it was found
func (s *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := s.store[name]

	if !ok && s.Outer != nil {
		obj, ok = s.Outer.Resolve(name)
		if !ok {
			return obj, ok
		}

		if obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := s.defineFree(obj)

		return free, true
	}

	return obj, ok
}
