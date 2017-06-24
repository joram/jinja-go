package evaluate

import (
	"errors"
	_ "errors"
	"fmt"
	"regexp"
)

// regular expressions
var (
	whitespace_re = regexp.MustCompile(`\s+`)
	string_re1    = regexp.MustCompile(`'([^'\\]*(?:\\.[^'\\]*)*)'`)
	string_re2    = regexp.MustCompile(`"([^"\\]*(?:\\.[^"\\]*)*)"`)
	integer_re    = regexp.MustCompile(`\d+`)
	float_re      = regexp.MustCompile(`([+-]?\d+\.\d+)`)
	newline_re    = regexp.MustCompile(`(\r\n|\r|\n)`)
	variable_re   = regexp.MustCompile(`([a-zA-Z0-9_]+)`)
)

// Token definitions
var (
	TOKEN_ADD                 = "add"
	TOKEN_ASSIGN              = "assign"
	TOKEN_COLON               = "colon"
	TOKEN_COMMA               = "comma"
	TOKEN_DIV                 = "div"
	TOKEN_DOT                 = "dot"
	TOKEN_EQ                  = "eq"
	TOKEN_FLOORDIV            = "floordiv"
	TOKEN_GT                  = "gt"
	TOKEN_GTEQ                = "gteq"
	TOKEN_LBRACE              = "lbrace"
	TOKEN_LBRACKET            = "lbracket"
	TOKEN_LPAREN              = "lparen"
	TOKEN_LT                  = "lt"
	TOKEN_LTEQ                = "lteq"
	TOKEN_MOD                 = "mod"
	TOKEN_MUL                 = "mul"
	TOKEN_NE                  = "ne"
	TOKEN_PIPE                = "pipe"
	TOKEN_POW                 = "pow"
	TOKEN_RBRACE              = "rbrace"
	TOKEN_RBRACKET            = "rbracket"
	TOKEN_RPAREN              = "rparen"
	TOKEN_SEMICOLON           = "semicolon"
	TOKEN_SUB                 = "sub"
	TOKEN_TILDE               = "tilde"
	TOKEN_WHITESPACE          = "whitespace"
	TOKEN_FLOAT               = "float"
	TOKEN_INTEGER             = "integer"
	TOKEN_NAME                = "name"
	TOKEN_STRING              = "string"
	TOKEN_OPERATOR            = "operator"
	TOKEN_BLOCK_BEGIN         = "block_begin"
	TOKEN_BLOCK_END           = "block_end"
	TOKEN_VARIABLE            = "variable"
	TOKEN_RAW_BEGIN           = "raw_begin"
	TOKEN_RAW_END             = "raw_end"
	TOKEN_COMMENT_BEGIN       = "comment_begin"
	TOKEN_COMMENT_END         = "comment_end"
	TOKEN_COMMENT             = "comment"
	TOKEN_LINESTATEMENT_BEGIN = "linestatement_begin"
	TOKEN_LINESTATEMENT_END   = "linestatement_end"
	TOKEN_LINECOMMENT_BEGIN   = "linecomment_begin"
	TOKEN_LINECOMMENT_END     = "linecomment_end"
	TOKEN_LINECOMMENT         = "linecomment"
	TOKEN_DATA                = "data"
	TOKEN_INITIAL             = "initial"
	TOKEN_EOF                 = "eof"
	TOKEN_NEWLINE             = "newline"
)

var operators = map[string]string{
	"==": TOKEN_EQ,
	"/+": TOKEN_ADD,
	"-":  TOKEN_SUB,
	"/":  TOKEN_DIV,
	"//": TOKEN_FLOORDIV,
	"*":  TOKEN_MUL,
	"%":  TOKEN_MOD,
	"**": TOKEN_POW,
	"~":  TOKEN_TILDE,
	"[":  TOKEN_LBRACKET,
	"]":  TOKEN_RBRACKET,
	"(":  TOKEN_LPAREN,
	")":  TOKEN_RPAREN,
	"{":  TOKEN_LBRACE,
	"}":  TOKEN_RBRACE,
	"!=": TOKEN_NE,
	">":  TOKEN_GT,
	">=": TOKEN_GTEQ,
	"<":  TOKEN_LT,
	"<=": TOKEN_LTEQ,
	"=":  TOKEN_ASSIGN,
	".":  TOKEN_DOT,
	":":  TOKEN_COLON,
	"|":  TOKEN_PIPE,
	",":  TOKEN_COMMA,
	";":  TOKEN_SEMICOLON,
}

