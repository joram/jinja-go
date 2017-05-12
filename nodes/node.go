package nodes

const (
	NODE_TYPE_ROOT     = iota
	NODE_TYPE_IFLSE    = iota
	NODE_TYPE_IF       = iota
	NODE_TYPE_ELSE     = iota
	NODE_TYPE_ENDIF    = iota
	NODE_TYPE_FOR      = iota
	NODE_TYPE_ENDFOR   = iota
	NODE_TYPE_COMMENT  = iota
	NODE_TYPE_TEXT     = iota
	NODE_TYPE_VARIABLE = iota
)

type INode interface {
	Type() int
	ToString() string
	GetChildren() []*INode
	Append(child *INode)
}

type HasChildren struct {
	children []*INode
}

func (node HasChildren) Append(child *INode) {
	node.children = append(node.children, child)
}

func (node HasChildren) GetChildren() []*INode {
	return node.children
}

type HasNoChildren struct{}

func (node HasNoChildren) GetChildren() []*INode {
	return nil
}

func (node HasNoChildren) Append(child *INode) {
	panic("adding child to node with no children")
}
