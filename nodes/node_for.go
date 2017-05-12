package nodes

import "fmt"

type ForNode struct {
	HasChildren
}

func (node *ForNode) ToString() string {
	return fmt.Sprintf("<FOR>")
}

func (node *ForNode) Type() int {
	return NODE_TYPE_FOR
}
