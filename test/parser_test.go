package test

import (
	"GoParser/internal/parser"
	"GoParser/internal/token"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{"1.33 + 2 +3 * 3", 12.33},
		{"(1.33 + 2^2) * 3", 15.99},
		{"sin(1.33 + 2^2) * 3", -2.4457926433498907},
		{"5 * (cos(2) + 10)", 47.91926581726429},
		{"-1 + 2 * 3", 5},
		{"-5 - 10", -15},
	}

	for i, tt := range tests {
		tokens := getTokenizer().SetInputString(tt.input).Parse().GetParsedToken()
		p := parser.New(tokens)
		p.Parse()
		sum := p.GetResult()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}

func getTokenizer() *token.Tokenizer {
	return token.NewTokenizer().
		AddToken(token.WHITESPACE, "^ ").
		AddToken(token.PLUS, "^\\+").
		AddToken(token.MINUS, "^-").
		AddToken(token.DIV, "^/").
		AddToken(token.MULT, "^\\*").
		AddToken(token.MOD, "^%").
		AddToken(token.POW, "^\\^").
		AddToken(token.OPEN_BRACKET, "^\\(").
		AddToken(token.CLOSE_BRACKET, "^\\)").
		AddToken(token.NUMBER, "^\\d+(\\.\\d+)?").
		AddToken(token.FUNCTION, "^sin|^cos|^tan|^sqrt")
	//AddToken(token.INTEGER, "^[0-9]+").
	//AddToken(token.DOUBLE,"^[0-9]+\\.[0-9]+")
}
