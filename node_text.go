package jinja_go

import "fmt"

type TextNode struct {
	children []*INode
	text     string
}

func (node TextNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node TextNode) close() {
}

func (node TextNode) isClosed() bool {
	return true
}

func (node TextNode) toString() string {
	return fmt.Sprintf("<STR>%s</STR>", node.text)
}
