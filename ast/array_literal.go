package ast

import (
	"bytes"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// ArrayLiteral - holds the token: '[' and an array of expressions (Elements)
type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode() {}

// TokenLiteral returns the ArrayLiteral's Literal (the string) and satisfies the Node interface.
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }

// String - loops through the array literal's elements and prints them. Satisfies our Node interface
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
