package jinja

type Parser struct {
	tokens    chan Token
	headToken Token
	tokenizer Tokenizer
	nodes     []NodeInter
}

func (parser *Parser) Parse(template string) {
	parser.tokens = make(chan Token)
	parser.tokenizer = NewTokenizer(template)
	go parser.tokenizer.GetTokens(parser.tokens)
	parser.headToken = <-parser.tokens // prime look ahead cache

	eof := false
	for !eof {
		n := parser.getNode()
		eof = n == nil
		if !eof {
			parser.nodes = append(parser.nodes, n)
		}
	}
}

func (parser *Parser) getToken() Token {
	currHead := parser.headToken
	next := <-parser.tokens
	for next.Definition.tokenType == SPACE {
		return parser.getToken()
	}
	parser.headToken = next
	return currHead
}

func (parser *Parser) peekToken() Token {
	return parser.headToken
}

func (parser *Parser) getExpressionNode() NodeInter {

	t := parser.getToken()
	if parser.peekToken().Definition.closesBlock {
		if t.Definition.tokenType == VARIABLE {
			return VariableNode{t.Value}
		}
	}

	if parser.peekToken().Definition.operator {
		opToken := parser.getToken()
		return ExpressionNode{
			leftChild:  VariableNode{t.Value},
			rightChild: parser.getExpressionNode(), // woo recursion. this is gonna bite me in the ass.
			operation:  opToken.Definition.tokenType,
		}
	}

	//	todo: more than single node expressions
	return nil
}

func (parser *Parser) getNode() NodeInter {
	firstToken := parser.getToken()
	if firstToken.Definition.tokenType == EOF {
		return nil
	}

	if firstToken.Definition.tokenType == TEXT {
		return ValueNode{firstToken.Value}
	}

	if firstToken.Definition.tokenType == OUTPUT_BEGIN {
		expression := parser.getExpressionNode()
		closeToken := parser.getToken() // assert they are the same types
		if closeToken.Definition.tokenType != OUTPUT_END {
			panic("didn't find a proper closing to the output")
		}
		return expression
	}

	return nil
}
