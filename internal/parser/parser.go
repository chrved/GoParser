package parser

import (
	"GoParser/internal/token"
	"GoParser/internal/util"
)

const (
	NUM = "NUM"
	OP  = "OP"
	FUN = "FUN"
)

type Parser struct {
	input []token.Token
}

func New(tokens []token.Token) *Parser {
	return &Parser{input: tokens}
}
func (p *Parser) Parse() {
	var output []token.Token
	opStack := util.NewStack()

	for _, tok := range p.input {
		tokType := getTokenType(tok)
		switch tokType {
		case NUM:
			output = append(output, tok)
			break
		case OP:
			// CHANGE THIS
			opStack.Push(tok)

		}
	}
}

func (p *Parser) CreateAST() {
}

func getTokenType(tok token.Token) string {
	if tok.Type == token.DOUBLE || tok.Type == token.INTEGER {
		return NUM
	}
	if tok.Type == token.PLUS ||
		tok.Type == token.MINUS ||
		tok.Type == token.DIV ||
		tok.Type == token.MULT ||
		tok.Type == token.MOD ||
		tok.Type == token.POW {
		return OP
	}
	return token.ILLEGAL
}

///* This implementation does not implement composite functions,functions with variable number of arguments, and unary operators. */
//
//while there are tokens to be read:
//read a token.
//	if the token is a number, then:
//		push it to the output queue.
//	else if the token is a function then:
//		push it onto the operator stack
//	else if the token is an operator then:
//while ((there is an operator at the top of the operator stack)
//and ((the operator at the top of the operator stack has greater precedence)
//or (the operator at the top of the operator stack has equal precedence and the token is left associative))
//and (the operator at the top of the operator stack is not a left parenthesis)):
//pop operators from the operator stack onto the output queue.
//push it onto the operator stack.
//else if the token is a left parenthesis (i.e. "("), then:
//push it onto the operator stack.
//else if the token is a right parenthesis (i.e. ")"), then:
//while the operator at the top of the operator stack is not a left parenthesis:
//pop the operator from the operator stack onto the output queue.
///* If the stack runs out without finding a left parenthesis, then there are mismatched parentheses. */
//if there is a left parenthesis at the top of the operator stack, then:
//pop the operator from the operator stack and discard it
///* After while loop, if operator stack not null, pop everything to output queue */
//if there are no more tokens to read then:
//while there are still operator tokens on the stack:
///* If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses. */
//pop the operator from the operator stack onto the output queue.
//exit.
