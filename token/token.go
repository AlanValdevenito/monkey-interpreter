package token

// TokenType is a string that represents the type of a token.
type TokenType string

// Token represents a lexical token with its type and literal value.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Special type. Signifies a token/character we don't know about.
	EOF     = "EOF" // Special type. Signifies we've reached the end of the file/input.

	// Identifiers + literals
	IDENT  = "IDENT"  // add, foobar, x, y, ...
	INT    = "INT"    // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)