package parser

import (
	"fmt"

	"github.com/AlanValdevenito/monkey-interpreter/ast"
	"github.com/AlanValdevenito/monkey-interpreter/scanner"
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

type (
	prefixParseFn func() ast.Expression
)

const (
	_ int = iota // We use iota to create a set of constants that represent the precedence of different operators. The blank identifier (_) is used to skip the zero value, so that our precedence levels start at 1.
	LOWEST
	EQUALS      // Example: ==
	LESSGREATER // Example: > or <
	SUM         // Example: +
	PRODUCT     // Example: *
	PREFIX      // Example: -X or !X
	CALL        // Example: myFunction(X)
)

type Parser struct {
	s *scanner.Scanner

	currentToken token.Token
	peekToken    token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
}

func New(s *scanner.Scanner) *Parser {
	p := &Parser{
		s:      s,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	// Registering the prefix parse functions for the different token types.
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)

	return p
}

// ----- Public methods -----

// ParseProgram parses the entire input and returns the root node of the AST.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // Root node of the AST
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}

	return program
}

// Errors returns a slice of error messages encountered during parsing.
func (p *Parser) Errors() []string {
	return p.errors
}

// ----- Private methods -----

// nextToken advances the parser to the next token in the input.
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.s.NextToken()
}

// parseStatement parses a single statement based on the current token.
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement parses a let statement and returns an AST node representing it.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement parses a return statement and returns an AST node representing it.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we encounter a semicolon
	for !p.currentTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement parses an expression statement and returns an AST node representing it.
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression parses an expression based on the current token and the given precedence level. It uses the prefix and infix parse functions registered for the current token type to parse the expression correctly.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()

	return leftExp
}

// peekError adds an error message to the parser's error list when the next token does not match the expected type.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// ----- Pratt parsing methods -----

// parseIdentifier parses an identifier and returns an AST node representing it.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

// ----- Utility methods -----

// currentTokenIs checks if the current token matches the given type.
func (p *Parser) currentTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

// peekTokenIs checks if the next token matches the given type.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks if the next token matches the given type and advances the parser if it does.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

// registerPrefix registers a prefix parse function for a given token type.
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
