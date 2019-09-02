package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// PostfixExpression - holds the token we're operating on and postfix operator (--, ++)
type PostfixExpression struct {
	Token    token.Token
	Operator string
}

func (pe *PostfixExpression) expressionNode() {}

// TokenLiteral returns the PostfixExpression's Literal and satisfies the Node interface.
func (pe *PostfixExpression) TokenLiteral() string { return pe.Token.Literal }

// String - returns a string representation of the left side expression, the operator, and
// the right side expression (5 * 5) and satisfies our Node interface
func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Token.Literal)
	out.WriteString(pe.Operator)
	out.WriteString(")")

	return out.String()
}
