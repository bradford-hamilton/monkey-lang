package token

const (
	ILLEGAL = "ILLEGAL" // Token/character we don't know about
	EOF     = "EOF"     // End of file

	// Identifiers & literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	NUMBER     = "NUMBER"     // 1, 2, 3, 4, 5

	// Operators
	EQUAL         = "="
	PLUS          = "+"
	MINUS         = "-"
	BANG          = "!"
	STAR          = "*"
	SLASH         = "/"
	EQUAL_EQUAL   = "=="
	LESS_EQUAL    = "<"
	GREATER_EQUAL = ">"
	BANG_EQUAL    = "!="

	// Delimiters
	COMMA       = ","
	SEMICOLON   = ";"
	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier checks our keywords map for the scanned keyword. If it finds one, then
// the keywords TokenType is returned. If not the user defined IDENTIFIER is returned
func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}
