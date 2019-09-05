package lexer

import (
	"github.com/bradford-hamilton/monkey-lang/token"
)

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

func newToken(tokenType token.TokenType, char ...byte) token.Token {
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

func (l *Lexer) skipSingleLineComment() {
	for l.char != '\n' && l.char != 0 {
		l.readChar()
	}
	l.skipWhitespace()
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
			t = token.Token{Type: token.EqualEqual, Literal: literal}
		} else {
			t = newToken(token.Equal, l.char)
		}
	case '+':
		if l.peek() == '+' {
			ch := l.char
			l.readChar()
			t = token.Token{
				Type:    token.PlusPlus,
				Literal: string(ch) + string(l.char),
			}
		} else {
			t = newToken(token.Plus, l.char)
		}
	case '-':
		if l.peek() == '-' {
			ch := l.char
			l.readChar()
			t = token.Token{
				Type:    token.MinusMinus,
				Literal: string(ch) + string(l.char),
			}
		} else {
			t = newToken(token.Minus, l.char)
		}
	case '!':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{Type: token.BangEqual, Literal: literal}
		} else {
			t = newToken(token.Bang, l.char)
		}
	case '*':
		t = newToken(token.Star, l.char)
	case '/':
		if l.peek() == '/' {
			l.skipSingleLineComment()
			return l.NextToken()
		}
		t = newToken(token.Slash, l.char)
	case '%':
		t = newToken(token.Mod, l.char)
	case '<':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			t = newToken(token.LessEqual, ch, l.char)
		} else {
			t = newToken(token.Less, l.char)
		}
	case '>':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			t = newToken(token.GreaterEqual, ch, l.char)
		} else {
			t = newToken(token.Greater, l.char)
		}
	case '&':
		if l.peek() == '&' {
			ch := l.char
			l.readChar()
			t = newToken(token.And, ch, l.char)
		}
	case '|':
		if l.peek() == '|' {
			ch := l.char
			l.readChar()
			t = newToken(token.Or, ch, l.char)
		}
	case ',':
		t = newToken(token.Comma, l.char)
	case ':':
		t = newToken(token.Colon, l.char)
	case ';':
		t = newToken(token.Semicolon, l.char)
	case '(':
		t = newToken(token.LeftParen, l.char)
	case ')':
		t = newToken(token.RightParen, l.char)
	case '{':
		t = newToken(token.LeftBrace, l.char)
	case '}':
		t = newToken(token.RightBrace, l.char)
	case '[':
		t = newToken(token.LeftBracket, l.char)
	case ']':
		t = newToken(token.RightBracket, l.char)
	case '"':
		t.Type = token.String
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
			t.Type = token.Integer
			return t
		} else {
			t = newToken(token.Illegal, l.char)
		}
	}

	l.readChar()
	return t
}
