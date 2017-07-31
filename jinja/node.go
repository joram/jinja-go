package jinja

import "fmt"

type NodeInter interface {
	String() string
}

type Node struct {
	tokens   []Token
	complete bool
}

func (n Node) isComplete() bool {
	return n.complete
}

func (n Node) String() string {
	s := ""
	for _, t := range n.tokens {
		s += fmt.Sprintf("'%+v'", t.String())
	}
	return s
}

type ExpressionNode struct {
	leftChild  NodeInter
	rightChild NodeInter
	operation  int
}

func (n ExpressionNode) isComplete() bool { return true }

func (n ExpressionNode) String() string { return "expression" }

type VariableNode struct {
	variableName string
}

func (n VariableNode) isComplete() bool { return true }

func (n VariableNode) String() string { return "<VAR: " + n.variableName + ">" }

type ValueNode struct {
	value interface{}
}

func (n ValueNode) isComplete() bool { return true }

func (n ValueNode) String() string { return fmt.Sprintf("<VAL: %v>", n.value) }
