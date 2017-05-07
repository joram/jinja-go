package jinja_go

import "fmt"

type CommentNode struct {
	children []*INode
	body     string
}

func (node *CommentNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node *CommentNode) close() {}

func (node *CommentNode) isClosed() bool {
	return true
}

func (node *CommentNode) toString() string {
	return fmt.Sprintf("<COMMENT>")
}
