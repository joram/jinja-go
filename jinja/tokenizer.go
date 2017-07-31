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
	for _, def := range tokenDefinitions {
		if def.opensBlock {
			re := regexp.MustCompile(def.tokenString)
			indices := re.FindAllIndex(b, -1)
			if indices != nil {
				for _, index := range indices {
					blockBeginIndices = append(blockBeginIndices, index[0])
				}
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

func (t *Tokenizer) nextTextToken() Token {
	for _, beginIndex := range t.blockBeginIndices {
		if beginIndex >= t.currPosition {
			return Token{
				getDefinitionByType(TEXT),
				t.body[t.currPosition:beginIndex],
				t.currLine,
			}
		}
	}

	// string until EOF
	return Token{
		getDefinitionByType(TEXT),
		t.body[t.currPosition:t.length],
		t.currLine,
	}
}

func (t *Tokenizer) next() Token {
	if t.currPosition == t.length {
		return Token{getDefinitionByType(EOF), "", 0}
	}

	if !t.inBlock {
		token := t.nextTextToken()
		t.inBlock = true
		t.currPosition += len(token.Value)
		return token
	}

	for _, def := range tokenDefinitions {
		//fmt.Printf("does '%s' start with '%s' ?\n", t.body[t.currPosition:t.currPosition+5], def.tokenString)
		if len(def.tokenString) > 0 && strings.HasPrefix(t.body[t.currPosition:], def.tokenString) {
			token := Token{
				def,
				t.body[t.currPosition : t.currPosition+len(def.tokenString)],
				t.currLine,
			}
			if token.Definition.closesBlock {
				t.inBlock = false
			}
			t.currPosition += len(def.tokenString)
			return token
		}
	}

	nextWord := t.nextWord()
	t.currPosition += len(nextWord)
	return Token{
		getDefinitionByType(VARIABLE),
		nextWord,
		t.currLine,
	}
}

func (t *Tokenizer) nextWord() string {
	re := regexp.MustCompile("[^ ]+")
	indices := re.FindAllIndex([]byte(t.body[t.currPosition:]), -1)

	firstVariable := indices[0]
	if firstVariable[0] != 0 {
		panic("a token should start now")
	}

	variableLength := firstVariable[1]
	return t.body[t.currPosition : t.currPosition+variableLength]
}

func (tokenizer *Tokenizer) GetTokens(tokens chan Token) {
	foundEOF := false
	i := 0
	for !foundEOF {
		t := tokenizer.next()
		foundEOF = t.Definition.tokenType == EOF
		tokens <- t
		//fmt.Printf("%d: %v\n", i, t)
		i += 1
	}
	close(tokens)
}
