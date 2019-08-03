package ast

import "github.com/bradford-hamilton/monkey-lang/token"

// Node - nodes in our ast will provide a TokenLiteral method for debugging
type Node interface {
	TokenLiteral() string
}

// Statement - must provide statementNode and TokenLiteral method. Statements do not produce values.
type Statement interface {
	Node
	statementNode()
}

// Expression - must provide expressionNode and TokenLiteral method. Expressions produce values.
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

// LetStatement - Name holds the identifier of the binding and Value for the expression
// that produces the value.
type LetStatement struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the LetStatement's Literal and satisfies the Node interface.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier -
type Identifier struct {
	Token token.Token // The token.IDENTIFIER token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the Identifier's Literal and satisfies the Node interface.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // The 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the ReturnStatement's Literal and satisfies the Node interface.
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
