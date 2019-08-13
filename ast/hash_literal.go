package ast

import (
	"bytes"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// HashLiteral - holds the '{' token and the pairs in the hash
type HashLiteral struct {
	Token token.Token // The '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

// TokenLiteral returns Token Literal and satisfies the Node interface.
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }

// String - iterates over the hl's pairs and prints string representation
// of the hash. Satisfies our Node interface
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
