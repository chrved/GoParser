package token

import (
	"regexp"
)

type tokenInfo struct {
	regex     *regexp.Regexp
	tokenType TokenType
}
type Tokenizer struct {
	tokens map[string]tokenInfo
	result []Token

	input        string
	readPosition int
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		tokens: make(map[string]tokenInfo, 0),
		result: make([]Token, 0),
	}
}

func (t *Tokenizer) AddToken(token TokenType, regex string) *Tokenizer {
	t.tokens[string(token)] = tokenInfo{
		regex:     regexp.MustCompile(regex),
		tokenType: token,
	}
	return t
}
func (t *Tokenizer) SetInputString(input string) *Tokenizer {
	t.input = input
	return t
}
func (t *Tokenizer) Parse() *Tokenizer {

	for {
		tok := t.getNextToken()
		t.result = append(t.result, tok)
		if tok.Type == EPSILON || tok.Type == ILLEGAL {
			break
		}
	}
	return t
}
func (t *Tokenizer) GetParsedToken() []Token {
	return t.result
}

func (t *Tokenizer) GetTokenInfo() map[string]tokenInfo {
	return t.tokens
}

func (t *Tokenizer) getNextToken() Token {
	if len(t.input) == 0 {
		return NewToken(EPSILON, "EPSILON")
	}
	for key, tokeninfo := range t.tokens {
		if index := tokeninfo.regex.FindStringIndex(t.input); index != nil {
			tok := NewToken(TokenType(key), t.input[index[0]:index[1]])
			t.input = t.input[index[1]:]
			return tok
		}
	}
	return NewToken(ILLEGAL, "ILLEGAL")
}

func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}
