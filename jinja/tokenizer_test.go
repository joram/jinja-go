package jinja

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	items := []struct {
		Text   string
		Tokens []Token
	}{
		{"before {{ foo }} middle {{ bar }} after", []Token{
			{EXTERNAL_CONTENT, "before", 1},
			{EXPRESSION_BEGIN, "{{", 1},
			{IDENTIFIER, "foo", 1},
			{EXPRESSION_END, "}}", 1},
			{EXTERNAL_CONTENT, "middle", 1},
			{EXPRESSION_BEGIN, "{{", 1},
			{IDENTIFIER, "bar", 1},
			{EXPRESSION_END, "}}", 1},
			{EXTERNAL_CONTENT, "after", 1},
			{Type: EOF},
		}},
	}

	for _, item := range items {
		tokenizer := NewTokenizer(item.Text)
		index := 0
		for token := range tokenizer.C {
			if index >= len(item.Tokens) {
				t.Errorf("Received an extra token: %v", token.String())
				t.FailNow()
			}
			assert.Equal(t, item.Tokens[index], token)
			index++
		}
	}
}

func BenchmarkTokenizerSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tokenizer := NewTokenizer("before {{ foo }} middle {{ bar }} after")
		for range tokenizer.C {
		}
	}
}
