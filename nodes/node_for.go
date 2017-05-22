package nodes

import "fmt"

type ForNode struct {
	HasChildren
}

func NewForNode() INode {
	return &ForNode{NewHasChildren("FOR")}
}

func (node *ForNode) ToString() string {
	return fmt.Sprintf("<FOR>")
}

func (node *ForNode) Type() int {
	return NODE_TYPE_FOR
}

func (node *ForNode) Render(context map[string]interface{}) string {
	return "" // TODO
}
