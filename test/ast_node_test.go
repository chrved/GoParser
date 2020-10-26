package test

import (
	"GoParser/internal/ast/node"
	"testing"
)

func TestConstNode(t *testing.T) {
	c := node.NewConstantNode("5", true)
	if c.GetValue() != 5 {
		t.Fatalf("Wrong value %v", c)
	}
	cc := node.NewConstantNode("5.2", false)
	if cc.GetValue() != -5.2 {
		t.Fatalf("Wrong value %v", cc)
	}
}

func TestAddNode(t *testing.T) {
	tests := []struct {
		input []node.Node
		want  float64
	}{
		{[]node.Node{node.NewConstantNode("5", true), node.NewConstantNode("5", true)}, 10},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", true)}, 0},
		{[]node.Node{node.NewConstantNode("-5", false), node.NewConstantNode("5", true)}, 10},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", false)}, -10},
	}

	for i, tt := range tests {
		add := node.NewAdditionNode(tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}

func TestSubNode(t *testing.T) {
	tests := []struct {
		input []node.Node
		want  float64
	}{
		{[]node.Node{node.NewConstantNode("5", true), node.NewConstantNode("5", true)}, 0},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", true)}, -10},
		{[]node.Node{node.NewConstantNode("-5", true), node.NewConstantNode("5", true)}, -10},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", false)}, 0},
	}

	for i, tt := range tests {
		add := node.NewSubtractionNode(tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}
func TestMultNode(t *testing.T) {
	tests := []struct {
		input []node.Node
		want  float64
	}{
		{[]node.Node{node.NewConstantNode("5", true), node.NewConstantNode("5", true)}, 25},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", true)}, -25},
		{[]node.Node{node.NewConstantNode("-5", true), node.NewConstantNode("5", true)}, -25},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", false)}, 25},
	}

	for i, tt := range tests {
		add := node.NewMultiplicationNode(tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}
func TestDivNode(t *testing.T) {
	tests := []struct {
		input []node.Node
		want  float64
	}{
		{[]node.Node{node.NewConstantNode("5", true), node.NewConstantNode("5", true)}, 1},
		{[]node.Node{node.NewConstantNode("10", true), node.NewConstantNode("2", true)}, 5},
		{[]node.Node{node.NewConstantNode("-5", true), node.NewConstantNode("5", true)}, -1},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", false)}, 1},
	}

	for i, tt := range tests {
		add := node.NewDivisionNode(tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}
func TestExpoNode(t *testing.T) {
	tests := []struct {
		input []node.Node
		want  float64
	}{
		{[]node.Node{node.NewConstantNode("5", true), node.NewConstantNode("2", true)}, 25},
		{[]node.Node{node.NewConstantNode("10", true), node.NewConstantNode("1", true)}, 10},
		{[]node.Node{node.NewConstantNode("2", true), node.NewConstantNode("5", true)}, 32},
		{[]node.Node{node.NewConstantNode("5", false), node.NewConstantNode("5", false)}, -0.00032},
	}

	for i, tt := range tests {
		add := node.NewExponentiationNode(tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}
func TestFuncNode(t *testing.T) {
	tests := []struct {
		inpStr string
		input  node.Node
		want   float64
	}{
		{"sin", node.NewConstantNode("2", true), 0.9092974268256816},
		{"Sqrt", node.NewConstantNode("25", true), 5},
	}

	for i, tt := range tests {
		add := node.NewFunctionNode(tt.inpStr, tt.input)
		sum := add.GetValue()
		if sum != tt.want {
			t.Fatalf("tests[%d] - got wrong value: %f want: %f", i, sum, tt.want)
		}
	}
}
func TestNode(t *testing.T) {
	//3*2^4 + sqrt(1+3) = 50
	mult := node.NewMultiplicationNode(
		[]node.Node{
			node.NewConstantNode("3", true),
			node.NewExponentiationNode(
				[]node.Node{
					node.NewConstantNode("2", true),
					node.NewConstantNode("4", true),
				}),
		})
	sqrt := node.NewFunctionNode(
		"sqrt",
		node.NewAdditionNode(
			[]node.Node{
				node.NewConstantNode("1", true),
				node.NewConstantNode("3", true),
			}))
	add := node.NewAdditionNode([]node.Node{
		mult,
		sqrt,
	})

	sum := add.GetValue()
	want := 50.0

	if sum != want {
		t.Fatalf("Got wrong value: %f want: %f", sum, want)
	}
}
