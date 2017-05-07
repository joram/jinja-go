package jinja_go

import "fmt"

type RootNode struct {
	children []*INode
}

func (node *RootNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node *RootNode) close() {}

func (node *RootNode) isClosed() bool {
	return false
}

func (node *RootNode) toString() string {
	return fmt.Sprintf("<ROOT>")
}
