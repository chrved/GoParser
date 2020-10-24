package parser

import (
	"GoParser/internal/token"
	"fmt"
)

const (
	NUM = "NUM"
	OP  = "OP"
	FUN = "FUN"
)

type Parser struct {
	input     []token.Token
	lookahead token.Token
}

func New(tokens []token.Token) *Parser {
	return &Parser{input: tokens}
}
func (p *Parser) Parse() {
	p.getNextToken()
	fmt.Printf("START: \n %v\n%v\n", p.lookahead, p.input)
	p.expression()

	if p.lookahead.Type != token.EPSILON {
		panic("END not ok")
	}
}
func (p *Parser) expression() {
	fmt.Printf("Enter %s\n", "Expression")

	p.signed_term()
	p.expression_sum()

	fmt.Printf("Leave %s\n", "Expression")
}
func (p *Parser) expression_sum() {
	fmt.Printf("Enter %s\n", "Expression_sum")
	if p.lookahead.Type == token.PLUS {
		fmt.Printf("Found %s %v\n", "+", p.lookahead)
		p.getNextToken()
		p.signed_term()

		fmt.Printf("Leave %s\n", "Expression_sum")
		return
	}
	if p.lookahead.Type == token.MINUS {
		fmt.Printf("Found %s %v\n", "-", p.lookahead)
		p.getNextToken()
		p.signed_term()

		fmt.Printf("Leave %s\n", "Expression_sum")
		return
	}

}
func (p *Parser) signed_term() {
	fmt.Printf("Enter %s\n", "Signed_Term")
	if p.lookahead.Type == token.MINUS {
		fmt.Printf("Found %s %v\n", "-", p.lookahead)
		p.getNextToken()
		p.term()

		fmt.Printf("Leave %s\n", "Signed_Term")
		return
	}
	p.term()

	fmt.Printf("Leave %s\n", "Signed_Term")
}
func (p *Parser) term() {
	fmt.Printf("Enter %s\n", "Term")

	p.factor()

	fmt.Printf("Leave %s\n", "Term")
}
func (p *Parser) factor() {
	fmt.Printf("Enter %s\n", "Factor")

	p.primary()

	fmt.Printf("Leave %s\n", "Factor")
}
func (p *Parser) primary() {
	fmt.Printf("Enter %s\n", "Primary")

	p.number()

	fmt.Printf("Leave %s\n", "Primary")
}
func (p *Parser) number() {
	fmt.Printf("Enter %s\n", "Number")

	if p.lookahead.Type == token.NUMBER {
		fmt.Printf("Found %s %v\n", "Number", p.lookahead)
		p.getNextToken()
	} else {
		panic("Not a NUMBER")
	}

	fmt.Printf("Leave %s\n", "Number")
}
func (p *Parser) getNextToken() {
	if p.input[0].Type == token.WHITESPACE {
		p.removeWhitespace()
	}
	p.lookahead = p.input[0]
	fmt.Printf("Get next token %v\n", p.lookahead)
	p.input = p.input[1:]

}
func (p *Parser) removeWhitespace() {
	fmt.Println("REMOVING WHITESPACE's")
	for {
		if p.input[0].Type == token.WHITESPACE {
			p.input = p.input[1:]
		} else {
			break
		}
	}
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
