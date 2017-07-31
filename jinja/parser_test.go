package jinja

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	testStrings := []string{
		//"before {{ foo }} middle {{ bar }} after",
		"before {{ 1 + 2 }} middle {{ bar }} after",
		//"before {% if 'a string with spaces in it' == foo %} middle {{ bar }} after",
	}
	for _, s := range testStrings {

		fmt.Printf("parsing: %v\n", s)

		parser := Parser{}
		parser.Parse(s)

		nodes := []NodeInter{}
		for _, node := range parser.nodes {
			nodes = append(nodes, node)
		}
		fmt.Printf("Nodes Parsed: %s\n", nodes)
	}
	t.Errorf("")
}
