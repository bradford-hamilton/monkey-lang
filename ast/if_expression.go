package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// IfExpression - holds the token, the condition expression and the consequence & alternative
// block statements. Structure: if (<condition>) <consequence> else <alternative>
type IfExpression struct {
	Token       token.Token // The IF token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the IfExpression's Literal and satisfies the Node interface.
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String - returns a string representation of the IfExpression with the consequence and
// also the alteritive if it is not nil. Satisfies our Node interface
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
