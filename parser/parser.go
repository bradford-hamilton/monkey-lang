package parser

import (
	"fmt"
	"strconv"

	"github.com/bradford-hamilton/monkey-lang/ast"
	"github.com/bradford-hamilton/monkey-lang/lexer"
	"github.com/bradford-hamilton/monkey-lang/token"
)

// Prefix expressions are often referred to as unary expressions: -x
// Infix expressions are often referred to as binary expressions: 5 * 5

// Define operator precedence constants
const (
	LOWEST      = iota + 1
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFunction(x)
	INDEX       // array[index]
)

// Define operator precedence table
var precedences = map[token.TokenType]int{
	token.EQUAL_EQUAL:   EQUALS,
	token.BANG_EQUAL:    EQUALS,
	token.LESS_EQUAL:    LESSGREATER,
	token.GREATER_EQUAL: LESSGREATER,
	token.PLUS:          SUM,
	token.MINUS:         SUM,
	token.SLASH:         PRODUCT,
	token.STAR:          PRODUCT,
	token.LEFT_PAREN:    CALL,
	token.LEFT_BRACKET:  INDEX,
}

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

// Parser holds a Lexer, the currentToken, and the peekToken
type Parser struct {
	lexer  *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken    token.Token

	prefixParseFuncs map[token.TokenType]prefixParseFunc
	infixParseFuncs  map[token.TokenType]infixParseFunc
}

// New takes a Lexer, creates a Parser with that Lexer, sets the current and
// peek tokens, and returns the Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.prefixParseFuncs = make(map[token.TokenType]prefixParseFunc)
	p.registerPrefix(token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LEFT_PAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.LEFT_BRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LEFT_BRACE, p.parseHashLiteral)

	p.infixParseFuncs = make(map[token.TokenType]infixParseFunc)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.STAR, p.parseInfixExpression)
	p.registerInfix(token.EQUAL_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.BANG_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.LESS_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.GREATER_EQUAL, p.parseInfixExpression)
	p.registerInfix(token.LEFT_PAREN, p.parseCallExpression)
	p.registerInfix(token.LEFT_BRACKET, p.parseIndexExpr)

	// Read two tokens, so currentToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
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

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
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

func (p *Parser) peekTokenPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currenTokenPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
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

	p.nextToken()

	stmt.Value = p.parseExpr(LOWEST)

	if fl, ok := stmt.Value.(*ast.FunctionLiteral); ok {
		fl.Name = stmt.Name.Value
	}

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpr(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExprStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpr(LOWEST)

	if p.peekTokenTypeIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpr(precedence int) ast.Expression {
	prefix := p.prefixParseFuncs[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFuncError(p.currentToken.Type)
		return nil
	}
	leftExpr := prefix()

	for !p.peekTokenTypeIs(token.SEMICOLON) && precedence < p.peekTokenPrecedence() {
		infix := p.infixParseFuncs[p.peekToken.Type]
		if infix == nil {
			return leftExpr
		}

		p.nextToken()
		leftExpr = infix(leftExpr)
	}

	return leftExpr
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("Could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.currentToken,
		Value: p.currentTokenTypeIs(token.TRUE),
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	expr := p.parseExpr(LOWEST)

	if !p.expectPeekType(token.RIGHT_PAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenTypeIs(token.RIGHT_BRACE) && !p.currentTokenTypeIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Token: p.currentToken}

	if !p.expectPeekType(token.LEFT_PAREN) {
		return nil
	}

	p.nextToken()
	expr.Condition = p.parseExpr(LOWEST)

	if !p.expectPeekType(token.RIGHT_PAREN) {
		return nil
	}

	if !p.expectPeekType(token.LEFT_BRACE) {
		return nil
	}

	expr.Consequence = p.parseBlockStatement()

	if p.peekTokenTypeIs(token.ELSE) {
		p.nextToken()

		if !p.expectPeekType(token.LEFT_BRACE) {
			return nil
		}

		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenTypeIs(token.RIGHT_PAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{
			Token: p.currentToken,
			Value: p.currentToken.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeekType(token.RIGHT_PAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeekType(token.LEFT_PAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeekType(token.LEFT_BRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.currentToken}
	array.Elements = p.parseExprList(token.RIGHT_BRACKET)

	return array
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.currentToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenTypeIs(token.RIGHT_BRACE) {
		p.nextToken()
		key := p.parseExpr(LOWEST)

		if !p.expectPeekType(token.COLON) {
			return nil
		}

		p.nextToken()

		value := p.parseExpr(LOWEST)
		hash.Pairs[key] = value

		if !p.peekTokenTypeIs(token.RIGHT_BRACE) && !p.expectPeekType(token.COMMA) {
			return nil
		}
	}

	if !p.expectPeekType(token.RIGHT_BRACE) {
		return nil
	}

	return hash
}

func (p *Parser) parseIndexExpr(left ast.Expression) ast.Expression {
	expr := &ast.IndexExpression{Token: p.currentToken, Left: left}
	p.nextToken()
	expr.Index = p.parseExpr(LOWEST)

	if !p.expectPeekType(token.RIGHT_BRACKET) {
		return nil
	}

	return expr
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenTypeIs(token.RIGHT_PAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpr(LOWEST))

	for p.peekTokenTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpr(LOWEST))
	}

	if !p.expectPeekType(token.RIGHT_PAREN) {
		return nil
	}

	return args
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expr := &ast.CallExpression{
		Token:    p.currentToken,
		Function: function,
	}
	expr.Arguments = p.parseExprList(token.RIGHT_PAREN)

	return expr
}

func (p *Parser) parseExprList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenTypeIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpr(LOWEST))

	for p.peekTokenTypeIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpr(LOWEST))
	}

	if !p.expectPeekType(end) {
		return nil
	}

	return list
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()
	expr.Right = p.parseExpr(PREFIX)

	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}

	precedence := p.currenTokenPrecedence()
	p.nextToken()
	expr.Right = p.parseExpr(precedence)

	return expr
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFunc) {
	p.prefixParseFuncs[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFunc) {
	p.infixParseFuncs[tokenType] = fn
}

func (p *Parser) noPrefixParseFuncError(t token.TokenType) {
	msg := fmt.Sprintf("No prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}
