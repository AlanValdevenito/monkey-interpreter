package parser

import (
	"github.com/AlanValdevenito/monkey-interpreter/ast"
	"github.com/AlanValdevenito/monkey-interpreter/scanner"
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

type Parser struct {
	s *scanner.Scanner

	currentToken token.Token
	peekToken    token.Token
}

func New(s *scanner.Scanner) *Parser {
	p := &Parser{s: s}

	p.nextToken()
	p.nextToken()

	return p
}

// ----- Public methods -----

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

// ----- Private methods -----

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.s.NextToken()
}
