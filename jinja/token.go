package jinja

import "fmt"

type Token struct {
	Type  string
	Value string
	line  int
}

// Tokens definitions
const (
	ADD              = "+"
	ASSIGN           = "=="
	COLON            = ":"
	COMMA            = ","
	DIV              = "/"
	DOT              = "."
	EQ               = "="
	FLOORDIV         = "//"
	GT               = ">"
	GTEQ             = ">="
	LBRACE           = "{"
	LBRACKET         = "["
	LPAREN           = "("
	LT               = "<"
	LTEQ             = "<="
	MOD              = "%"
	MUL              = "*"
	NE               = "!="
	PIPE             = "|"
	POW              = "^"
	RBRACE           = "}"
	RBRACKET         = "]"
	RPAREN           = ")"
	SEMICOLON        = ";"
	SUB              = "-"
	FLOAT            = "FLOAT"
	INTEGER          = "INTEGER"
	IDENTIFIER       = "IDENTIFIER"
	EXTERNAL_CONTENT = "EXTERNAL_CONTENT"
	ILLEGAL          = "ILLEGAL"
	BLOCK_BEGIN
	BLOCK_END
	RAW_BEGIN
	RAW_END
	DATA
	INITIAL
	EOF = "EOF"
	NEWLINE

	LINESTATEMENT_BEGIN
	LINESTATEMENT_END
	LINECOMMENT_BEGIN
	LINECOMMENT_END
	LINECOMMENT

	COMMENT_BEGIN    = "{#"
	COMMENT_END      = "#}"
	EXPRESSION_BEGIN = "{{"
	EXPRESSION_END   = "}}"
)

var tokenNames = map[string]string{
	EQ:               "EQ",
	ADD:              "ADD",
	SUB:              "SUB",
	FLOORDIV:         "FLOORDIV",
	DIV:              "DIV",
	MUL:              "MUL",
	MOD:              "MOD",
	POW:              "POW",
	LBRACKET:         "LBRACKET",
	RBRACKET:         "RBRACKET",
	LPAREN:           "LPAREN",
	RPAREN:           "RPAREN",
	NE:               "NE",
	GT:               "GT",
	GTEQ:             "GTEQ",
	LT:               "LT",
	LTEQ:             "LTEQ",
	ASSIGN:           "ASSIGN",
	DOT:              "DOT",
	COLON:            "COLON",
	PIPE:             "PIPE",
	COMMA:            "COMMA",
	SEMICOLON:        "SEMICOLON",
	EXPRESSION_BEGIN: "EXPRESSION_BEGIN",
	EXPRESSION_END:   "EXPRESSION_END",
	COMMENT_BEGIN:    "COMMENT_BEGIN",
	COMMENT_END:      "COMMENT_END",
	EXTERNAL_CONTENT: EXTERNAL_CONTENT,
	EOF:              EOF,
	ILLEGAL:          ILLEGAL,
}

var blockPairs = []struct{ begin, end string }{
	{"{%", "%}"},
	{"{#", "#}"},
	{"{{", "}}"},
}

func (token *Token) closesBlock() bool {
	for _, blockPair := range blockPairs {
		if blockPair.end == token.Value {
			return true
		}
	}
	return false
}

func (token *Token) String() string {
	tokenType := tokenNames[token.Type]
	return fmt.Sprintf("<%v '%s'>", tokenType, token.Value)
}
