package nodes

import "fmt"

type EndForNode struct {
	HasNoChildren
}

func NewEndForNode(content string) INode {
	return &EndForNode{NewHasNoChildren("ENDFOR", content)}
}

func (node EndForNode) ToString() string {
	return fmt.Sprintf("</FOR>")
}

func (node *EndForNode) Type() int {
	return NODE_TYPE_ENDFOR
}

func (node *EndForNode) Render(context map[string]interface{}) string {
	return ""
}
