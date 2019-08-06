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

// PrefixExpression - holds the token, a string version of the operator, and the expression to the right of it
type PrefixExpression struct {
	Token    token.Token // The prefix token (! or -)
	Operator string      // string (either "!" or "-")
	Right    Expression  // The expression to the right of the operator
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the PrefixExpression's Literal and satisfies the Node interface.
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String - returns a string representation of the operator followed by it's expression to the
// right (-5) and satisfies our Node interface
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression - holds the token, the expression to the left of it, a string version of
// the operator, and the expression to the right of it
type InfixExpression struct {
	Token    token.Token // The operator token (+, -, *, etc)
	Left     Expression
	Operator string // string (examples: "+", "-", "*", etc)
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the InfixExpression's Literal and satisfies the Node interface.
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String - returns a string representation of the left side expression, the operator, and
// the right side expression (5 * 5) and satisfies our Node interface
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

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

// IfExpression - holds the token, the condition expression and the consequence & alternative
// block statements. Structure: if (<condition>) <consequence> else <alternative>
type IfExpression struct {
	Token       token.Token // The IF token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the IfExpression's Literal and satisfies the Node interface.
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String - returns a string representation of the IfExpression with the consequence and
// also the alteritive if it is not nil. Satisfies our Node interface
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}
