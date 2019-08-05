package ast

import (
	"bytes"

	"github.com/bradford-hamilton/monkey-lang/token"
)

// Node - nodes in our ast will provide a TokenLiteral method for debugging
type Node interface {
	TokenLiteral() string
	String() string
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

// String returns a buffer containing the programs Statements as strings.
func (p *RootNode) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String - returns a string representation of the LetStatement and satisfies our Node interface
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier - holds IDENTIFIER token and it's value (add, foobar, x, y, ...)
type Identifier struct {
	Token token.Token // The token.IDENTIFIER token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the Identifier's Literal and satisfies the Node interface.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String - returns a string representation of the Identifier and satisfies our Node interface
func (i *Identifier) String() string { return i.Value }

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

// ExpressionStatement - holds the first token of the expression and the expression
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the ExpressionStatement's Literal and satisfies the Node interface.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String - returns a string representation of the ExpressionStatement and satisfies our Node interface
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

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
