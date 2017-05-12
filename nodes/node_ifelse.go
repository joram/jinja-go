package nodes

type IfElseNode struct {
	IfNode   INode
	ElseNode INode
	EndNode  INode
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
