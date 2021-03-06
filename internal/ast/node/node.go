package node

type NodeType string

const (
	VARIABLE_NODE = "VARIABLE_NODE"
	CONSTANT_NODE = "CONSTANT_NODE"

	ADDITION_NODE    = "ADDITION_NODE"
	SUBTRACTION_NODE = "SUBTRACTION_NODE"

	MULTIPLICATION_NODE = "MULTIPLICATION_NODE"
	DIVISION_NODE       = "DIVISION_NODE"

	EXPONENTIATION_NODE = "EXPONENTIATION_NODE"

	FUNCTION_NODE = "FUNCTION_NODE"
)

type Node interface {
	GetType() NodeType
	GetValue() float64
	AddNode(node Node)
	SetPositive(pos bool)
}
