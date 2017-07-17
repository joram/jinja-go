package jinja

import (
	"regexp"
	"sort"
	"strings"
)

type Tokenizer struct {
	body              string
	length            int
	blockBeginIndices []int

	currPosition int
	currLine     int
	inBlock      bool
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

	return Tokenizer{
		s,
		len(s),
		blockBeginIndices,
		0,
		1,
		false,
	}
}

func (t *Tokenizer) nextStringToken() Token {
	for _, beginIndex := range t.blockBeginIndices {
		if beginIndex >= t.currPosition {
			return Token{
				STRING,
				t.body[t.currPosition:beginIndex],
				t.currLine,
			}
		}
	}

	// string until EOF
	return Token{
		STRING,
		t.body[t.currPosition:t.length],
		t.currLine,
	}
}

func (t *Tokenizer) nextVariableToken() Token {
	re := regexp.MustCompile("[a-zA-Z]+")
	indices := re.FindAllIndex([]byte(t.body[t.currPosition:]), -1)

	if len(indices) == 0 {
		panic("found no vars!")
	}

	firstVariable := indices[0]
	if firstVariable[0] != 0 {
		panic("a token should start now")
	}

	variableLength := firstVariable[1]
	return Token{
		VARIABLE,
		t.body[t.currPosition : t.currPosition+variableLength],
		t.currLine,
	}
}

func (t *Tokenizer) next() Token {
	if t.currPosition == t.length {
		return Token{Type: EOF}
	}

	if !t.inBlock {
		token := t.nextStringToken()
		t.inBlock = true
		t.currPosition += len(token.Value)
		return token
	}

	for operator, tokenType := range operators {
		if strings.HasPrefix(t.body[t.currPosition:], operator) {
			token := Token{
				tokenType,
				t.body[t.currPosition : t.currPosition+len(operator)],
				t.currLine,
			}
			if token.closesBlock() {
				t.inBlock = false
			}
			t.currPosition += len(operator)
			return token
		}
	}

	token := t.nextVariableToken()
	t.currPosition += len(token.Value)
	return token
}

func (t *Tokenizer) GetTokens(tokens chan Token) {
	token := Token{}
	for token.Type != EOF {
		token = t.next()
		tokens <- token
	}
	tokens <- token
}
