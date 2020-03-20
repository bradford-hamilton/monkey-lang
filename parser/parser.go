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
	Lowest      = iota + 1
	Equals      // =
	Logical     // && and ||
	LessGreater // > or <
	Sum         // +
	Product     // *
	Mod         // %
	Prefix      // -x or !x
	Call        // myFunction(x)
	Index       // array[index], hash[key]
)

// Define operator precedence table
var precedences = map[token.Type]int{
	token.EqualEqual:   Equals,
	token.BangEqual:    Equals,
	token.Less:         LessGreater,
	token.Greater:      LessGreater,
	token.LessEqual:    LessGreater,
	token.GreaterEqual: LessGreater,
	token.Plus:         Sum,
	token.Minus:        Sum,
	token.Slash:        Product,
	token.Star:         Product,
	token.Mod:          Mod,
	token.And:          Logical,
	token.Or:           Logical,
	token.LeftParen:    Call,
	token.LeftBracket:  Index,
}

type (
	prefixParseFunc  func() ast.Expression
	infixParseFunc   func(ast.Expression) ast.Expression
	postfixParseFunc func() ast.Expression
)

// Parser holds a Lexer, its errors, the currentToken, peekToken (next token), and
// prevToken (used for ++ and --), as well as the prefix/infix/postfix functions
type Parser struct {
	lexer  *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken    token.Token
	prevToken    token.Token

	prefixParseFuncs  map[token.Type]prefixParseFunc
	infixParseFuncs   map[token.Type]infixParseFunc
	postfixParseFuncs map[token.Type]postfixParseFunc
}

// New takes a Lexer, creates a Parser with that Lexer, sets the current and
// peek tokens, and returns the Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:             l,
		errors:            []string{},
		prefixParseFuncs:  make(map[token.Type]prefixParseFunc),
		infixParseFuncs:   make(map[token.Type]infixParseFunc),
		postfixParseFuncs: make(map[token.Type]postfixParseFunc),
	}

	// Register all of our prefix parse funcs
	p.registerPrefix(token.Identifier, p.parseIdentifier)
	p.registerPrefix(token.Integer, p.parseIntegerLiteral)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)
	p.registerPrefix(token.True, p.parseBoolean)
	p.registerPrefix(token.False, p.parseBoolean)
	p.registerPrefix(token.LeftParen, p.parseGroupedExpression)
	p.registerPrefix(token.If, p.parseIfExpression)
	p.registerPrefix(token.Function, p.parseFunctionLiteral)
	p.registerPrefix(token.String, p.parseStringLiteral)
	p.registerPrefix(token.LeftBracket, p.parseArrayLiteral)
	p.registerPrefix(token.LeftBrace, p.parseHashLiteral)

	// Register all of our infix parse funcs
	p.registerInfix(token.Plus, p.parseInfixExpression)
	p.registerInfix(token.Minus, p.parseInfixExpression)
	p.registerInfix(token.Slash, p.parseInfixExpression)
	p.registerInfix(token.Star, p.parseInfixExpression)
	p.registerInfix(token.Mod, p.parseInfixExpression)
	p.registerInfix(token.EqualEqual, p.parseInfixExpression)
	p.registerInfix(token.BangEqual, p.parseInfixExpression)
	p.registerInfix(token.Less, p.parseInfixExpression)
	p.registerInfix(token.Greater, p.parseInfixExpression)
	p.registerInfix(token.LessEqual, p.parseInfixExpression)
	p.registerInfix(token.GreaterEqual, p.parseInfixExpression)
	p.registerInfix(token.LeftParen, p.parseCallExpression)
	p.registerInfix(token.LeftBracket, p.parseIndexExpr)
	p.registerInfix(token.And, p.parseInfixExpression)
	p.registerInfix(token.Or, p.parseInfixExpression)

	// Register all of our postfix parse funcs
	p.registerPostfix(token.PlusPlus, p.parsePostfixExpression)
	p.registerPostfix(token.MinusMinus, p.parsePostfixExpression)

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
	case token.Let:
		return p.parseLetStatement()
	case token.Const:
		return p.parseConstStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseIdentifier() ast.Expression {
	postfix := p.postfixParseFuncs[p.peekToken.Type]
	if postfix != nil {
		p.nextToken()
		return postfix()
	}

	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) nextToken() {
	p.prevToken = p.currentToken
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// Errors is simply a helper function that returns the parser's errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf(
		"Line: %d: Expected next token to be %s, got: %s instead", p.currentToken.Line, t, p.peekToken.Type,
	)
	p.errors = append(p.errors, msg)
}

