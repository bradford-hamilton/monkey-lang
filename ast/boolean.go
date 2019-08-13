package ast

import "github.com/bradford-hamilton/monkey-lang/token"

// Boolean - holds the token and it's value (a boolean)
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns the Boolean's Literal and satisfies the Node interface.
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String - returns a string representation of the boolean and satisfies our Node interface
func (b *Boolean) String() string { return b.Token.Literal }
