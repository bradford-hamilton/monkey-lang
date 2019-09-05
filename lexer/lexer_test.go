package lexer

import (
	"testing"

	"github.com/bradford-hamilton/monkey-lang/token"
)

func TestNextToken(t *testing.T) {
	input := `
let five = 5;
let ten = 10;
let add = func(x, y) {
	x + y;
};
let result = add(five, ten);

!-*/5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;

"foobar"
"foo bar"

/* multiline comment */

[1, 2];

{"foo": "bar"}

true && false;
true || false;

// This is a comment above the number 10
10

10 // This is a comment to the right of 10

const cantChangeMe = "neato";

10 % 3;

five++
five--

5 >= 5;
5 <= 5;

/*
	multiline comment
*/
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Equal, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Equal, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Equal, "="},
		{token.Function, "func"},
		{token.LeftParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Equal, "="},
		{token.Identifier, "add"},
		{token.LeftParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RightParen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Star, "*"},
		{token.Slash, "/"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.Less, "<"},
		{token.Integer, "10"},
		{token.Greater, ">"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.LeftParen, "("},
		{token.Integer, "5"},
		{token.Less, "<"},
		{token.Integer, "10"},
		{token.RightParen, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Integer, "10"},
		{token.EqualEqual, "=="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.BangEqual, "!="},
		{token.Integer, "9"},
		{token.Semicolon, ";"},
		{token.String, "foobar"},
		{token.String, "foo bar"},
		{token.LeftBracket, "["},
		{token.Integer, "1"},
		{token.Comma, ","},
		{token.Integer, "2"},
		{token.RightBracket, "]"},
		{token.Semicolon, ";"},
		{token.LeftBrace, "{"},
		{token.String, "foo"},
		{token.Colon, ":"},
		{token.String, "bar"},
		{token.RightBrace, "}"},
		{token.True, "true"},
		{token.And, "&&"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.True, "true"},
		{token.Or, "||"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.Integer, "10"},
		{token.Const, "const"},
		{token.Identifier, "cantChangeMe"},
		{token.Equal, "="},
		{token.String, "neato"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.Mod, "%"},
		{token.Integer, "3"},
		{token.Semicolon, ";"},
		{token.Identifier, "five"},
		{token.PlusPlus, "++"},
		{token.Identifier, "five"},
		{token.MinusMinus, "--"},
		{token.Integer, "5"},
		{token.GreaterEqual, ">="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.LessEqual, "<="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected: %q, Got: %q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected: %q, Got: %q", i, tt.expectedLiteral, token.Literal)
		}
	}
}
