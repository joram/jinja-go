package jinja_go

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type Template struct {
	Config   Configuration
	Content  string
	rootNode RootNode
	stack    []INode
}

func (template *Template) topNode() INode {
	return template.stack[len(template.stack)-1]
}

func (template *Template) popNode() INode {
	node := template.topNode()
	template.topNode().close()
	template.stack = template.stack[:len(template.stack)-1]
	return node
}
func (template *Template) cleanStack() {
	for template.topNode().isClosed() {
		template.popNode()
	}
}

func (template *Template) debugPrint(node INode) {
	tabs := fmt.Sprintf("%d", len(template.stack)-1)
	for i := 0; i < len(template.stack); i++ {
		tabs += "  "
	}
	fmt.Printf("%s%s\n", tabs, node.toString())
}

func (template *Template) addNode(node INode) {
	template.cleanStack()
	if template.topNode().toString() == "<IF>" && node.toString() == "<ELSE>" {
		template.popNode()
	}
	if node.toString() == "</IFELSE>" {
		closedNode := template.popNode()
		for closedNode.toString() != "<IFELSE>" {
			closedNode = template.popNode()
		}
	}
	template.debugPrint(node)

	template.topNode().append(&node)
	template.cleanStack()
	template.stack = append(template.stack, node)

}

func (template *Template) Compile(content string) error {
	template.Content = content
	template.stack = []INode{&template.rootNode}

	var err error
	for err == nil {
		nodes, length, err := template.GetNode(content)
		for _, node := range nodes {
			template.addNode(node)
		}

		if err != nil {
			return err
		}
		content = content[length:]
		time.Sleep(time.Second)
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func (template *Template) GetNode(body string) ([]INode, int64, error) {
	followingNodeStart, _ := template.GetNodeStartingIndex(body)
	if followingNodeStart > 0 {
		end := followingNodeStart
		node := TextNode{
			text: body[:end],
		}
		return []INode{node}, end, nil
	}

	end, s := template.GetNodeEndingIndex(body)
	end += int64(len(s))
	if end == -1 {
		if len(body) > 0 {
			return []INode{&TextNode{text: body}}, int64(len(body)), io.EOF
		}
		return nil, 0, io.EOF
	}

	if s == template.Config.BlockEndString {
		blockString := body[:end]
		blockType := strings.Replace(blockString, template.Config.BlockStartString, "", 1)
		blockType = strings.TrimLeft(blockType, " ")
		blockType = blockType[:strings.Index(blockType, " ")]
		nodes := []INode{}
		if blockType == "if" {
			ifelseNode := IfElseNode{}
			ifNode := IfNode{closed: false}
			nodes = []INode{
				&ifelseNode,
				&ifNode,
			}
		}
		if blockType == "else" {
			nodes = []INode{&ElseNode{}}
		}
		if blockType == "endif" {
			nodes = []INode{&EndIfNode{}}
		}
		return nodes, end, nil
	}
	return []INode{}, end, nil
}

func (template *Template) GetNodeStartingIndex(s string) (int64, string) {
	return template.GetNextOccurance(s, []string{
		template.Config.CommentStartString,
		template.Config.VariableStartString,
		template.Config.BlockStartString,
	})
}

func (template *Template) GetNodeEndingIndex(s string) (int64, string) {
	return template.GetNextOccurance(s, []string{
		template.Config.CommentEndString,
		template.Config.VariableEndString,
		template.Config.BlockEndString,
	})
}

func (template *Template) GetNextOccurance(s string, options []string) (int64, string) {
	min := -1
	minOption := ""
	for _, option := range options {
		index := strings.Index(s, option)
		if index > -1 && (index < min || min == -1) {
			min = index
			minOption = option
		}
	}
	return int64(min), minOption
}
