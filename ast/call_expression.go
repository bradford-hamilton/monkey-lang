package ast

import (
	"bytes"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// CallExpression - holds the token, the function expression, and its arguments ([]Expression).
// Structure: <expression>(<comma separated expressions>)
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns the CallExpression's Literal and satisfies the Node interface.
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

// String - returns a string representation of the CallExpression. Prints the function
// and its arguments. Satisfies our Node interface
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
