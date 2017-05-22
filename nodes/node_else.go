package nodes

import "fmt"

type ElseNode struct {
	HasChildren
}

func NewElseNode() INode {
	return &ElseNode{NewHasChildren("ELSE")}
}

func (node ElseNode) ToString() string {
	return fmt.Sprintf("<ELSE>")
}

func (node *ElseNode) Type() int {
	return NODE_TYPE_ELSE
}

func (node *ElseNode) Render(context map[string]interface{}) string {
	s := ""
	for _, child := range node.Children {
		s += (*child).Render(context)
	}
	return s
}
