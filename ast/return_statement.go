package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// ReturnStatement - pretty self explanatory, holds RETURN token and return value
type ReturnStatement struct {
	Token       token.Token // The 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the ReturnStatement's Literal and satisfies the Node interface.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// String - returns a string representation of the ReturnStatement and satisfies our Node interface
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
