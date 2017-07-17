package jinja_go

import (
	"io"
	"strings"
	"sync"
)

var nodeTypes = map[string]int{
	"if":     NODE_TYPE_IF,
	"else":   NODE_TYPE_ELSE,
	"endif":  NODE_TYPE_ENDIF,
	"for":    NODE_TYPE_FOR,
	"endfor": NODE_TYPE_ENDFOR,
}

func GetNodes(content string, c chan *Node, wg *sync.WaitGroup, config Configuration) {
	defer wg.Done()
	var err error
	for err == nil {
		node, length, err := GetNode(content, config)
		content = content[length:]
		c <- node
		if err != nil {
			break
		}
	}
	if err != io.EOF && err != nil {
		panic(err)
	}
	close(c)
}

func GetNode(body string, config Configuration) (*Node, int64, error) {
	followingNodeStart, _ := GetNodeStartingIndex(body, config)
	if followingNodeStart > 0 {
		end := followingNodeStart
		node := NewNode(NODE_TYPE_TEXT, body[:end])
		return &node, end, nil
	}

	end, s := GetNodeEndingIndex(body, config)
	end += int64(len(s))
	if end == -1 {
		if len(body) > 0 {
			node := NewNode(NODE_TYPE_TEXT, body)
			return &node, int64(len(body)), io.EOF
		}
		return nil, 0, io.EOF
	}

	if s == config.BlockEndString {
		blockString := body[:end]
		blockType := strings.Replace(blockString, config.BlockStartString, "", 1)
		blockType = strings.TrimLeft(blockType, " ")
		blockType = blockType[:strings.Index(blockType, " ")]
		blockString = strings.TrimLeft(blockString, config.BlockStartString)
		blockString = strings.TrimRight(blockString, config.BlockEndString)

		node := NewNode(nodeTypes[blockType], blockString)
		return &node, end, nil
	}

	if s == config.VariableEndString {
		node := NewNode(NODE_TYPE_VARIABLE, body[:end])
		return &node, end, nil
	}

	if s == config.CommentEndString {
		node := NewNode(NODE_TYPE_COMMENT, body[:end])
		return &node, end, nil
	}

	return nil, end, nil
}

func GetNodeStartingIndex(s string, config Configuration) (int64, string) {
	return GetNextOccurance(s, []string{
		config.CommentStartString,
		config.VariableStartString,
		config.BlockStartString,
	})
}

func GetNodeEndingIndex(s string, config Configuration) (int64, string) {
	return GetNextOccurance(s, []string{
		config.CommentEndString,
		config.VariableEndString,
		config.BlockEndString,
	})
}

func GetNextOccurance(s string, options []string) (int64, string) {
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
