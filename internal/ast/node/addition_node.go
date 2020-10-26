package node

type AdditionNode struct {
	nodeType NodeType
	nodes    []Node
}

func NewAdditionNode(values []Node) *AdditionNode {
	return &AdditionNode{
		nodeType: ADDITION_NODE,
		nodes:    values,
	}
}
func (a *AdditionNode) GetType() NodeType {
	return a.nodeType
}
func (a *AdditionNode) GetValue() float64 {
	var result float64
	result = a.nodes[0].GetValue()
	for _, node := range a.nodes[1:] {
		result += node.GetValue()
	}

	return result
}
func (a *AdditionNode) SetPositive(pos bool) {}
func (a *AdditionNode) AddNode(node Node) {
	a.nodes = append(a.nodes, node)
}
