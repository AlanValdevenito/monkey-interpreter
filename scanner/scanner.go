package scanner

import (
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

// Scanner is responsible for reading the input string and producing tokens.
type Scanner struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a new instance of Scanner with the given input string.
func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()
	return s
}

// ----- Public methods -----

// NextToken returns the next token from the input.
func (s *Scanner) NextToken() token.Token {
	var tok token.Token

	s.skipWhitespace()

	switch s.ch {
	case '(': 
		tok = newToken(token.LPAREN, s.ch)
	case ')':
		tok = newToken(token.RPAREN, s.ch)
	case '{':
		tok = newToken(token.LBRACE, s.ch)
	case '}':
		tok = newToken(token.RBRACE, s.ch)
	case '=':
		if eqToken, isDoubleEqual := s.makeTwoCharToken(token.EQ, '='); isDoubleEqual {
			tok = eqToken
		} else {
			tok = newToken(token.ASSIGN, s.ch)
		}
	case '+':
		tok = newToken(token.PLUS, s.ch)
	case '-':
		tok = newToken(token.MINUS, s.ch)
	case '/':
		tok = newToken(token.SLASH, s.ch)
	case '*':
		tok = newToken(token.ASTERISK, s.ch)
	case '<':
		tok = newToken(token.LT, s.ch)
	case '>':
		tok = newToken(token.GT, s.ch)
	case '!':
		if notEqToken, isNotEqual := s.makeTwoCharToken(token.NOT_EQ, '='); isNotEqual {
			tok = notEqToken
		} else {
			tok = newToken(token.BANG, s.ch)
		}
	case ',':
		tok = newToken(token.COMMA, s.ch)
	case ';':
		tok = newToken(token.SEMICOLON, s.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(s.ch) {
			tok.Literal = s.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(s.ch) {
			tok.Type = token.INT
			tok.Literal = s.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, s.ch)
		}
	}

	s.readChar()
	return tok
}

// ----- Private methods -----

// readChar reads the next character from the input and advances the positions accordingly.
func (s *Scanner) readChar() {
	if s.readPosition >= len(s.input) {
		s.ch = 0 // ASCII code for NUL, signifies end of input
	} else {
		s.ch = s.input[s.readPosition]
	}
	s.position = s.readPosition
	s.readPosition++
}

// newToken creates a new token with the given type and character.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter checks if the given character is a letter (a-z or A-Z).
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// readIdentifier reads an identifier from the input and returns it as a string.
func (s *Scanner) readIdentifier() string {
	position := s.position
	for isLetter(s.ch) {
		s.readChar()
	}
	return s.input[position:s.position]
}

// skipWhitespace advances the scanner's position past any whitespace characters.
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' || s.ch == '\n' || s.ch == '\r' {
		s.readChar()
	}
}

// isDigit checks if the given character is a digit (0-9).
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readNumber reads a number from the input and returns it as a string.
func (s *Scanner) readNumber() string {
	position := s.position
	for isDigit(s.ch) {
		s.readChar()
	}
	return s.input[position:s.position]
}

// peekChar returns the next character without advancing the scanner's position.
func (s *Scanner) peekChar() byte {
	if s.readPosition >= len(s.input) {
		return 0
	}
	return s.input[s.readPosition]
}

// makeTwoCharToken checks if the next character matches 'expected', and if so, advances and returns a token with the combined literal.
func (s *Scanner) makeTwoCharToken(tokenType token.TokenType, expected byte) (token.Token, bool) {
	if s.peekChar() == expected {
		ch := s.ch
		s.readChar()
		return token.Token{Type: tokenType, Literal: string(ch) + string(s.ch)}, true
	}
	return token.Token{}, false
}