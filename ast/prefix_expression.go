package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// PrefixExpression - holds the token, a string version of the operator, and the expression to the right of it
type PrefixExpression struct {
	Token    token.Token // The prefix token (! or -)
	Operator string      // String (either "!" or "-")
	Right    Expression  // The expression to the right of the operator
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the PrefixExpression's Literal and satisfies the Node interface.
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String - returns a string representation of the operator followed by it's expression to the
// right (-5) and satisfies our Node interface
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
