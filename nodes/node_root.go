package nodes

type RootNode struct {
	HasChildren
}

func (node *RootNode) ToString() string {
	return "<ROOT>"
}

func (node *RootNode) Type() int {
	return NODE_TYPE_ROOT
}
