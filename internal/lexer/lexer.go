package lexer

import (
	"GoParser/internal/token"
	"regexp"
)

type Lexer struct {
	input         string
	position      int
	readPosition  int
	char          byte
	tokenOpRegex  map[string]*regexp.Regexp
	tokenNumRegex map[string]*regexp.Regexp
}

func New(input string) *Lexer {
	l := &Lexer{
		input:         input,
		position:      0,
		tokenOpRegex:  createOpRegex(),
		tokenNumRegex: createNumRegex(),
	}

	return l
}

func (l *Lexer) ParseInput() []token.Token {
	var tokens []token.Token

	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == token.EPSILON || tok.Type == token.ILLEGAL {
			return tokens
		}
	}
}

func (l *Lexer) NextToken() token.Token {
	if len(l.input) == 0 {
		return newToken(token.EPSILON, "EPSILON")
	}

	if ind := l.tokenNumRegex[token.NUMBER].FindStringIndex(l.input); ind != nil {
		if decNum := l.tokenNumRegex[token.DOUBLE].FindStringIndex(l.input); decNum != nil {
			tok := newToken(token.TokenType("DOUBLE"), l.input[decNum[0]:decNum[1]])
			l.input = l.input[decNum[1]:]
			return tok
		} else if intNum := l.tokenNumRegex[token.INTEGER].FindStringIndex(l.input); intNum != nil {
			tok := newToken(token.TokenType("INTEGER"), l.input[intNum[0]:intNum[1]])
			l.input = l.input[intNum[1]:]
			return tok
		}
	}
	for key, val := range l.tokenOpRegex {
		if ind := val.FindStringIndex(l.input); ind != nil {
			tok := newToken(token.TokenType(key), l.input[ind[0]:ind[1]])
			l.input = l.input[ind[1]:]
			return tok
		}
	}
	return newToken(token.ILLEGAL, "ILLEGAL")
}

func createOpRegex() map[string]*regexp.Regexp {
	m := make(map[string]*regexp.Regexp)
	m[token.WHITESPACE] = regexp.MustCompile("^ ")
	m[token.PLUS] = regexp.MustCompile("^\\+")
	m[token.MINUS] = regexp.MustCompile("^-")
	m[token.DIV] = regexp.MustCompile("^/")
	m[token.MULT] = regexp.MustCompile("^\\*")
	m[token.MOD] = regexp.MustCompile("^%")
	m[token.POW] = regexp.MustCompile("^\\^")
	m[token.OPEN_BRACKET] = regexp.MustCompile("^\\(")
	m[token.CLOSE_BRACKET] = regexp.MustCompile("^\\)")

	return m
}

func createNumRegex() map[string]*regexp.Regexp {
	m := make(map[string]*regexp.Regexp)
	m[token.NUMBER] = regexp.MustCompile("^\\d+(\\.\\d+)?")
	m[token.INTEGER] = regexp.MustCompile("^[0-9]+")
	m[token.DOUBLE] = regexp.MustCompile("^[0-9]+\\.[0-9]+")

	return m
}
func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: literal,
	}
}
