package jinja_go

import "fmt"

type VariableNode struct {
	children []*INode
	body     string
}

func (node *VariableNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node *VariableNode) close() {}

func (node *VariableNode) isClosed() bool {
	return true
}

func (node *VariableNode) toString() string {
	return fmt.Sprintf("<VAR>")
}
