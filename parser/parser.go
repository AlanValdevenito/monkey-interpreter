package parser

import (
	"fmt"

	"github.com/AlanValdevenito/monkey-interpreter/ast"
	"github.com/AlanValdevenito/monkey-interpreter/scanner"
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

type Parser struct {
	s *scanner.Scanner

	currentToken token.Token
	peekToken    token.Token

	errors []string
}

func New(s *scanner.Scanner) *Parser {
	p := &Parser{
		s:      s,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	return p
}

// ----- Public methods -----

// ParseProgram parses the entire input and returns the root node of the AST.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // Root node of the AST
	program.Statements = []ast.Statement{}

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
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
	default:
		return nil
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

// peekError adds an error message to the parser's error list when the next token does not match the expected type.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
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
