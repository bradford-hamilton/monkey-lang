package ast

import "github.com/bradford-hamilton/monkey-lang/token"

// StringLiteral - holds the token and it's value (string)
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

// TokenLiteral returns the StringLiteral's Literal (the string) and satisfies the Node interface.
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

// String - returns a string representation of the StringLiteral and satisfies our Node interface
func (sl *StringLiteral) String() string { return sl.Token.Literal }
