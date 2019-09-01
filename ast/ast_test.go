package ast

import (
	"testing"

	"github.com/bradford-hamilton/monkey-lang/token"
)

func TestString(t *testing.T) {
	program := &RootNode{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.Let, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}
}

func TestArrayLiteral(t *testing.T) {
	arrLit := &ArrayLiteral{
		Token: token.Token{Type: token.LeftBracket, Literal: "["},
		Elements: []Expression{
			&IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "8"},
				Value: 8,
			},
			&IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "13"},
				Value: 13,
			},
		},
	}

	if arrLit.TokenLiteral() != "[" {
		t.Errorf("Wrong TokenLiteral for ArrayLiteral. Expected: '['. Got: %s", arrLit.TokenLiteral())
	}

	if arrLit.String() != "[8, 13]" {
		t.Errorf("Wrong String representation for ArrayLiteral. Expected: [8, 13]. Got: %s", arrLit.String())
	}
}

func TestBlockStatement(t *testing.T) {
	bs := &BlockStatement{
		Token: token.Token{Type: token.LeftBrace, Literal: "{"},
		Statements: []Statement{
			&ConstStatement{
				Token: token.Token{Type: token.Const, Literal: "const"},
				Name: &Identifier{
					Token: token.Token{Type: token.String, Literal: "number"},
					Value: "number",
				},
				Value: &IntegerLiteral{
					Token: token.Token{Type: token.Integer, Literal: "66"},
					Value: 66,
				},
			},
		},
	}

	if bs.TokenLiteral() != "{" {
		t.Errorf("Wrong TokenLiteral for BlockStatement. Expected: '{'. Got: %s", bs.TokenLiteral())
	}

	if bs.String() != "const number = 66;" {
		t.Errorf("Wrong String representation for BlockStatement. Expected: 'const number = 66;'. Got: %s", bs.String())
	}
}

func TestBoolean(t *testing.T) {
	b := &Boolean{
		Token: token.Token{Type: token.True, Literal: "true"},
		Value: true,
	}

	if b.TokenLiteral() != "true" {
		t.Errorf("Wrong TokenLiteral for Boolean. Expected: 'true'. Got: %s", b.TokenLiteral())
	}

	if b.String() != "true" {
		t.Errorf("Wrong String representation for Boolean. Expected: 'true'. Got: %s", b.String())
	}
}

func TestCallExpression(t *testing.T) {
	ce := &CallExpression{
		Token: token.Token{Type: token.LeftParen, Literal: "("},
		Function: &FunctionLiteral{
			Token:      token.Token{Type: token.Function, Literal: "func"},
			Parameters: []*Identifier{},
			Body:       &BlockStatement{},
			Name:       "add",
		},
		Arguments: []Expression{
			&IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "1"},
				Value: 1,
			},
			&IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "2"},
				Value: 2,
			},
		},
	}

	if ce.TokenLiteral() != "(" {
		t.Errorf("Wrong TokenLiteral for CallExpression. Expected: '('. Got: %s", ce.TokenLiteral())
	}

	if ce.String() != "func<add>() (1, 2)" {
		t.Errorf("Wrong String representation for CallExpression. Expected: 'func<add>() (1, 2)'. Got: %s", ce.String())
	}
}

func TestConstStatement(t *testing.T) {
	cs := &ConstStatement{
		Token: token.Token{Type: token.Const, Literal: "const"},
		Name: &Identifier{
			Token: token.Token{Type: token.String, Literal: "monkey"},
			Value: "monkey",
		},
		Value: &Identifier{
			Token: token.Token{Type: token.String, Literal: "lang"},
			Value: "lang",
		},
	}

	if cs.TokenLiteral() != "const" {
		t.Errorf("Wrong TokenLiteral for ConstStatement. Expected: 'const'. Got: %s", cs.TokenLiteral())
	}

	if cs.String() != "const monkey = lang;" {
		t.Errorf("Wrong String representation for ConstStatement. Expected: 'const monkey = lang;'. Got: %s", cs.String())
	}
}

func TestExpressionStatement(t *testing.T) {
	es := &ExpressionStatement{
		Token: token.Token{Type: token.Integer, Literal: "1000"},
		Expression: &IntegerLiteral{
			Token: token.Token{Type: token.Integer, Literal: "1000"},
			Value: 1000,
		},
	}

	if es.TokenLiteral() != "1000" {
		t.Errorf("Wrong TokenLiteral for ExpressionStatement. Expected: '1000'. Got: %s", es.TokenLiteral())
	}

	if es.String() != "1000" {
		t.Errorf("Wrong String representation for ExpressionStatement. Expected: '1000'. Got: %s", es.String())
	}
}

func TestFunctionLiteral(t *testing.T) {
	fl := &FunctionLiteral{
		Token: token.Token{Type: token.Function, Literal: "func"},
		Parameters: []*Identifier{
			&Identifier{
				Token: token.Token{Type: token.String, Literal: "arg1"},
				Value: "arg1",
			},
			&Identifier{
				Token: token.Token{Type: token.String, Literal: "arg2"},
				Value: "arg2",
			},
		},
		Body: &BlockStatement{
			Token: token.Token{Type: token.LeftBrace, Literal: "{"},
			Statements: []Statement{
				&ConstStatement{
					Token: token.Token{Type: token.Const, Literal: "const"},
					Name: &Identifier{
						Token: token.Token{Type: token.String, Literal: "number"},
						Value: "number",
					},
					Value: &IntegerLiteral{
						Token: token.Token{Type: token.Integer, Literal: "66"},
						Value: 66,
					},
				},
			},
		},
		Name: "add",
	}

	if fl.TokenLiteral() != "func" {
		t.Errorf("Wrong TokenLiteral for FunctionLiteral. Expected: 'func'. Got: %s", fl.TokenLiteral())
	}

	if fl.String() != "func<add>(arg1, arg2) const number = 66;" {
		t.Errorf("Wrong String representation for FunctionLiteral. Expected: 'func<add>(arg1, arg2) const number = 66;'. Got: %s", fl.String())
	}
}

func TestHashLiteral(t *testing.T) {
	hl := &HashLiteral{
		Token: token.Token{Type: token.LeftBrace, Literal: "{"},
		Pairs: map[Expression]Expression{
			&StringLiteral{
				Token: token.Token{Type: token.String, Literal: "num1"},
				Value: "num1",
			}: &IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "999"},
				Value: 999,
			},
			&StringLiteral{
				Token: token.Token{Type: token.String, Literal: "num2"},
				Value: "num2",
			}: &IntegerLiteral{
				Token: token.Token{Type: token.Integer, Literal: "11"},
				Value: 11,
			},
		},
	}

	if hl.TokenLiteral() != "{" {
		t.Errorf("Wrong TokenLiteral for HashLiteral. Expected: '{'. Got: %s", hl.TokenLiteral())
	}

	if hl.String() != "{num1:999, num2:11}" {
		t.Errorf("Wrong String representation for HashLiteral. Expected: '{num1:999, num2:11}'. Got: %s", hl.String())
	}
}

func TestIdentifier(t *testing.T) {
	ident := &Identifier{
		Token: token.Token{Type: token.String, Literal: "monkey"},
		Value: "monkey",
	}

	if ident.TokenLiteral() != "monkey" {
		t.Errorf("Wrong TokenLiteral for Identifier. Expected: 'monkey'. Got: %s", ident.TokenLiteral())
	}

	if ident.String() != "monkey" {
		t.Errorf("Wrong String representation for Identifier. Expected: 'monkey'. Got: %s", ident.String())
	}
}