type TokenType struct {
	tokenType string
	regex     *regexp.Regexp
}
type Token struct {
	index      int
	tokenType  string
	tokenValue string
	Children   []*Token
}

func (token *Token) Ignore() bool {
	for _, ignorableType := range []string{TOKEN_WHITESPACE, TOKEN_INITIAL, TOKEN_EOF} {
		if token.tokenType == ignorableType {
			return true
		}
	}
	return false
}

func (token *Token) IsTrue(context map[string]interface{}) bool {
	if token.tokenType == TOKEN_EQ {
		if len(token.Children) != 2 {
			panic(fmt.Sprintf("EQ node has %v children", len(token.Children)))
		}
		leftVal := token.Children[0].Value(context)
		rightVal := token.Children[1].Value(context)
		fmt.Printf("%v ?= %v\n", token.Children[0].tokenValue, token.Children[1].tokenValue)
		fmt.Printf("%v ?= %v\n", leftVal, rightVal)
		return leftVal == rightVal
	}

	panic(errors.New(fmt.Sprintf("IsTrue called on %v node", token.tokenType)))
}

func (token *Token) Value(context map[string]interface{}) interface{} {
	return context[token.tokenValue]
}

func allTokenTypes() []TokenType {
	tokenTypes := []TokenType{
		{TOKEN_WHITESPACE, whitespace_re},
		{TOKEN_STRING, string_re1},
		{TOKEN_STRING, string_re2},
		{TOKEN_INTEGER, integer_re},
		{TOKEN_FLOAT, float_re},
		{TOKEN_NEWLINE, newline_re},
	}
	for val, tokenType := range operators {
		newToken := TokenType{tokenType, regexp.MustCompile(regexp.QuoteMeta(val))}
		tokenTypes = append(tokenTypes, newToken)
	}
	return tokenTypes
}

func NextToken(curr Token, s string) Token {
	index := curr.index + len(curr.tokenValue)

	if len(s) == 0 {
		return Token{index, TOKEN_EOF, "", []*Token{}}
	}

	for _, tokenType := range allTokenTypes() {
		indices := tokenType.regex.FindIndex([]byte(s))
		for _, index := range indices {
			if index == 0 {
				tokenValue := tokenType.regex.FindString(s)
				return Token{index, tokenType.tokenType, tokenValue, []*Token{}}

			}
		}
	}

	tokenValue := variable_re.FindString(s)
	return Token{index, TOKEN_VARIABLE, tokenValue, []*Token{}}
}

func BuildTree(tokens []Token) *Token {
	topNode := &(tokens[0])
	leafNode := topNode
	tokens = tokens[1:]
	for {
		if len(tokens) == 0 {
			break
		}

		// pop
		currNode := &(tokens[0])
		tokens = tokens[1:]

		if currNode.tokenType == TOKEN_EQ {
			currNode.Children = append(currNode.Children, topNode)
			topNode = currNode
			leafNode = currNode
			continue
		}

		if currNode.tokenType == TOKEN_VARIABLE {
			leafNode.Children = append(leafNode.Children, currNode)
			leafNode = currNode
		}

	}
	return topNode
}

func Evaluate(s string, context map[string]interface{}) (bool, error) {
	tokens := Tokenize(s)
	parentNode := BuildTree(tokens) // TODO: store tokenized strings on nodes

	fmt.Printf("%v\n", tokens)
	fmt.Printf("evaluating %v\n", parentNode.tokenType)
	if parentNode.tokenType == TOKEN_EQ {
		return parentNode.IsTrue(context), nil
	}

	panic(fmt.Sprintf("evaluation of non-evaluable token type %v", parentNode.tokenType))
	return false, nil
}

func Tokenize(s string) []Token {
	tokens := []Token{}
	prev := Token{0, TOKEN_INITIAL, "", []*Token{}}

	itter := 0
	for {
		curr := NextToken(prev, s)
		curr.index += prev.index
		if !curr.Ignore() {
			tokens = append(tokens, curr)
		}

		if curr.tokenType == TOKEN_EOF {
			break
		}
		itter += 1
		if itter > 10 {
			panic(errors.New("10 iteratons of nothing"))
			break
		}

		s = s[len(curr.tokenValue):]
		prev = curr
	}

	return tokens
}
