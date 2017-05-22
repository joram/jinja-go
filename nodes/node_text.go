package nodes

import "fmt"

type TextNode struct {
	HasNoChildren
	text string
}

func NewTextNode(content string) INode {
	node := TextNode{NewHasNoChildren("TEXT"), content}
	return &node
}

func (node TextNode) ToString() string {
	return fmt.Sprintf("<STR>%s</STR>", node.text)
}

func (node *TextNode) Type() int {
	return NODE_TYPE_TEXT
}

func (node *TextNode) Render(context map[string]interface{}) string {
	return node.text
}
