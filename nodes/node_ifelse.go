package nodes

type IfElseNode struct {
	HasChildren
	IfNode   INode
	ElseNode INode
	EndNode  INode
}

func NewIfElseNode() IfElseNode {
	return IfElseNode{NewHasChildren("IFELSE"), nil, nil, nil}
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
	return "" // TODO
}
