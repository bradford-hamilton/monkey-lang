package lexer

import "github.com/bradford-hamilton/monkey-lang/token"

// Lexer performs our lexical analysis/scanning
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

// New creates and returns a pointer to the Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// End of input (haven't read anything yet or EOF)
		// 0 is ASCII code for "NUL" character
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken switches through the lexer's current char and creates a new token.
// It then it calls readChar() to advance the lexer and it returns the token
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.char {
	case '=':
		t = newToken(token.ASSIGN, l.char)
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case '(':
		t = newToken(token.LPAREN, l.char)
	case ')':
		t = newToken(token.RPAREN, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case '+':
		t = newToken(token.PLUS, l.char)
	case '{':
		t = newToken(token.LBRACE, l.char)
	case '}':
		t = newToken(token.RBRACE, l.char)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	}

	l.readChar()

	return t
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}
