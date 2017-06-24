package nodes

import (
	"fmt"
	"github.com/joram/jinja-go/evaluate"
	"strings"
)

type IfElseNode struct {
	HasChildren
	IfNode   INode
	ElseNode INode
	EndNode  INode
}

func NewIfElseNode(content string) IfElseNode {
	return IfElseNode{NewHasChildren("IFELSE", content), nil, nil, nil}
}

func (node *IfElseNode) ToString() string {
	return "<IFELSE>"
}

func (node *IfElseNode) Type() int {
	return NODE_TYPE_IFLSE
}

func (node *IfElseNode) GetChildren() []*INode {
	return []*INode{
		&node.IfNode,
		&node.ElseNode,
		&node.EndNode,
	}
}

func (node *IfElseNode) Append(child *INode) {}

func (node *IfElseNode) Render(context map[string]interface{}) string {
	s := node.IfNode.GetContent()
	s = strings.Replace(s, "if", "", 1)
	isTrue, err := evaluate.Evaluate(s, context)

	if err != nil {
		panic(err)
	}
	if isTrue {
		return node.IfNode.Render(context)
	}
	if node.ElseNode != nil {
		fmt.Printf("rendering else node: %v\n", node.ElseNode.Render(context))
		return node.ElseNode.Render(context)
	}
	println("no else node")
	return ""
}
