package jinja_go

import (
	"encoding/json"
	"fmt"
)

const (
	NODE_TYPE_ROOT     = iota
	NODE_TYPE_IFELSE   = iota
	NODE_TYPE_IF       = iota
	NODE_TYPE_ELSE     = iota
	NODE_TYPE_ENDIF    = iota
	NODE_TYPE_FOR      = iota
	NODE_TYPE_ENDFOR   = iota
	NODE_TYPE_COMMENT  = iota
	NODE_TYPE_TEXT     = iota
	NODE_TYPE_VARIABLE = iota
)

var NODE_TYPES = []int{
	NODE_TYPE_ROOT,
	NODE_TYPE_IFELSE,
	NODE_TYPE_IF,
	NODE_TYPE_ELSE,
	NODE_TYPE_ENDIF,
	NODE_TYPE_FOR,
	NODE_TYPE_ENDFOR,
	NODE_TYPE_COMMENT,
	NODE_TYPE_TEXT,
	NODE_TYPE_VARIABLE,
}

var HAS_CHILDREN = map[int]bool{
	NODE_TYPE_ROOT:     true,
	NODE_TYPE_IFELSE:   true,
	NODE_TYPE_IF:       true,
	NODE_TYPE_ELSE:     true,
	NODE_TYPE_ENDIF:    false,
	NODE_TYPE_FOR:      true,
	NODE_TYPE_ENDFOR:   false,
	NODE_TYPE_COMMENT:  false,
	NODE_TYPE_TEXT:     false,
	NODE_TYPE_VARIABLE: false,
}

var NODE_NAME = map[int]string{
	NODE_TYPE_ROOT:     "ROOT",
	NODE_TYPE_IFELSE:   "IFELSE",
	NODE_TYPE_IF:       "IF",
	NODE_TYPE_ELSE:     "ELSE",
	NODE_TYPE_ENDIF:    "ENDIF",
	NODE_TYPE_FOR:      "FOR",
	NODE_TYPE_ENDFOR:   "ENDFOR",
	NODE_TYPE_COMMENT:  "COMMENT",
	NODE_TYPE_TEXT:     "TEXT",
	NODE_TYPE_VARIABLE: "VAR",
}

var NODE_ENDS = map[int]bool{
	NODE_TYPE_ROOT:     false,
	NODE_TYPE_IFELSE:   false,
	NODE_TYPE_IF:       false,
	NODE_TYPE_ELSE:     false,
	NODE_TYPE_ENDIF:    true,
	NODE_TYPE_FOR:      false,
	NODE_TYPE_ENDFOR:   true,
	NODE_TYPE_COMMENT:  false,
	NODE_TYPE_TEXT:     false,
	NODE_TYPE_VARIABLE: false,
}

type Node struct {
	//Render(map[string]interface{}) string
	Type     int
	Children []*Node
	Content  string
}

func (n *Node) MarshalJSON() ([]byte, error) {
	if n.HasChildren() {
		return json.Marshal(struct {
			Type     string
			Children []*Node
		}{
			Type:     n.ToString(),
			Children: n.Children,
		})
	}

	return json.Marshal(struct {
		Type string
	}{
		Type: n.ToString(),
	})

}

func NewNode(nodeType int, content string) Node {
	ValidateNodeType(nodeType)
	return Node{nodeType, []*Node{}, content}
}

func ValidateNodeType(nodeType int) {
	for _, validNodeType := range NODE_TYPES {
		if validNodeType == nodeType {
			return
		}
	}
	panic(fmt.Sprintf("%d is not a valid node type", nodeType))
}

func (node *Node) IsEnd() bool {
	ValidateNodeType(node.Type)
	return NODE_ENDS[node.Type]
}

func (node *Node) ToString() string {
	ValidateNodeType(node.Type)
	return NODE_NAME[node.Type]
}

func (node *Node) HasChildren() bool {
	ValidateNodeType(node.Type)
	return HAS_CHILDREN[node.Type]
}

func (node *Node) Append(child *Node) {
	if !node.HasChildren() {
		panic("appending to a node that doesn't allow children")
	}
	node.Children = append(node.Children, child)
}

func (node *Node) Render(context map[string]interface{}) string {
	switch node.Type {
	case NODE_TYPE_IFELSE:
		return node.Children[0].Render(context)
	}

	if node.HasChildren() {
		s := ""
		for _, child := range node.Children {
			s += child.Render(context)
		}
		return s
	}

	switch node.Type {
	case NODE_TYPE_TEXT:
		return node.Content
	case NODE_TYPE_VARIABLE:
		return "" // TODO parse instead
	default:
		return ""
	}
}
