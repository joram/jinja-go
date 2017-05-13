package nodes

import "fmt"

type EndForNode struct {
	HasNoChildren
}

func NewEndForNode() INode {
	return &EndForNode{NewHasNoChildren("ENDFOR")}
}

func (node EndForNode) ToString() string {
	return fmt.Sprintf("</FOR>")
}

func (node *EndForNode) Type() int {
	return NODE_TYPE_ENDFOR
}
