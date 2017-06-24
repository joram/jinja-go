package nodes

import "fmt"

type ForNode struct {
	HasChildren
}

func NewForNode(content string) INode {
	return &ForNode{NewHasChildren("FOR", content)}
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
