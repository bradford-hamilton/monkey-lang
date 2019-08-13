package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// BlockStatement - holds the token "{", and a slice of statements
type BlockStatement struct {
	Token      token.Token // The { token
	Statements []Statement
}

func (bs *BlockStatement) expressionNode() {}

// TokenLiteral returns the BlockStatement's Literal and satisfies the Node interface.
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String - returns a string representation of the statements inside the block
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
