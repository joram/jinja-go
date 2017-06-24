package nodes

import (
	"fmt"
)

type IfNode struct {
	HasChildren
	Content string
}

func NewIfNode(content string) INode {
	return &IfNode{NewHasChildren("IF", content), content}
}

func (node *IfNode) ToString() string {
	return fmt.Sprintf("<IF>")
}

func (node *IfNode) Type() int {
	return NODE_TYPE_IF
}

func (node *IfNode) Render(context map[string]interface{}) string {
	rendered := ""
	for _, child := range node.Children {
		rendered += (*child).Render(context)
	}
	return rendered
}
