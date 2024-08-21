package ast

import (
	"Interpreter/token"
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// Root
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (program *Program) String() string {
	var out bytes.Buffer
	
	for _, statement := range program.Statements {
		out.WriteString(statement.String())
	}
	
	return out.String()
}

type LetStatement struct {
	Token token.Token //token.LET
	Name  *Identifier
	Value Expression
}

func (letStatement *LetStatement) statementNode() {}
func (letStatement *LetStatement) TokenLiteral() string {
	return letStatement.Token.Literal
}

func (letStatement *LetStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(letStatement.TokenLiteral() + " ")
	out.WriteString(letStatement.Name.String())
	out.WriteString(" = ")
	
	if letStatement.Value != nil {
		out.WriteString(letStatement.Value.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode() {}
func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}

func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer
	
	out.WriteString(returnStatement.TokenLiteral() + " ")
	
	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}
	
	out.WriteString(";")
	
	return out.String()
}

type ExpressionStatement struct {
	token token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {return es.TokenLiteral()}

func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.String()
	}
	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (ident *Identifier) expressionNode() {}
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Literal
}

func (identifier *Identifier) String() string {
	return identifier.Value
}
