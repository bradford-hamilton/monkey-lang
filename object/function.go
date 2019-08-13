package object

import (
	"bytes"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/ast"
)

// Function holds Parameters as a slice of *Identifier, a Body which is a *ast.BlockStatement
// and a pointer to it's environment
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns our Function's ObjectType (FunctionObj)
func (f *Function) Type() ObjectType { return FunctionObj }

// Inspect returns a string representation of the function definition
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("func")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}
