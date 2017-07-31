package jinja

import "fmt"

type TokenDefinition struct {
	tokenType   int
	tokenName   string
	tokenString string
	comparator  bool
	opensBlock  bool
	closesBlock bool
	operator    bool
}
type Token struct {
	Definition TokenDefinition
	Value      string
	line       int
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

	TEXT
	SPACE
	SINGLEQUOTE
	DOUBLEQUOTE
)

var tokenDefinitions = []TokenDefinition{

	// Comparators
	{EQ, "EQ", "==", true, false, false, false},
	{NE, "NE", "!=", true, false, false, false},
	{GT, "GT", ">", true, false, false, false},
	{GTEQ, "GTEQ", ">=", true, false, false, false},
	{LT, "LT", "<", true, false, false, false},
	{LTEQ, "LTEQ", "<=", true, false, false, false},

	// Operators
	{ADD, "ADD", "+", false, false, false, true},
	{SUB, "SUB", "-", false, false, false, true},
	{FLOORDIV, "FLOORDIV", "//", false, false, false, true},
	{DIV, "DIV", "/", false, false, false, true},
	{MUL, "MUL", "*", false, false, false, true},
	{MOD, "MOD", "%", false, false, false, true},
	{POW, "POW", "^", false, false, false, true},

	// blocks
	{OUTPUT_BEGIN, "OUTPUT_BEGIN", "{{", false, true, false, false},
	{OUTPUT_END, "OUTPUT_END", "}}", false, false, true, false},
	{LOGIC_BEGIN, "LOGIC_BEGIN", "%}", false, true, false, false},
	{LOGIC_END, "LOGIC_END", "{%", false, false, true, false},
	{COMMENT_BEGIN, "COMMENT_BEGIN", "{#", false, true, false, false},
	{COMMENT_END, "COMMENT_END", "#}", false, false, true, false},

	{LBRACKET, "LBRACKET", "[", false, false, false, false},
	{RBRACKET, "RBRACKET", "]", false, false, false, false},
	{LPAREN, "LPAREN", "(", false, false, false, false},
	{RPAREN, "RPAREN", ")", false, false, false, false},
	{ASSIGN, "ASSIGN", "=", false, false, false, false},
	{DOT, "DOT", ".", false, false, false, false},
	{COLON, "COLON", ":", false, false, false, false},
	{PIPE, "PIPE", "|", false, false, false, false},
	{COMMA, "COMMA", ",", false, false, false, false},
	{SEMICOLON, "SEMICOLON", ";", false, false, false, false},
	{SPACE, "SPACE", " ", false, false, false, false},
	{SINGLEQUOTE, "SINGLEQUOTE", "'", false, false, false, false},
	{DOUBLEQUOTE, "DOUBLEQUOTE", "\"", false, false, false, false},

	// shouldn't be searched for
	{STRING, "STRING", "", false, false, false, false},
	{VARIABLE, "VARIABLE", "", false, false, false, false},
	{EOF, "EOF", "", false, false, false, false},
	{TEXT, "TEXT", "", false, false, false, false},
	{TILDE, "TILDE", "", false, false, false, false},
}

func getDefinitionByType(tokenType int) TokenDefinition {
	for _, def := range tokenDefinitions {
		if def.tokenType == tokenType {
			return def
		}
	}
	panic("none found of type " + string(tokenType))
}

func getDefinitionByString(tokenString string) TokenDefinition {
	for _, def := range tokenDefinitions {
		if def.tokenString == tokenString {
			return def
		}
	}
	panic("none found")
}

func (token Token) isStatement() bool {
	if token.Definition.tokenType == STRING {
		return true
	}
	return false
}

func (token Token) String() string {
	return fmt.Sprintf("<%v '%s'>", token.Definition.tokenName, token.Value)
}
