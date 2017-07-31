package jinja

import "fmt"

type Token struct {
	Type  TokenType
	Value string
	line  int
}

type TokenType string

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
	EOF              = "EOF"
	COMMENT_BEGIN    = "{#"
	COMMENT_END      = "#}"
	EXPRESSION_BEGIN = "{{"
	EXPRESSION_END   = "}}"

	// Keywords
	NOT      = "not"
	IF       = "if"
	AND      = "and"
	OR       = "or"
	ELIF     = "elif"
	ELSE     = "else"
	IN       = "in"
	DEL      = "del"
	TRUE     = "True"
	FALSE    = "False"
	YIELD    = "yield"
	IS       = "is"
	LAMBDA   = "lambda"
	FOR      = "for"
	CONTINUE = "continue"
	NONE     = "none"
)

var keywords = map[string]TokenType{
	"not":      NOT,
	"if":       IF,
	"and":      AND,
	"or":       OR,
	"elif":     ELIF,
	"in":       IN,
	"del":      DEL,
	"True":     TRUE,
	"False":    FALSE,
	"yield":    YIELD,
	"is":       IS,
	"lambda":   LAMBDA,
	"for":      FOR,
	"none":     NONE,
	"else":     ELSE,
	"continue": CONTINUE,
}

var tokenNames = map[TokenType]string{
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
	NOT:              "not",
	IF:               "if",
	AND:              "and",
	OR:               "or",
	ELIF:             "elif",
	ELSE:             "else",
	IN:               "in",
	DEL:              "del",
	TRUE:             "True",
	FALSE:            "False",
	YIELD:            "yield",
	IS:               "is",
	LAMBDA:           "lambda",
	FOR:              "for",
	CONTINUE:         "continue",
	NONE:             "none",
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
