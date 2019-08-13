package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// InfixExpression - holds the token, the expression to the left of it, a string version of
// the operator, and the expression to the right of it
type InfixExpression struct {
	Token    token.Token // The operator token (+, -, *, etc)
	Left     Expression
	Operator string // string (examples: "+", "-", "*", etc)
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the InfixExpression's Literal and satisfies the Node interface.
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String - returns a string representation of the left side expression, the operator, and
// the right side expression (5 * 5) and satisfies our Node interface
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
