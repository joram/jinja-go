package nodes

import "fmt"

type EndIfNode struct {
	HasNoChildren
}

func NewEndIfNode() INode {
	return &EndIfNode{NewHasNoChildren("ENDIF")}
}

func (node EndIfNode) ToString() string {
	return fmt.Sprintf("</IFELSE>")
}

func (node *EndIfNode) Type() int {
	return NODE_TYPE_ENDIF
}
