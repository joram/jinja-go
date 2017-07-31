package jinja

import (
	"regexp"
	"sort"
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
	return Token{IDENTIFIER, t.body[position:end], t.currLine}
}

func (t *Tokenizer) next() Token {
	if t.currPosition == len(t.body) {
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
	case '{':
		if t.peek() == '{' {
			t.readChar()
			token = Token{EXPRESSION_BEGIN, EXPRESSION_BEGIN, t.currLine}
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
	default:
		token = t.readIdentifier()
	}

	t.readChar()
	return token
	//for operator, tokenType := range operators {
	//	if strings.HasPrefix(t.body[t.currPosition:], operator) {
	//		token := Token{
	//			tokenType,
	//			t.body[t.currPosition : t.currPosition+len(operator)],
	//			t.currLine,
	//		}
	//		if token.closesBlock() {
	//			t.inBlock = false
	//		}
	//		t.currPosition += len(operator)
	//		return token
	//	}
	//}
	//
	//token := t.readIdentifier()
	//t.currPosition += len(token.Value)
	//return token
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
