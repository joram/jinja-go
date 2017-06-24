package nodes

import "fmt"

type EndIfNode struct {
	HasNoChildren
}

func NewEndIfNode(content string) INode {
	return &EndIfNode{NewHasNoChildren("ENDIF", content)}
}

func (node EndIfNode) ToString() string {
	return fmt.Sprintf("</IFELSE>")
}

func (node *EndIfNode) Type() int {
	return NODE_TYPE_ENDIF
}

func (node *EndIfNode) Render(context map[string]interface{}) string {
	return ""
}
