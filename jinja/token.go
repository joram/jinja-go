package jinja

import "fmt"

type Token struct {
	Type  int
	Value string
	line  int
}

// Token definitions
const (
	ADD = iota
	ASSIGN
	COLON
	COMMA
	DIV
	DOT
	EQ
	FLOORDIV
	GT
	GTEQ
	LBRACE
	LBRACKET
	LPAREN
	LT
	LTEQ
	MOD
	MUL
	NE
	PIPE
	POW
	RBRACE
	RBRACKET
	RPAREN
	SEMICOLON
	SUB
	TILDE
	WHITESPACE
	FLOAT
	INTEGER
	NAME
	STRING
	OPERATOR
	BLOCK_BEGIN
	BLOCK_END
	VARIABLE
	RAW_BEGIN
	RAW_END
	DATA
	INITIAL
	EOF
	NEWLINE

	LINESTATEMENT_BEGIN
	LINESTATEMENT_END
	LINECOMMENT_BEGIN
	LINECOMMENT_END
	LINECOMMENT

	COMMENT_BEGIN
	COMMENT_END

	LOGIC_BEGIN
	LOGIC_END

	OUTPUT_BEGIN
	OUTPUT_END

	SPACE
	DOUBLEQUOTE
)

var operators = map[string]int{
	"==": EQ,
	"/+": ADD,
	"-":  SUB,
	"/":  DIV,
	"//": FLOORDIV,
	"*":  MUL,
	"%":  MOD,
	"**": POW,
	"~":  TILDE,
	"[":  LBRACKET,
	"]":  RBRACKET,
	"(":  LPAREN,
	")":  RPAREN,
	//"{":  LBRACE,
	//"}":  RBRACE,
	"!=": NE,
	">":  GT,
	">=": GTEQ,
	"<":  LT,
	"<=": LTEQ,
	"=":  ASSIGN,
	".":  DOT,
	":":  COLON,
	"|":  PIPE,
	",":  COMMA,
	";":  SEMICOLON,
	" ":  SPACE,
	"\"": DOUBLEQUOTE,

	"{{": OUTPUT_BEGIN,
	"}}": OUTPUT_END,
	"{%": LOGIC_BEGIN,
	"%}": LOGIC_END,
	"{#": COMMENT_BEGIN,
	"#}": COMMENT_END,
}

var operatorNames = map[int]string{
	EQ:          "EQ",
	ADD:         "ADD",
	SUB:         "SUB",
	FLOORDIV:    "FLOORDIV",
	DIV:         "DIV",
	MUL:         "MUL",
	MOD:         "MOD",
	POW:         "POW",
	TILDE:       "TILDE",
	LBRACKET:    "LBRACKET",
	RBRACKET:    "RBRACKET",
	LPAREN:      "LPAREN",
	RPAREN:      "RPAREN",
	NE:          "NE",
	GT:          "GT",
	GTEQ:        "GTEQ",
	LT:          "LT",
	LTEQ:        "LTEQ",
	ASSIGN:      "ASSIGN",
	DOT:         "DOT",
	COLON:       "COLON",
	PIPE:        "PIPE",
	COMMA:       "COMMA",
	SEMICOLON:   "SEMICOLON",
	SPACE:       "SPACE",
	DOUBLEQUOTE: "DOUBLEQUOTE",

	OUTPUT_BEGIN:  "OUTPUT_BEGIN",
	OUTPUT_END:    "OUTPUT_END",
	LOGIC_BEGIN:   "LOGIC_BEGIN",
	LOGIC_END:     "LOGIC_END",
	COMMENT_BEGIN: "COMMENT_BEGIN",
	COMMENT_END:   "COMMENT_END",
	STRING:        "STRING",
	VARIABLE:      "VARIABLE",
	EOF:           "EOF",
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
	tokenType := operatorNames[token.Type]
	return fmt.Sprintf("<%v '%s'>", tokenType, token.Value)
}
