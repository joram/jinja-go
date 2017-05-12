package nodes

import "fmt"

type ElseNode struct {
	HasChildren
}

func (node ElseNode) ToString() string {
	return fmt.Sprintf("<ELSE>")
}

func (node *ElseNode) Type() int {
	return NODE_TYPE_ELSE
}
