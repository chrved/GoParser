package util

import "GoParser/internal/token"

type stack struct {
	stack []token.Token
}

func NewStack() *stack {
	return &stack{stack: make([]token.Token, 0)}
}

func (s *stack) Pop() token.Token {
	index := len(s.stack) - 1
	tok := s.stack[index]
	s.stack = s.stack[:index]
	return tok
}

func (s *stack) Push(tok token.Token) {
	s.stack = append(s.stack, tok)
}

func (s *stack) Empty() bool {
	return len(s.stack) == 0
}
