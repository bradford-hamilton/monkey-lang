package lexer

import (
	"fmt"
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

let snake_case_with_question_mark? = true;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
	}{
		{token.Let, "let", 1},
		{token.Identifier, "five", 1},
		{token.Equal, "=", 1},
		{token.Integer, "5", 1},
		{token.Semicolon, ";", 1},
		{token.Let, "let", 2},
		{token.Identifier, "ten", 2},
		{token.Equal, "=", 2},
		{token.Integer, "10", 2},
		{token.Semicolon, ";", 2},
		{token.Let, "let", 3},
		{token.Identifier, "add", 3},
		{token.Equal, "=", 3},
		{token.Function, "func", 3},
		{token.LeftParen, "(", 3},
		{token.Identifier, "x", 3},
		{token.Comma, ",", 3},
		{token.Identifier, "y", 3},
		{token.RightParen, ")", 3},
		{token.LeftBrace, "{", 3},
		{token.Identifier, "x", 4},
		{token.Plus, "+", 4},
		{token.Identifier, "y", 4},
		{token.Semicolon, ";", 4},
		{token.RightBrace, "}", 5},
		{token.Semicolon, ";", 5},
		{token.Let, "let", 6},
		{token.Identifier, "result", 6},
		{token.Equal, "=", 6},
		{token.Identifier, "add", 6},
		{token.LeftParen, "(", 6},
		{token.Identifier, "five", 6},
		{token.Comma, ",", 6},
		{token.Identifier, "ten", 6},
		{token.RightParen, ")", 6},
		{token.Semicolon, ";", 6},
		{token.Bang, "!", 8},
		{token.Minus, "-", 8},
		{token.Star, "*", 8},
		{token.Slash, "/", 8},
		{token.Integer, "5", 8},
		{token.Semicolon, ";", 8},
		{token.Integer, "5", 9},
		{token.Less, "<", 9},
		{token.Integer, "10", 9},
		{token.Greater, ">", 9},
		{token.Integer, "5", 9},
		{token.Semicolon, ";", 9},
		{token.If, "if", 11},
		{token.LeftParen, "(", 11},
		{token.Integer, "5", 11},
		{token.Less, "<", 11},
		{token.Integer, "10", 11},
		{token.RightParen, ")", 11},
		{token.LeftBrace, "{", 11},
		{token.Return, "return", 12},
		{token.True, "true", 12},
		{token.Semicolon, ";", 12},
		{token.RightBrace, "}", 13},
		{token.Else, "else", 13},
		{token.LeftBrace, "{", 13},
		{token.Return, "return", 14},
		{token.False, "false", 14},
		{token.Semicolon, ";", 14},
		{token.RightBrace, "}", 15},
		{token.Integer, "10", 17},
		{token.EqualEqual, "==", 17},
		{token.Integer, "10", 17},
		{token.Semicolon, ";", 17},
		{token.Integer, "10", 18},
		{token.BangEqual, "!=", 18},
		{token.Integer, "9", 18},
		{token.Semicolon, ";", 18},
		{token.String, "foobar", 20},
		{token.String, "foo bar", 21},
		{token.LeftBracket, "[", 25},
		{token.Integer, "1", 25},
		{token.Comma, ",", 25},
		{token.Integer, "2", 25},
		{token.RightBracket, "]", 25},
		{token.Semicolon, ";", 25},
		{token.LeftBrace, "{", 27},
		{token.String, "foo", 27},
		{token.Colon, ":", 27},
		{token.String, "bar", 27},
		{token.RightBrace, "}", 27},
		{token.True, "true", 29},
		{token.And, "&&", 29},
		{token.False, "false", 29},
		{token.Semicolon, ";", 29},
		{token.True, "true", 30},
		{token.Or, "||", 30},
		{token.False, "false", 30},
		{token.Semicolon, ";", 30},
		{token.Integer, "10", 33},
		{token.Integer, "10", 35},
		{token.Const, "const", 37},
		{token.Identifier, "cantChangeMe", 37},
		{token.Equal, "=", 37},
		{token.String, "neato", 37},
		{token.Semicolon, ";", 37},
		{token.Integer, "10", 39},
		{token.Mod, "%", 39},
		{token.Integer, "3", 39},
		{token.Semicolon, ";", 39},
		{token.Identifier, "five", 41},
		{token.PlusPlus, "++", 41},
		{token.Identifier, "five", 42},
		{token.MinusMinus, "--", 42},
		{token.Integer, "5", 44},
		{token.GreaterEqual, ">=", 44},
		{token.Integer, "5", 44},
		{token.Semicolon, ";", 44},
		{token.Integer, "5", 45},
		{token.LessEqual, "<=", 45},
		{token.Integer, "5", 45},
		{token.Semicolon, ";", 45},
		{token.Let, "let", 49},
		{token.Identifier, "snake_case_with_question_mark?", 49},
		{token.Equal, "=", 49},
		{token.True, "true", 49},
		{token.Semicolon, ";", 49},
		{token.EOF, "", 50},
	}

	l := New(input)

	for i, tt := range tests {
		token := l.NextToken()
		fmt.Printf("Line: %d: Expected: %d", token.Line, tt.expectedLine)
		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. Expected: %q, Got: %q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. Expected: %q, Got: %q", i, tt.expectedLiteral, token.Literal)
		}

		if token.Line != tt.expectedLine {
			t.Fatalf("tests[%d] - line wrong. Expected: %q, Got: %q", i, tt.expectedLine, token.Line)
		}
	}
}
