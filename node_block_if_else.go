package jinja_go

import "fmt"

type IfElseNode struct {
	ifNode   *INode
	elseNode *INode
	endNode  *INode
	closed   bool
}

func (node *IfElseNode) append(child *INode) {
	if node.ifNode == nil {
		node.ifNode = child
		return
	}
	if node.elseNode == nil {
		node.elseNode = child
		return
	}
	if node.endNode == nil {
		node.endNode = child
		return
	}

}

func (node *IfElseNode) close() {
	node.closed = true
}

func (node *IfElseNode) isClosed() bool {
	return node.endNode != nil
}

func (node *IfElseNode) toString() string {
	return "<IFELSE>"
}

type IfNode struct {
	children []*INode
	closed   bool
}

func (node *IfNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node *IfNode) close() {
	node.closed = true
}

func (node *IfNode) isClosed() bool {
	return node.closed
}

func (node *IfNode) toString() string {
	return fmt.Sprintf("<IF>")
}

type ElseNode struct {
	children []*INode
	closed   bool
}

func (node ElseNode) append(child *INode) {
	node.children = append(node.children, child)
}

func (node ElseNode) toString() string {
	return fmt.Sprintf("<ELSE>")
}
func (node *ElseNode) close() {
	node.closed = true
}

func (node ElseNode) isClosed() bool {
	return node.closed
}

type EndIfNode struct {
	closed bool
}

func (node EndIfNode) append(child *INode) {}

func (node EndIfNode) toString() string {
	return fmt.Sprintf("</IFELSE>")
}
func (node *EndIfNode) close() {
	node.closed = true
}
func (node EndIfNode) isClosed() bool {
	return true
}
