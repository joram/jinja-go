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
	Render(map[string]interface{}) string
}

type HasChildren struct {
	Type     string
	Children []*INode
}

func NewHasChildren(t string) HasChildren {
	return HasChildren{t, []*INode{}}
}

func (node *HasChildren) Append(child *INode) {
	node.Children = append(node.Children, child)
}

func (node HasChildren) GetChildren() []*INode {
	return node.Children
}

type HasNoChildren struct {
	Type string
}

func NewHasNoChildren(t string) HasNoChildren {
	return HasNoChildren{t}
}

func (node HasNoChildren) GetChildren() []*INode {
	return nil
}

func (node HasNoChildren) Append(child *INode) {
	panic("adding child to node with no Children")
}
