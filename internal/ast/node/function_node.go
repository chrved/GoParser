package node

import (
	"math"
	"strings"
)

type FunctionNode struct {
	nodeType     NodeType
	nodes        Node
	functionType string
}

func NewFunctionNode(functionType string, value Node) *FunctionNode {
	return &FunctionNode{
		nodeType:     FUNCTION_NODE,
		nodes:        value,
		functionType: strings.ToUpper(functionType),
	}
}
func (f *FunctionNode) GetType() NodeType {
	return f.nodeType
}
func (f *FunctionNode) GetValue() float64 {
	switch f.functionType {
	case "SIN":
		return math.Sin(f.nodes.GetValue())
	case "COS":
		return math.Cos(f.nodes.GetValue())
	case "TAN":
		return math.Tan(f.nodes.GetValue())
	case "SQRT":
		return math.Sqrt(f.nodes.GetValue())
	}
	panic("Unknown function " + f.functionType)
}
func (f *FunctionNode) SetPositive(pos bool) {}
func (f *FunctionNode) AddNode(node Node)    {}
