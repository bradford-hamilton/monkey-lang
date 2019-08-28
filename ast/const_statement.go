package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// ConstStatement - Name holds the identifier of the binding and Value for the expression
// that produces the value.
type ConstStatement struct {
	Token token.Token // The token.Const token
	Name  *Identifier
	Value Expression
}

func (ls *ConstStatement) statementNode() {}

// TokenLiteral returns the ConstStatement's Literal and satisfies the Node interface.
func (ls *ConstStatement) TokenLiteral() string { return ls.Token.Literal }

// String - returns a string representation of the ConstStatement and satisfies our Node interface
func (ls *ConstStatement) String() string {
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
