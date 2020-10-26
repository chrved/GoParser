package node

type DivisionNode struct {
	nodeType NodeType
	nodes    []Node
}

func NewDivisionNode(values []Node) *DivisionNode {
	return &DivisionNode{
		nodeType: DIVISION_NODE,
		nodes:    values,
	}
}
func (d *DivisionNode) GetType() NodeType {
	return d.nodeType
}
func (d *DivisionNode) GetValue() float64 {
	var result float64
	result = d.nodes[0].GetValue()
	for _, node := range d.nodes[1:] {
		result /= node.GetValue()
	}

	return result
}
func (d *DivisionNode) SetPositive(pos bool) {}
func (d *DivisionNode) AddNode(node Node) {
	d.nodes = append(d.nodes, node)
}
