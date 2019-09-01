package ast

import (
	"testing"

	"github.com/bradford-hamilton/monkey-lang/token"
)

func TestRootNode(t *testing.T) {
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
	emptyProgram := &RootNode{}

	if program.TokenLiteral() != "let" {
		t.Errorf("program.TokenLiteral() wrong. Expected: 'let'. Got: %q", program.TokenLiteral())
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}

	if emptyProgram.TokenLiteral() != "" {
		t.Errorf("emptyProgram.String() wrong. Expected \"\" Got: %q", program.String())
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
	blankExpr := &ExpressionStatement{}

	if es.TokenLiteral() != "1000" {
		t.Errorf("Wrong TokenLiteral for ExpressionStatement. Expected: '1000'. Got: %s", es.TokenLiteral())
	}

	if es.String() != "1000" {
		t.Errorf("Wrong String representation for ExpressionStatement. Expected: '1000'. Got: %s", es.String())
	}

	if blankExpr.String() != "" {
		t.Errorf("Wrong String representation for empty ExpressionStatement. Expected: empty string \" \". Got: %s", es.String())
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

func TestIfExpression(t *testing.T) {
	ie := &IfExpression{
		Token: token.Token{Type: token.If, Literal: "if"},
		Condition: &Boolean{
			Token: token.Token{Type: token.True, Literal: "true"},
			Value: true,
		},
		Consequence: &BlockStatement{
			Token: token.Token{Type: token.LeftBrace, Literal: "{"},
			Statements: []Statement{
				&ConstStatement{
					Token: token.Token{Type: token.Const, Literal: "const"},
					Name: &Identifier{
						Token: token.Token{Type: token.String, Literal: "whenTrue"},
						Value: "whenTrue",
					},
					Value: &Boolean{
						Token: token.Token{Type: token.Integer, Literal: "true"},
						Value: true,
					},
				},
			},
		},
		Alternative: &BlockStatement{
			Token: token.Token{Type: token.LeftBrace, Literal: "{"},
			Statements: []Statement{
				&ConstStatement{
					Token: token.Token{Type: token.Const, Literal: "const"},
					Name: &Identifier{
						Token: token.Token{Type: token.String, Literal: "whenFalse"},
						Value: "whenFalse",
					},
					Value: &Boolean{
						Token: token.Token{Type: token.Integer, Literal: "false"},
						Value: false,
					},
				},
			},
		},
	}

	if ie.TokenLiteral() != "if" {
		t.Errorf("Wrong TokenLiteral for IfExpression. Expected: 'if'. Got: %s", ie.TokenLiteral())
	}

	if ie.String() != "if true const whenTrue = true; else const whenFalse = false;" {
		t.Errorf("Wrong String representation for IfExpression. Expected: 'if true const whenTrue = true; else const whenFalse = false;'. Got: %s", ie.String())
	}
}

func TestIndexExpression(t *testing.T) {
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

	ie := &IndexExpression{
		Token: token.Token{Type: token.LeftBracket, Literal: "["},
		Left:  arrLit,
		Index: &IntegerLiteral{
			Token: token.Token{Type: token.Integer, Literal: "0"},
			Value: 0,
		},
	}

	if ie.TokenLiteral() != "[" {
		t.Errorf("Wrong TokenLiteral for IndexExpression. Expected: '['. Got: %s", ie.TokenLiteral())
	}

	if ie.String() != "([8, 13][0])" {
		t.Errorf("Wrong String representation for IndexExpression. Expected: '([8, 13][0])'. Got: %s", ie.String())
	}
}

func TestInfixExpression(t *testing.T) {
	ie := &InfixExpression{
		Token: token.Token{Type: token.Plus, Literal: "+"},
		Left: &IntegerLiteral{
			Token: token.Token{Type: token.Integer, Literal: "5"},
			Value: 5,
		},
		Operator: "+",
		Right: &IntegerLiteral{
			Token: token.Token{Type: token.Integer, Literal: "10"},
			Value: 10,
		},
	}

	if ie.TokenLiteral() != "+" {
		t.Errorf("Wrong TokenLiteral for InfixExpression. Expected: '+'. Got: %s", ie.TokenLiteral())
	}

	if ie.String() != "(5 + 10)" {
		t.Errorf("Wrong String representation for InfixExpression. Expected: '(5 + 10)'. Got: %s", ie.String())
	}
}

func TestIntegerLiteral(t *testing.T) {
	il := &IntegerLiteral{
		Token: token.Token{Type: token.Integer, Literal: "10"},
		Value: 10,
	}

	if il.TokenLiteral() != "10" {
		t.Errorf("Wrong TokenLiteral for IntegerLiteral. Expected: '10'. Got: %s", il.TokenLiteral())
	}

	if il.String() != "10" {
		t.Errorf("Wrong String representation for IntegerLiteral. Expected: '10'. Got: %s", il.String())
	}
}

func TestLetStatement(t *testing.T) {
	cs := &LetStatement{
		Token: token.Token{Type: token.Let, Literal: "let"},
		Name: &Identifier{
			Token: token.Token{Type: token.String, Literal: "monkey"},
			Value: "monkey",
		},
		Value: &Identifier{
			Token: token.Token{Type: token.String, Literal: "lang"},
			Value: "lang",
		},
	}

	if cs.TokenLiteral() != "let" {
		t.Errorf("Wrong TokenLiteral for LetStatement. Expected: 'let'. Got: %s", cs.TokenLiteral())
	}

	if cs.String() != "let monkey = lang;" {
		t.Errorf("Wrong String representation for LetStatement. Expected: 'let monkey = lang;'. Got: %s", cs.String())
	}
}

func TestPrefixExpression(t *testing.T) {
	pe := &PrefixExpression{
		Token: token.Token{Type: token.Minus, Literal: "-"},
		Right: &IntegerLiteral{
			Token: token.Token{Type: token.Integer, Literal: "5"},
			Value: 5,
		},
		Operator: "-",
	}

	if pe.TokenLiteral() != "-" {
		t.Errorf("Wrong TokenLiteral for PrefixExpression. Expected: '-'. Got: %s", pe.TokenLiteral())
	}

	if pe.String() != "(-5)" {
		t.Errorf("Wrong String representation for PrefixExpression. Expected: '(-5)'. Got: %s", pe.String())
	}
}

func TestReturnStatement(t *testing.T) {
	rs := &ReturnStatement{
		Token: token.Token{Type: token.Return, Literal: "return"},
		ReturnValue: &StringLiteral{
			Token: token.Token{Type: token.String, Literal: "monkeys"},
			Value: "monkeys",
		},
	}

	if rs.TokenLiteral() != "return" {
		t.Errorf("Wrong TokenLiteral for ReturnStatement. Expected: 'return'. Got: %s", rs.TokenLiteral())
	}

	if rs.String() != "return monkeys;" {
		t.Errorf("Wrong String representation for ReturnStatement. Expected: 'return monkeys;'. Got: %s", rs.String())
	}
}

func TestStringLiteral(t *testing.T) {
	sl := &StringLiteral{
		Token: token.Token{Type: token.String, Literal: "this string is so literal"},
		Value: "this string is so literal",
	}

	if sl.TokenLiteral() != "this string is so literal" {
		t.Errorf("Wrong TokenLiteral for StringLiteral. Expected: 'this string is so literal'. Got: %s", sl.TokenLiteral())
	}

	if sl.String() != "this string is so literal" {
		t.Errorf("Wrong String representation for StringLiteral. Expected: 'this string is so literal'. Got: %s", sl.String())
	}
}
