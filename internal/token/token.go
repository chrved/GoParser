package token

type Token struct {
	Type    TokenType
	Literal string
}
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EPSILON = "EPSILON"

	PLUS  = "PLUS"
	MINUS = "MINUS"
	DIV   = "DIV"
	MULT  = "MULT"
	MOD   = "MOD"

	POW = "POW"

	OPEN_BRACKET = "OPEN_BRACKET"
	CLOSE_BRACKET = "CLOSE_BRACKET"

	WHITESPACE = "WHITESPACE"

	NUMBER = "NUMBER"
	INTEGER = "INTEGER"
	DOUBLE = "DOUBLE"
)

//INTEGER ^[0-9]+
//DOUBLE [0-9]+\.[0-9]+

//WHITESPACE ^\s
//OPEN_BRACKET \(
//CLOSE_BRACKET \)
//PLUS [+]
//MINUS [-]
//DIV [\/]
//MULT [*]
//MOD [%]
//POW [\^]

//FUNCTION ^sin|cos|exp|ln|sqrt

//tokenizer.add("[a-zA-Z][a-zA-Z0-9_]*", TokenType.VARIABLE); // variable
