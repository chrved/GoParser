package node

type SubtractionNode struct {
	nodeType NodeType
	nodes    []Node
}

func NewSubtractionNode(values []Node) *SubtractionNode {
	return &SubtractionNode{
		nodeType: SUBTRACTION_NODE,
		nodes:    values,
	}
}
func (s *SubtractionNode) GetType() NodeType {
	return s.nodeType
}
func (s *SubtractionNode) GetValue() float64 {
	var result float64
	result = s.nodes[0].GetValue()
	for _, node := range s.nodes[1:] {
		result -= node.GetValue()
	}

	return result
}
func (s *SubtractionNode) SetPositive(pos bool) {}
func (s *SubtractionNode) AddNode(node Node) {
	s.nodes = append(s.nodes, node)
}
