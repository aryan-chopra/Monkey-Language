package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"Interpreter/token"
	"fmt"
)

type Parser struct {
	lexer *lexer.Lexer
	errors []string
	curToken token.Token
	peekToken token.Token
	
	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

func (parser *Parser) registerPrefix(tokenType token.TokenType, prefixFn prefixParseFn) {
	parser.prefixParseFns[tokenType] = prefixFn
}

func (parser *Parser) registerInfix(tokenType token.TokenType, infixFn infixParseFn) {
	parser.infixParseFns[tokenType] = infixFn 
}

func newParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{lexer: lexer, errors: []string{}}
	
	parser.nextToken()
	parser.nextToken()
	
	return parser
}

func (parser *Parser) nextToken() {
	parser.curToken = parser.peekToken
	parser.peekToken = parser.lexer.GetNextToken()
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) peekError(token token.TokenType) {
	message := fmt.Sprintf("Expected next token to be %s, got %s instead", token, parser.peekToken.Type)
	parser.errors = append(parser.errors, message)
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	
	for parser.curToken.Type != token.EOF {
		statement := parser.parseStatement()
		
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}
	
	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.curToken.Type {
		case token.LET:
			return parser.parseLetStatement()
		case token.RETURN:
			return parser.parseReturnStatement()
		default:
			return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	letStatement := &ast.LetStatement{Token: parser.curToken}
	
	if !parser.expectedTokenOnPeek(token.IDENT) {
		return nil
	}
	
	letStatement.Name = &ast.Identifier{Token: parser.curToken, Value: parser.curToken.Literal}
	
	if !parser.expectedTokenOnPeek(token.ASSIGN) {
		return nil
	}
	
	for !parser.hasCurrentToken(token.SEMICOLON) {
		parser.nextToken()
	}
	
	return letStatement
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	returnStatement := &ast.ReturnStatement{Token: parser.curToken}
	
	parser.nextToken()
	
	for !parser.hasCurrentToken(token.SEMICOLON) {
		parser.nextToken()
	}
	
	return returnStatement
}

func (parser *Parser) hasCurrentToken(expectedToken token.TokenType) bool {
	return parser.curToken.Type == expectedToken
}

func (parser *Parser) hasPeekToken(expectedToken token.TokenType) bool {
	return parser.peekToken.Type == expectedToken
}

func (parser *Parser) expectedTokenOnPeek(token token.TokenType) bool {
	if parser.hasPeekToken(token) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(token)
		return false
	}
}
