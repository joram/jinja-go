package nodes

import "fmt"

type VariableNode struct {
	HasNoChildren
	body string
}

func NewVariableNode(body string) INode {
	return &VariableNode{NewHasNoChildren("VAR", body), body}
}
func (node *VariableNode) ToString() string {
	return fmt.Sprintf("<VAR>")
}

func (node *VariableNode) Type() int {
	return NODE_TYPE_VARIABLE
}

func (node *VariableNode) Render(context map[string]interface{}) string {
	return "" // TODO
}
