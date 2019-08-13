package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// IndexExpression - holds the '[' token, the object being accessed, and the index
type IndexExpression struct {
	Token token.Token // The '[' token
	Left  Expression  // The object being accessed
	Index Expression  // Can be any expression, but must produce an integer
}

func (ie *IndexExpression) expressionNode() {}

// TokenLiteral returns the IndexExpression's Literal and satisfies the Node interface.
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

// String - returns string representation of the IndexExpression: (leftExpr[indexExpr]).
// Satisfies our Node interface
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
