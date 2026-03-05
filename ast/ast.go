package ast

import (
	"github.com/AlanValdevenito/monkey-interpreter/token"
)

// This file defines the abstract syntax tree (AST) structures for the Monkey programming language.

// Program is the root node of every AST our parser produces.
// A program consists of a series of statements.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ---------- AST Node definition ----------

// Node is the base interface for all AST nodes.
// Every node in our AST has to implement the Node interface.
type Node interface {
	TokenLiteral() string // TokenLiteral returns the literal value of the token associated with this node. Will be used only for debugging and testing.}
}

// ---------- Statement and Expression interfaces ----------

// Statements are an interface that represents a statement in the Monkey language.
// Nodes that are statements will implement this interface.
type Statement interface {
	Node
	statementNode()
}

// Expressions are an interface that represents an expression in the Monkey language.
// Nodes that are expressions will implement this interface.
type Expression interface {
	Node
	expressionNode()
}

// ---------- Statement nodes ----------

// LetStatement represents a variable declaration in the Monkey language.
// The struct contains the token for the 'let' keyword, the name of the variable being declared, and the expression that represents the value being assigned to the variable.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {} // LetStatement struct implements the Statement interface, so we need to define the statementNode method.
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier represents a variable name in the Monkey language.
// The struct contains the token for the identifier and the string value of the identifier.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {} // Identifier struct implements the Expression interface, so we need to define the expressionNode method.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
