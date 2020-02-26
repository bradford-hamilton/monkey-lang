package lexer

import (
	"github.com/bradford-hamilton/monkey-lang/token"
)

// Lexer performs our lexical analysis/scanning
type Lexer struct {
	input        []rune
	char         rune // current char under examination
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	line         int  // line number for better error reporting, etc
}

// New creates and returns a pointer to the Lexer
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
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
	return string(l.input[position:l.position])
}

func newToken(tokenType token.Type, line int, char ...rune) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
		Line:    line,
	}
}

// Allow a-z, A-Z, _, ?
func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' || char == '?'
}

func isInteger(char rune) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) readInteger() string {
	position := l.position

	for isInteger(l.char) {
		l.readChar()
	}

	return string(l.input[position:l.position])
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		if l.char == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) skipSingleLineComment() {
	for l.char != '\n' && l.char != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipMultiLineComment() {
	endFound := false

	for !endFound {
		if l.char == 0 {
			endFound = true
		}

		if l.char == '*' && l.peek() == '/' {
			endFound = true
			l.readChar()
		}

		l.readChar()
	}

	l.skipWhitespace()
}

func (l *Lexer) peek() rune {
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
			t = token.Token{
				Type:    token.EqualEqual,
				Literal: literal,
				Line:    l.line,
			}
		} else {
			t = newToken(token.Equal, l.line, l.char)
		}
	case '+':
		if l.peek() == '+' {
			ch := l.char
			l.readChar()
			t = token.Token{
				Type:    token.PlusPlus,
				Literal: string(ch) + string(l.char),
				Line:    l.line,
			}
		} else {
			t = newToken(token.Plus, l.line, l.char)
		}
	case '-':
		if l.peek() == '-' {
			ch := l.char
			l.readChar()
			t = token.Token{
				Type:    token.MinusMinus,
				Literal: string(ch) + string(l.char),
				Line:    l.line,
			}
		} else {
			t = newToken(token.Minus, l.line, l.char)
		}
	case '!':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			literal := string(ch) + string(l.char)
			t = token.Token{
				Type:    token.BangEqual,
				Literal: literal,
				Line:    l.line,
			}
		} else {
			t = newToken(token.Bang, l.line, l.char)
		}
	case '*':
		t = newToken(token.Star, l.line, l.char)
	case '/':
		if l.peek() == '/' {
			l.skipSingleLineComment()
			return l.NextToken()
		}
		if l.peek() == '*' {
			l.skipMultiLineComment()
			return l.NextToken()
		}
		t = newToken(token.Slash, l.line, l.char)
	case '%':
		t = newToken(token.Mod, l.line, l.char)
	case '<':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			t = newToken(token.LessEqual, l.line, ch, l.char)
		} else {
			t = newToken(token.Less, l.line, l.char)
		}
	case '>':
		if l.peek() == '=' {
			ch := l.char
			l.readChar()
			t = newToken(token.GreaterEqual, l.line, ch, l.char)
		} else {
			t = newToken(token.Greater, l.line, l.char)
		}
	case '&':
		if l.peek() == '&' {
			ch := l.char
			l.readChar()
			t = newToken(token.And, l.line, ch, l.char)
		}
	case '|':
		if l.peek() == '|' {
			ch := l.char
			l.readChar()
			t = newToken(token.Or, l.line, ch, l.char)
		}
	case ',':
		t = newToken(token.Comma, l.line, l.char)
	case ':':
		t = newToken(token.Colon, l.line, l.char)
	case ';':
		t = newToken(token.Semicolon, l.line, l.char)
	case '(':
		t = newToken(token.LeftParen, l.line, l.char)
	case ')':
		t = newToken(token.RightParen, l.line, l.char)
	case '{':
		t = newToken(token.LeftBrace, l.line, l.char)
	case '}':
		t = newToken(token.RightBrace, l.line, l.char)
	case '[':
		t = newToken(token.LeftBracket, l.line, l.char)
	case ']':
		t = newToken(token.RightBracket, l.line, l.char)
	case '"':
		t.Type = token.String
		t.Literal = l.readString()
		t.Line = l.line
	case 0:
		t.Literal = ""
		t.Type = token.EOF
		t.Line = l.line
	default:
		if isLetter(l.char) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			t.Line = l.line
			return t
		} else if isInteger(l.char) {
			t.Literal = l.readInteger()
			t.Type = token.Integer
			t.Line = l.line
			return t
		} else {
			t = newToken(token.Illegal, l.line, l.char)
		}
	}

	l.readChar()
	return t
}
