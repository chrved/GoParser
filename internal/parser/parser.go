package parser

import (
	"GoParser/internal/ast/node"
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

	result node.Node
}

func New(tokens []token.Token) *Parser {
	return &Parser{input: tokens}
}
func (p *Parser) GetResult() float64 {
	return p.result.GetValue()
}
func (p *Parser) Parse() {
	p.getNextToken()
	fmt.Printf("START: \n %v\n%v\n", p.lookahead, p.input)
	p.result = p.expression()

	if p.lookahead.Type != token.EPSILON {
		panic("END not ok")
	}
	fmt.Printf("END: \n %v\n%v\n", p.lookahead, p.input)
}
func (p *Parser) expression() node.Node {
	//expression : signed_term expression_sum
	exp := p.signed_term()
	return p.expression_sum(exp)
}
func (p *Parser) expression_sum(arg node.Node) node.Node {
	//expression_sum: PLUS-MINUS term expression_sum | EPSILON
	if p.lookahead.Type == token.PLUS || p.lookahead.Type == token.MINUS {
		token_sign := p.lookahead
		p.getNextToken()
		term := p.term()
		var sumnode node.Node
		if token_sign.Type == token.PLUS {
			sumnode = node.NewAdditionNode([]node.Node{arg, term})
		} else {
			sumnode = node.NewSubtractionNode([]node.Node{arg, term})
		}

		return p.expression_sum(sumnode)
	}
	return arg
}

func (p *Parser) term() node.Node {
	fac := p.factor()
	return p.term_sum(fac)
}
func (p *Parser) term_sum(arg node.Node) node.Node {
	if p.lookahead.Type == token.MULT || p.lookahead.Type == token.DIV {
		var mult_div_node node.Node
		if p.lookahead.Type == token.MULT {
			mult_div_node = node.NewMultiplicationNode([]node.Node{arg})
		} else {
			mult_div_node = node.NewMultiplicationNode([]node.Node{arg})
		}
		p.getNextToken()
		fac := p.signed_factor()
		mult_div_node.AddNode(fac)

		return p.term_sum(mult_div_node)
	}
	return arg
}
func (p *Parser) signed_term() node.Node {
	if p.lookahead.Type == token.MINUS || p.lookahead.Type == token.PLUS {
		//positive := true
		//if p.lookahead.Type == token.MINUS {
		//	positive = false
		//}
		p.getNextToken()

		term := p.term()
		//ERROR
		return term
	}
	return p.term()
}

func (p *Parser) factor() node.Node {
	arg := p.argument()
	return p.factor_sum(arg)
}
func (p *Parser) factor_sum(arg node.Node) node.Node {
	if p.lookahead.Type == token.POW {
		p.getNextToken()
		exponent := p.signed_factor()
		return node.NewExponentiationNode([]node.Node{arg, exponent})
	}

	return arg
}
func (p *Parser) signed_factor() node.Node {
	if p.lookahead.Type == token.MINUS || p.lookahead.Type == token.PLUS {
		positive := true
		if p.lookahead.Type == token.MINUS {
			positive = false
		}
		p.getNextToken()
		fac := p.factor()
		fac.SetPositive(positive)
		return node.NewAdditionNode([]node.Node{fac})
	}
	return p.factor()
}

func (p *Parser) argument() node.Node {
	if p.lookahead.Type == token.FUNCTION {
		funcType := p.lookahead.Literal
		p.getNextToken()
		funcValue := p.argument()
		return node.NewFunctionNode(funcType, funcValue)
	}
	if p.lookahead.Type == token.OPEN_BRACKET {
		p.getNextToken()
		exp := p.expression()

		if p.lookahead.Type != token.CLOSE_BRACKET {
			panic("Wrong token")
		}

		p.getNextToken()
		return exp
	}

	return p.value()
}
func (p *Parser) value() node.Node {
	if p.lookahead.Type == token.NUMBER {
		numberNode := node.NewConstantNode(p.lookahead.Literal, true)
		p.getNextToken()
		return numberNode
	} else {
		panic("Not a NUMBER")
	}
}

func (p *Parser) getNextToken() {
	if p.input[0].Type == token.WHITESPACE {
		p.removeWhitespace()
	}
	p.lookahead = p.input[0]
	p.input = p.input[1:]

}
func (p *Parser) removeWhitespace() {
	for {
		if p.input[0].Type == token.WHITESPACE {
			p.input = p.input[1:]
		} else {
			break
		}
	}
}

func printFound(token token.Token) {
	fmt.Printf("Found %v\n", token)
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
