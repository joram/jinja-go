package jinja

import (
	"regexp"
	"sort"
	"strings"
)

type Tokenizer struct {
	body              string
	blockBeginIndices []int
	ch                byte

	currPosition int
	currLine     int
	inBlock      bool

	C chan Token
}

func NewTokenizer(s string) Tokenizer {
	blockBeginIndices := []int{}
	b := []byte(s)
	for _, pair := range blockPairs {
		re := regexp.MustCompile(pair.begin)
		indices := re.FindAllIndex(b, -1)
		if indices != nil {
			for _, index := range indices {
				blockBeginIndices = append(blockBeginIndices, index[0])
			}
		}
	}
	sort.Ints(blockBeginIndices)

	tokenizer := Tokenizer{
		body:              s,
		blockBeginIndices: blockBeginIndices,
		currPosition:      0,
		currLine:          1,
		inBlock:           false,
		C:                 make(chan Token),
	}
	if len(s) > 0 {
		if s[0] == '{' {
			tokenizer.inBlock = true
		}
	}
	go tokenizer.run()
	return tokenizer
}

func (t *Tokenizer) run() {
	token := Token{}
	for token.Type != EOF {
		token = t.next()
		t.C <- token
	}
	close(t.C)
}

func (t *Tokenizer) readExternalContent() Token {
	for _, beginIndex := range t.blockBeginIndices {
		if beginIndex >= t.currPosition {
			return Token{
				EXTERNAL_CONTENT,
				t.body[t.currPosition : beginIndex-1],
				t.currLine,
			}
		}
	}

	// string until EOF
	return Token{
		EXTERNAL_CONTENT,
		t.body[t.currPosition:len(t.body)],
		t.currLine,
	}
}

func (t *Tokenizer) readIdentifier() Token {
	position := t.currPosition - 1
	for isLetter(t.ch) {
		t.readChar()
	}
	end := t.currPosition - 1
	tokenType := TokenType(IDENTIFIER)
	val := t.body[position:end]
	if _, ok := keywords[val]; ok {
		tokenType = TokenType(val)
	}
	return Token{tokenType, val, t.currLine}
}

func (t *Tokenizer) readNumber() Token {
	position := t.currPosition - 1
	for isNumber(t.ch) || t.ch == '.' {
		t.readChar()
	}
	end := t.currPosition - 1
	val := t.body[position:end]
	tokenType := TokenType(INTEGER)
	if strings.Contains(val, ".") {
		tokenType = FLOAT
	}
	return Token{tokenType, val, t.currLine}
}

func (t *Tokenizer) next() Token {
	if t.currPosition >= len(t.body) {
		return Token{Type: EOF}
	}

	if !t.inBlock {
		token := t.readExternalContent()
		t.inBlock = true
		t.currPosition += len(token.Value)
		return token
	}

	t.skipWhitespace()

	token := Token{ILLEGAL, string(t.ch), t.currLine}
	switch t.ch {
	case '+':
		token = Token{ADD, ADD, t.currLine}
	case '%':
		token = Token{MOD, MOD, t.currLine}
	case '/':
		if t.peek() == '/' {
			t.readChar()
			token = Token{FLOORDIV, FLOORDIV, t.currLine}
		} else {
			token = Token{DIV, DIV, t.currLine}
		}
	case '-':
		token = Token{SUB, SUB, t.currLine}
	case '*':
		token = Token{MUL, MUL, t.currLine}
	case '^':
		token = Token{POW, POW, t.currLine}
	case '|':
		token = Token{PIPE, PIPE, t.currLine}
	case '=':
		if t.peek() == '=' {
			t.readChar()
			token = Token{ASSIGN, ASSIGN, t.currLine}
		} else {
			token = Token{EQ, EQ, t.currLine}
		}
	case '>':
		if t.peek() == '=' {
			t.readChar()
			token = Token{GTEQ, GTEQ, t.currLine}
		} else {
			token = Token{GT, GT, t.currLine}
		}
	case '<':
		if t.peek() == '=' {
			t.readChar()
			token = Token{LTEQ, LTEQ, t.currLine}
		} else {
			token = Token{LT, LT, t.currLine}
		}
	case ':':
		token = Token{COLON, COLON, t.currLine}
	case ';':
		token = Token{SEMICOLON, SEMICOLON, t.currLine}
	case '.':
		token = Token{DOT, DOT, t.currLine}
	case ',':
		token = Token{COMMA, COMMA, t.currLine}
	case '[':
		token = Token{LBRACKET, LBRACKET, t.currLine}
	case ']':
		token = Token{RBRACKET, RBRACKET, t.currLine}
	case '(':
		token = Token{LPAREN, LPAREN, t.currLine}
	case ')':
		token = Token{RPAREN, RPAREN, t.currLine}
	case '{':
		if t.peek() == '{' {
			t.readChar()
			token = Token{EXPRESSION_BEGIN, EXPRESSION_BEGIN, t.currLine}
		} else if t.peek() == '#' {
			t.readChar()
			token = Token{COMMENT_BEGIN, COMMENT_BEGIN, t.currLine}
		} else {
			token = Token{LBRACE, LBRACE, t.currLine}
		}
	case '}':
		if t.peek() == '}' {
			t.readChar()
			token = Token{EXPRESSION_END, EXPRESSION_END, t.currLine}
			t.inBlock = false
		} else {
			token = Token{RBRACE, RBRACE, t.currLine}
		}
	case '#':
		if t.peek() == '}' {
			t.readChar()
			token = Token{COMMENT_END, COMMENT_END, t.currLine}
			t.inBlock = false
		}
	case '!':
		if t.peek() == '=' {
			t.readChar()
			token = Token{NE, NE, t.currLine}
		} else {
			token = t.readIdentifier()
		}
	default:
		if isNumber(t.peek()) {
			token = t.readNumber()
		} else {
			token = t.readIdentifier()
		}
	}

	t.readChar()
	return token
}

func (t *Tokenizer) peek() byte {
	return t.body[t.currPosition]
}

func (t *Tokenizer) skipWhitespace() {
	for t.ch == ' ' || t.ch == '\t' || t.ch == '\n' || t.ch == '\r' || t.ch == 0 {
		if t.ch == '\n' {
			t.currLine += 1
		}
		t.readChar()
	}
}

func (t *Tokenizer) readChar() {
	if t.currPosition >= len(t.body) {
		t.ch = 0
	} else {
		t.ch = t.body[t.currPosition]
	}
	t.currPosition += 1
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
