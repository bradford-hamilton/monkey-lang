package ast

import (
	"bytes"
)

// Node - nodes in our ast will provide a TokenLiteral method for debugging
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement - must provide statementNode, TokenLiteral, and String methods. Statements do not produce values.
type Statement interface {
	Node
	statementNode()
}

// Expression - must provide expressionNode, TokenLiteral, and String methods. Expressions produce values.
type Expression interface {
	Node
	expressionNode()
}

// RootNode of every AST our parser produces.
type RootNode struct {
	Statements []Statement
}

// TokenLiteral returns the RootNode's Literal and satisfies the Node interface.
func (p *RootNode) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// String returns a buffer containing the programs Statements as strings.
func (p *RootNode) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
