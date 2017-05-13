package jinja_go

import (
	"encoding/json"
	"github.com/joram/jinja-go/nodes"
	"sync"
)

type Template struct {
	Config   Configuration
	Content  string
	rootNode nodes.RootNode
	stack    []nodes.INode
}

func NewTemplate() Template {
	return Template{
		NewDefaultConfig(),
		"",
		nodes.NewRootNode(),
		[]nodes.INode{},
	}
}

func (template *Template) topNode() nodes.INode {
	return template.stack[len(template.stack)-1]
}

func (template *Template) popNode() nodes.INode {
	if len(template.stack) == 0 {
		return nil
	}
	node := template.topNode()
	template.stack = template.stack[:len(template.stack)-1]
	return node
}

func (template *Template) addNode(node nodes.INode) {
	template.topNode().Append(&node)
	template.stack = append(template.stack, node)
}

func (template *Template) Compile(content string) error {
	template.Content = content
	template.stack = []nodes.INode{&template.rootNode}
	ifElses := []nodes.IfElseNode{}

	c := make(chan nodes.INode)
	var wg sync.WaitGroup
	wg.Add(1)
	go GetNodes(content, c, &wg, template.Config)
	for node := range c {
		if node.Type() == nodes.NODE_TYPE_IF {
			ifElseNode := nodes.NewIfElseNode()
			ifElseNode.IfNode = node
			template.addNode(&ifElseNode)
			template.addNode(node)
			ifElses = append(ifElses, ifElseNode)
			continue
		}

		if node.Type() == nodes.NODE_TYPE_ELSE {
			poppedType := -1
			for poppedType != nodes.NODE_TYPE_IF {
				poppedType = template.popNode().Type()
			}
			ifElses[len(ifElses)-1].ElseNode = node
			template.addNode(node)
			continue
		}

		if node.Type() == nodes.NODE_TYPE_ENDIF {
			ifElses[len(ifElses)-1].EndNode = node

			poppedType := -1
			for poppedType != nodes.NODE_TYPE_IFLSE {
				poppedType = template.popNode().Type()
			}
			template.addNode(node)
			template.popNode()
			continue
		}

		if node.Type() == nodes.NODE_TYPE_FOR {
			template.addNode(node)
		}

		if node.Type() == nodes.NODE_TYPE_ENDFOR {
			poppedType := -1
			for poppedType != nodes.NODE_TYPE_FOR {
				poppedType = template.popNode().Type()
			}
			template.addNode(node)
			template.popNode()
			continue
		}

		if node.Type() == nodes.NODE_TYPE_TEXT {
			template.addNode(node)
			template.popNode()
			continue
		}
		if node.Type() == nodes.NODE_TYPE_COMMENT {
			template.addNode(node)
			template.popNode()
			continue
		}
	}
	wg.Wait()
	return nil
}

func (template Template) JSONTree() (string, error) {
	b, err := json.Marshal(template.rootNode)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
