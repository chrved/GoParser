package node

import "math"

type ExponentiationNode struct {
	nodeType NodeType
	nodes    []Node
}

func NewExponentiationNode(values []Node) *ExponentiationNode {
	return &ExponentiationNode{
		nodeType: EXPONENTIATION_NODE,
		nodes:    values,
	}
}
func (e *ExponentiationNode) GetType() NodeType {
	return e.nodeType
}
func (e *ExponentiationNode) GetValue() float64 {
	base := e.nodes[0].GetValue()
	exponent := e.nodes[1].GetValue()
	return math.Pow(base, exponent)
}
func (e *ExponentiationNode) SetPositive(pos bool) {}
func (e *ExponentiationNode) AddNode(node Node) {
	e.nodes = append(e.nodes, node)
}
