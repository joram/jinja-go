package nodes

import "fmt"

type IfNode struct {
	HasChildren
}

func (node *IfNode) ToString() string {
	return fmt.Sprintf("<IF>")
}

func (node *IfNode) Type() int {
	return NODE_TYPE_IF
}
