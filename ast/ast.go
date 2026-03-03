package ast

// This file defines the abstract syntax tree (AST) structures for the Monkey programming language.

// Node is the base interface for all AST nodes.
// Every node in our AST has to implement the Node interface.
type Node interface {
	TokenLiteral() string // TokenLiteral returns the literal value of the token associated with this node. Will be used only for debugging and testing.
}

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
