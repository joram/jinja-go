package jinja

type Parser struct {
	tokenizer Tokenizer
	root      Node
	headNode  *Node
}

type Node struct {
	token    Token
	parent   *Node
	children []*Node
}

func NewParser() Parser {
	return Parser{}
}

func (parser Parser) Parse(s string) {
	parser.tokenizer = NewTokenizer(s)
	parser.root = Node{}
	parser.headNode = &parser.root
	for token := range parser.tokenizer.C {
		if token.Type == EXTERNAL_CONTENT {
			thisNode := Node{
				token:  token,
				parent: parser.headNode,
			}
			parser.headNode.children = append(parser.headNode.children, &thisNode)
		}
		if token.Type == EXPRESSION_BEGIN {
			parser.parseExpression(token)
		}
	}
}

func (parser Parser) parseExpression(beginToken Token) {

}
