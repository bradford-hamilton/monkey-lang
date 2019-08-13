package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// LetStatement - Name holds the identifier of the binding and Value for the expression
// that produces the value.
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the LetStatement's Literal and satisfies the Node interface.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// String - returns a string representation of the LetStatement and satisfies our Node interface
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
