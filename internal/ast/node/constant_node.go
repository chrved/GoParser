package node

import (
	"strconv"
)

type ConstantNode struct {
	nodeType NodeType
	value    float64
	positive bool
}

func NewConstantNode(value string, positive bool) *ConstantNode {
	if val, err := strconv.ParseFloat(value, 64); err == nil {
		return &ConstantNode{
			nodeType: CONSTANT_NODE,
			value:    val,
			positive: positive,
		}
	}
	panic("Cant parse:" + value + " to float.")
}
func (c *ConstantNode) GetType() NodeType {
	return c.nodeType
}
func (c *ConstantNode) GetValue() float64 {
	if c.positive {
		return c.value
	} else {
		return c.value * -1
	}
}
func (c *ConstantNode) SetPositive(pos bool) {
	c.positive = pos
}
func (c *ConstantNode) AddNode(node Node) {}
