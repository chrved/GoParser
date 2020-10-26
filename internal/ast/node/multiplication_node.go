package node

type MultiplicationNode struct {
	nodeType NodeType
	nodes    []Node
}

func NewMultiplicationNode(values []Node) *MultiplicationNode {
	return &MultiplicationNode{
		nodeType: MULTIPLICATION_NODE,
		nodes:    values,
	}
}
func (m *MultiplicationNode) GetType() NodeType {
	return m.nodeType
}
func (m *MultiplicationNode) GetValue() float64 {
	var result float64
	result = m.nodes[0].GetValue()
	for _, node := range m.nodes[1:] {
		result *= node.GetValue()
	}

	return result
}
func (m *MultiplicationNode) SetPositive(pos bool) {}
func (m *MultiplicationNode) AddNode(node Node) {
	m.nodes = append(m.nodes, node)
}
