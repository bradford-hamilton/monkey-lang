package parser

import (
	"fmt"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/lexer"
	"github.com/bradford-hamilton/monkey-lang/token"
)

// Parser holds a Lexer, the currentToken, and the peekToken
type Parser struct {
	lexer        *lexer.Lexer
	errors       []string
	currentToken token.Token
	peekToken    token.Token
}

// New takes a Lexer, creates a Parser with that Lexer, sets the current and
// peek tokens, and returns the Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	// Read two tokens, so currentToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// Errors is simply a helper function that returns the parser's errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("Expected next token to be %s, got: %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) currentTokenTypeIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenTypeIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekType(t token.TokenType) bool {
	if p.peekTokenTypeIs(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeekType(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeekType(token.EQUAL) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon.
	for !p.currentTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// ParseProgram parses tokens and creates an AST. It returns the RootNode
// which holds a slice of Statements (and in turn, the rest of the tree)
func (p *Parser) ParseProgram() *ast.RootNode {
	rootNode := &ast.RootNode{}
	rootNode.Statements = []ast.Statement{}

	for !p.currentTokenTypeIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			rootNode.Statements = append(rootNode.Statements, stmt)
		}
		p.nextToken()
	}

	return rootNode
}
