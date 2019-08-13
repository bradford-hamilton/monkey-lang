package ast

import "github.com/bradford-hamilton/monkey-lang/token"

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