func (p *Parser) currentTokenTypeIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenTypeIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeekType(t token.Type) bool {
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

	return Lowest
}

func (p *Parser) currenTokenPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return Lowest
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeekType(token.Identifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeekType(token.Equal) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpr(Lowest)

	if fl, ok := stmt.Value.(*ast.FunctionLiteral); ok {
		fl.Name = stmt.Name.Value
	}

	if p.peekTokenTypeIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.currentToken}

	if !p.expectPeekType(token.Identifier) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeekType(token.Equal) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpr(Lowest)

	if fl, ok := stmt.Value.(*ast.FunctionLiteral); ok {
		fl.Name = stmt.Name.Value
	}

	if p.peekTokenTypeIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	stmt.ReturnValue = p.parseExpr(Lowest)

	if p.peekTokenTypeIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExprStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}
	stmt.Expression = p.parseExpr(Lowest)

	if p.peekTokenTypeIs(token.Semicolon) {
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

	for !p.peekTokenTypeIs(token.Semicolon) && precedence < p.peekTokenPrecedence() {
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
		msg := fmt.Sprintf("Line %d: Could not parse %q as integer", p.currentToken.Line, p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.currentToken,
		Value: p.currentTokenTypeIs(token.True),
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	expr := p.parseExpr(Lowest)

	if !p.expectPeekType(token.RightParen) {
		return nil
	}

	return expr
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.currentTokenTypeIs(token.RightBrace) && !p.currentTokenTypeIs(token.EOF) {
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

	if !p.expectPeekType(token.LeftParen) {
		return nil
	}

	p.nextToken()
	expr.Condition = p.parseExpr(Lowest)

	if !p.expectPeekType(token.RightParen) {
		return nil
	}

	if !p.expectPeekType(token.LeftBrace) {
		return nil
	}

	expr.Consequence = p.parseBlockStatement()

	if p.peekTokenTypeIs(token.Else) {
		p.nextToken()

		if !p.expectPeekType(token.LeftBrace) {
			return nil
		}

		expr.Alternative = p.parseBlockStatement()
	}

	return expr
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenTypeIs(token.RightParen) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
	identifiers = append(identifiers, ident)

	for p.peekTokenTypeIs(token.Comma) {
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{
			Token: p.currentToken,
			Value: p.currentToken.Literal,
		}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeekType(token.RightParen) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeekType(token.LeftParen) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeekType(token.LeftBrace) {
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
	array.Elements = p.parseExprList(token.RightBracket)

	return array
}

func (p *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: p.currentToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenTypeIs(token.RightBrace) {
		p.nextToken()
		key := p.parseExpr(Lowest)

		if !p.expectPeekType(token.Colon) {
			return nil
		}

		p.nextToken()

		value := p.parseExpr(Lowest)
		hash.Pairs[key] = value

		if !p.peekTokenTypeIs(token.RightBrace) && !p.expectPeekType(token.Comma) {
			return nil
		}
	}

	if !p.expectPeekType(token.RightBrace) {
		return nil
	}

	return hash
}

func (p *Parser) parseIndexExpr(left ast.Expression) ast.Expression {
	expr := &ast.IndexExpression{Token: p.currentToken, Left: left}
	p.nextToken()
	expr.Index = p.parseExpr(Lowest)

	if !p.expectPeekType(token.RightBracket) {
		return nil
	}

	return expr
}

func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenTypeIs(token.RightParen) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpr(Lowest))

	for p.peekTokenTypeIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpr(Lowest))
	}

	if !p.expectPeekType(token.RightParen) {
		return nil
	}

	return args
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expr := &ast.CallExpression{
		Token:    p.currentToken,
		Function: function,
	}
	expr.Arguments = p.parseExprList(token.RightParen)

	return expr
}

func (p *Parser) parseExprList(end token.Type) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenTypeIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpr(Lowest))

	for p.peekTokenTypeIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpr(Lowest))
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
	expr.Right = p.parseExpr(Prefix)

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

func (p *Parser) parsePostfixExpression() ast.Expression {
	return &ast.PostfixExpression{
		Token:    p.prevToken,
		Operator: p.currentToken.Literal,
	}
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFunc) {
	p.prefixParseFuncs[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.Type, fn infixParseFunc) {
	p.infixParseFuncs[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.Type, fn postfixParseFunc) {
	p.postfixParseFuncs[tokenType] = fn
}

func (p *Parser) noPrefixParseFuncError(t token.Type) {
	msg := fmt.Sprintf("Line %d: No prefix parse function for %s found", p.currentToken.Line, t)
	p.errors = append(p.errors, msg)
}
