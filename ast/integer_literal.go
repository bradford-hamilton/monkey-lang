package ast

import "github.com/bradford-hamilton/monkey-lang/token"

// IntegerLiteral - holds the token and it's value (int64)
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the IntegerLiteral's Literal and satisfies the Node interface.
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// String - returns a string representation of the IntegerLiteral and satisfies our Node interface
func (il *IntegerLiteral) String() string { return il.Token.Literal }
