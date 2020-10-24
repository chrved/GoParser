package test

import (
	"GoParser/internal/token"
	"reflect"
	"testing"
)

func TestTokenizerAddToken(t *testing.T) {

	tests := []struct {
		tokens    []token.TokenType
		input     []string
		nrOfToken int
	}{
		{[]token.TokenType{token.PLUS, token.MINUS}, []string{"^\\+", "^\\-"}, 2},
		{[]token.TokenType{token.PLUS, token.MINUS, token.PLUS, token.MINUS}, []string{"^\\+", "^\\-", "^\\*", "^/"}, 4},
	}

	for i, tt := range tests {
		tokenizer := token.NewTokenizer()
		for index, _ := range tt.tokens {
			tokenizer.AddToken(tt.tokens[index], tt.input[index])
		}

		if len(tokenizer.GetTokenInfo()) != tt.nrOfToken {
			t.Fatalf("tests[%d] - Nr of tokens wrong. expected=%q, got=%q",
				i, tt.nrOfToken, len(tokenizer.GetTokenInfo()))
		}
	}
}
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
		AddToken(token.NUMBER, "^\\d+(\\.\\d+)?")
	//AddToken(token.INTEGER, "^[0-9]+").
	//AddToken(token.DOUBLE,"^[0-9]+\\.[0-9]+")
}
