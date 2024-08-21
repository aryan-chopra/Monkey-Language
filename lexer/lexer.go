package lexer

import (
	"Interpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}


func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) GetNextToken() token.Token {
	var tok token.Token
	
	l.skipWhitespaces()
	
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else{
			tok = generateNewToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = generateNewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = generateNewToken(token.LPAREN, l.ch)
	case ')':
		tok = generateNewToken(token.RPAREN, l.ch)
	case ',':
		tok = generateNewToken(token.COMMA, l.ch)
	case '+':
		tok = generateNewToken(token.PLUS, l.ch)
	case '-':
		tok = generateNewToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else{
			tok = generateNewToken(token.BANG, l.ch)
		}
	case '*':
		tok = generateNewToken(token.ASTERISK, l.ch)
	case '/':
		tok = generateNewToken(token.SLASH, l.ch)
	case '<':
		tok = generateNewToken(token.LT, l.ch)
	case '>':
		tok = generateNewToken(token.GT, l.ch)
	case '{':
		tok = generateNewToken(token.LBRACE, l.ch)
	case '}':
		tok = generateNewToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdentifierType(tok.Literal)
			return tok
		} 
		if isInteger(l.ch) {
			tok.Literal = l.readInteger()
			tok.Type = token.INT
			return tok
		} else {
			tok = generateNewToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l* Lexer) readIdentifier() string {
	position := l.position
	
	for isLetter(l.ch) {
		l.readChar()
	}
	
	return l.input[position:l.position]
}

func (l *Lexer) readInteger() string {
	position := l.position
	
	for isInteger(l.ch) {
		l.readChar()
	}
	
	return l.input[position : l.position]
}

func (l *Lexer) peekChar() byte {
	if (l.readPosition >= len(l.input)) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isInteger(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func generateNewToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
