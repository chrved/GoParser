package test

import (
	"GoParser/internal/token"
	"GoParser/internal/util"
	"testing"
)

func TestStack(t *testing.T) {

	stack := util.NewStack()

	if !stack.Empty() {
		t.Error("stack not empty")
	}
	stack.Push(token.Token{Literal: "1"})

	if stack.Empty() {
		t.Error("stack not empty")
	}
	stack.Push(token.Token{Literal: "2"})
	stack.Push(token.Token{Literal: "3"})
	stack.Push(token.Token{Literal: "4"})

	tok := stack.Pop()
	if tok.Literal != "4" {
		t.Error("stack not returning right value")
	}
	tok = stack.Pop()
	if tok.Literal != "3" {
		t.Error("stack not returning right value")
	}
	tok = stack.Pop()
	if tok.Literal != "2" {
		t.Error("stack not returning right value")
	}
	tok = stack.Pop()
	if tok.Literal != "1" {
		t.Error("stack not returning right value")
	}
	if !stack.Empty() {
		t.Error("stack not empty")
	}
}
