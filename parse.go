package jinja_go

import (
	"github.com/joram/jinja-go/nodes"
	"io"
	"strings"
	"sync"
)

func GetNodes(content string, c chan nodes.INode, wg *sync.WaitGroup, config Configuration) {
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

func GetNode(body string, config Configuration) (nodes.INode, int64, error) {
	followingNodeStart, _ := GetNodeStartingIndex(body, config)
	if followingNodeStart > 0 {
		end := followingNodeStart
		return nodes.NewTextNode(body[:end]), end, nil
	}

	end, s := GetNodeEndingIndex(body, config)
	end += int64(len(s))
	if end == -1 {
		if len(body) > 0 {
			return nodes.NewTextNode(body), int64(len(body)), io.EOF
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
		if blockType == "if" {
			return nodes.NewIfNode(blockString), end, nil
		}
		if blockType == "else" {
			return nodes.NewElseNode(blockString), end, nil
		}
		if blockType == "endif" {
			return nodes.NewEndIfNode(blockString), end, nil
		}
		if blockType == "for" {
			return nodes.NewForNode(blockString), end, nil
		}
		if blockType == "endfor" {
			return nodes.NewEndForNode(blockString), end, nil
		}
		return nil, end, nil
	}

	if s == config.VariableEndString {
		return nodes.NewVariableNode(body[:end]), end, nil
	}

	if s == config.CommentEndString {
		return nodes.NewCommentNode(body[:end]), end, nil
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
