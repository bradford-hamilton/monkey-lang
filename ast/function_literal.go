package ast

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// FunctionLiteral - holds the token, the function params (a slice of *Identifier), and
// the function Body (*BlockStatement). Structure: func <parameters> <block statement>
type FunctionLiteral struct {
	Token      token.Token // The 'func' token
	Parameters []*Identifier
	Body       *BlockStatement
	Name       string
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns the FunctionLiteral's Literal and satisfies the Node interface.
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }

// String - returns a string representation of the FunctionLiteral. Prints it's token,
// params, and body. Satisfies our Node interface
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	if fl.Name != "" {
		out.WriteString(fmt.Sprintf("<%s>", fl.Name))
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}
