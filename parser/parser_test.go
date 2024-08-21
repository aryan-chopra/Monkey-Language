package parser

import (
	"Interpreter/ast"
	"Interpreter/lexer"
	"testing"
)

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	
	lexer := lexer.NewLexer(input)
	parser := newParser(lexer)
	
	program := parser.ParseProgram()
	checkParseErrors(t, parser)
	
	if len(program.Statements) != 1 {
		t.Fatalf("Not enough statements, got = %d", len(program.Statements))
	}
	
	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program Statement is not an expression statement, got = %q", program.Statements[0])
	}
	
	identifier, ok := statement.Expression.(*ast.Identifier)
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	
	lexer := lexer.NewLexer(input)
	parser := newParser(lexer)
	
	program := parser.ParseProgram()
	checkParseErrors(t, parser)
	
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3. got=%d", len(program.Statements))
	}
	
	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)
		
		if !ok {
			t.Errorf("Statement not *ast.returnStatement, got=%T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral not 'return', got %q", returnStatement.TokenLiteral())
		}
	}
}

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	lexer := lexer.NewLexer(input)
	parser := newParser(lexer)
	
	program := parser.ParseProgram()
	checkParseErrors(t, parser)
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. Got %d", len(program.Statements))
	}
	
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral() not 'let'. Got =%q", s.TokenLiteral())
		return false
	}
	
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.Statement. Got=%t", s)
		return false
	}
	
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. Got='%s'", name, letStmt.Name.Value)
		return false
	}
	
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. Got='%s'", name, letStmt.Name)
		return false
	}
	
	return true
}

func checkParseErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	
	if len(errors) == 0 {
		return
	}
	
	t.Errorf("Parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
