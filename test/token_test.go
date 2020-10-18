package test

import (
	"GoParser/internal/lexer"
	"GoParser/internal/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `+-*/^%() `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULT, "*"},
		{token.DIV, "/"},
		{token.POW, "^"},
		{token.MOD, "%"},
		{token.OPEN_BRACKET, "("},
		{token.CLOSE_BRACKET, ")"},
		{token.WHITESPACE, " "},
		{token.EPSILON, "EPSILON"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
func TestNextTokenILLEGAL(t *testing.T) {
	input := `+B`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.ILLEGAL, "ILLEGAL"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}