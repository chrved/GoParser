package test

import (
	"GoParser/internal/token"
	"reflect"
	"testing"
)

func TestTokenizerParse(t *testing.T) {

	tests := []struct {
		input string
		want  []token.Token
	}{
		{"1.33 + 2",
			[]token.Token{
				token.NewToken(token.NUMBER, "1.33"),
				token.NewToken(token.WHITESPACE, " "),
				token.NewToken(token.PLUS, "+"),
				token.NewToken(token.WHITESPACE, " "),
				token.NewToken(token.NUMBER, "2"),
				token.NewToken(token.EPSILON, "EPSILON")}},
		{"sin(2)",
			[]token.Token{
				token.NewToken(token.FUNCTION, "sin"),
				token.NewToken(token.OPEN_BRACKET, "("),
				token.NewToken(token.NUMBER, "2"),
				token.NewToken(token.CLOSE_BRACKET, ")"),
				token.NewToken(token.EPSILON, "EPSILON")}},
	}

	for i, tt := range tests {
		tok := createTokenizer().SetInputString(tt.input)
		tok.Parse()
		res := tok.GetParsedToken()

		if len(res) == 0 {
			t.Fatalf("tests[%d] - Nr of tokens is 0.", i)
		}
		if !reflect.DeepEqual(res, tt.want) {
			t.Fatalf("tests[%d] - Not equal \nwant: %v \ngot: %v", i, tt.want, res)
		}
	}
}
func createTokenizer() *token.Tokenizer {
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
