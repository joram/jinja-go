package nodes

import "fmt"

type VariableNode struct {
	body string
	HasChildren
}

func NewVariableNode(body string) INode {
	return &VariableNode{body: body}
}
func (node *VariableNode) ToString() string {
	return fmt.Sprintf("<VAR>")
}

func (node *VariableNode) Type() int {
	return NODE_TYPE_VARIABLE
}
