package test

import (
	"GoParser/internal/lexer"
	"GoParser/internal/token"
	"fmt"
	"reflect"
	"testing"
)
//
//-1 + 2
//5 * 5 + 10
//-5 - 10
//5 * (add(2, 3) + 10)


func TestParseInputLexer(t *testing.T) {
	tests := []struct {
		lexer *lexer.Lexer
		want    []token.Token
	}{
		{lexer.New("1 + 2"), []token.Token{
			{
				Type:    token.INTEGER,
				Literal: "1",
			},
			{
				Type:    token.WHITESPACE,
				Literal: " ",
			},
			{
				Type:    token.PLUS,
				Literal: "+",
			},
			{
				Type:    token.WHITESPACE,
				Literal: " ",
			},
			{
				Type:    token.INTEGER,
				Literal: "2",
			},
			{
				Type:    token.EPSILON,
				Literal: "EPSILON",
			},
		}},
		{lexer.New("1.33 + 2"), []token.Token{
			{
				Type:    token.DOUBLE,
				Literal: "1.33",
			},
			{
				Type:    token.WHITESPACE,
				Literal: " ",
			},
			{
				Type:    token.PLUS,
				Literal: "+",
			},
			{
				Type:    token.WHITESPACE,
				Literal: " ",
			},
			{
				Type:    token.INTEGER,
				Literal: "2",
			},
			{
				Type:    token.EPSILON,
				Literal: "EPSILON",
			},
		}},
	}

	for i, tt := range tests {
		tokens := tt.lexer.ParseInput()
		fmt.Println(tokens)
		if len(tokens) != len(tt.want) {
			t.Fatalf("tests[%d] - len diff. expected=%v, got=%v",
				i, tt.want, tokens)
		}
		if !reflect.DeepEqual(tokens, tt.want) {
			t.Fatalf("tests[%d] - slice diff. expected=%v, got=%v",
				i, tt.want, tokens)
		}
	}
}