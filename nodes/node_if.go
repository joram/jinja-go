package nodes

import "fmt"

type IfNode struct {
	HasChildren
}

func NewIfNode() INode {
	return &IfNode{NewHasChildren("IF")}
}

func (node *IfNode) ToString() string {
	return fmt.Sprintf("<IF>")
}

func (node *IfNode) Type() int {
	return NODE_TYPE_IF
}

func (node *IfNode) Render(context map[string]interface{}) string {
	return "" // TODO
}
