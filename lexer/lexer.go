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

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isInteger(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readInteger() string {
	position := l.position

	for isInteger(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken switches through the lexer's current char and creates a new token.
// It then it calls readChar() to advance the lexer and it returns the token
func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.EQUAL_EQUAL, Literal: literal}
		} else {
			t = newToken(token.EQUAL, l.char)
		}
	case '+':
		t = newToken(token.PLUS, l.char)
	case '-':
		t = newToken(token.MINUS, l.char)
	case '!':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.BANG_EQUAL, Literal: literal}
		} else {
			t = newToken(token.BANG, l.char)
		}
	case '*':
		t = newToken(token.STAR, l.char)
	case '/':
		t = newToken(token.SLASH, l.char)
	case '<':
		t = newToken(token.LESS_EQUAL, l.char)
	case '>':
		t = newToken(token.GREATER_EQUAL, l.char)
	case ',':
		t = newToken(token.COMMA, l.char)
	case ';':
		t = newToken(token.SEMICOLON, l.char)
	case '(':
		t = newToken(token.LEFT_PAREN, l.char)
	case ')':
		t = newToken(token.RIGHT_PAREN, l.char)
	case '{':
		t = newToken(token.LEFT_BRACE, l.char)
	case '}':
		t = newToken(token.RIGHT_BRACE, l.char)
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		} else if isInteger(l.char) {
			t.Literal = l.readInteger()
			t.Type = token.INTEGER
			return t
		} else {
			t = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return t
}
