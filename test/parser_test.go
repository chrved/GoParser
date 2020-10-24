package test

import (
	"GoParser/internal/parser"
	"GoParser/internal/token"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tokens := getTokenizer().SetInputString("1.33 + 2").Parse().GetParsedToken()

	p := parser.New(tokens)
	p.Parse()
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
		AddToken(token.NUMBER, "^\\d+(\\.\\d+)?")
	//AddToken(token.INTEGER, "^[0-9]+").
	//AddToken(token.DOUBLE,"^[0-9]+\\.[0-9]+")
}

func TestNew(t *testing.T) {
	type args struct {
		tokens []token.Token
	}
	tests := []struct {
		name string
		args args
		want *parser.Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser.New(tt.args.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestParser_Parse(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}

//func TestParser_expression(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func TestParser_factor(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func TestParser_getNextToken(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func TestParser_number(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func TestParser_primary(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func TestParser_term(t *testing.T) {
//	type fields struct {
//		input     []token.Token
//		lookahead token.Token
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &parser.Parser{
//				input:     tt.fields.input,
//				lookahead: tt.fields.lookahead,
//			}
//		})
//	}
//}
//
//func Test_getTokenType(t *testing.T) {
//	type args struct {
//		tok token.Token
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := parser.getTokenType(tt.args.tok); got != tt.want {
//				t.Errorf("getTokenType() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
