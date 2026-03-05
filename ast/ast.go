package ast

import (
	"bytes"

	"github.com/AlanValdevenito/monkey-interpreter/token"
)

// This file defines the abstract syntax tree (AST) structures for the Monkey programming language.

// Program is the root node of every AST our parser produces.
// A program consists of a series of statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the token associated with the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String returns a string representation of the entire program by concatenating the string representations of all its statements.
// The "real work" of this method is done by the String() method of each statement, which recursively calls the String() method of its child nodes.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// ---------- AST Node definition ----------

// Node is the base interface for all AST nodes.
// Every node in our AST has to implement the Node interface.
type Node interface {
	TokenLiteral() string // TokenLiteral returns the literal value of the token associated with this node. Will be used only for debugging and testing.}
	String() string       // String returns a string representation of the node and its children. Will be used for debugging and testing.
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier represents a variable name in the Monkey language.
// The struct contains the token for the identifier and the string value of the identifier.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {} // Identifier struct implements the Expression interface, so we need to define the expressionNode method.
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents a statement that consists of a single expression in the Monkey language.
// The struct contains the token for the first token in the expression and the expression itself.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
