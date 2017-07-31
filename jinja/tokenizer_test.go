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
		{"{# #}", []Token{
			{COMMENT_BEGIN, "{#", 1},
			{COMMENT_END, "#}", 1},
			{Type: EOF},
		}},
		{"{{ + / - * % // == = : , ;\n. > >= < <= [] {} () ^ | != }}", []Token{
			{EXPRESSION_BEGIN, "{{", 1},
			{ADD, ADD, 1},
			{DIV, DIV, 1},
			{SUB, SUB, 1},
			{MUL, MUL, 1},
			{MOD, MOD, 1},
			{FLOORDIV, FLOORDIV, 1},
			{ASSIGN, ASSIGN, 1},
			{EQ, EQ, 1},
			{COLON, COLON, 1},
			{COMMA, COMMA, 1},
			{SEMICOLON, SEMICOLON, 1},
			{DOT, DOT, 2},
			{GT, GT, 2},
			{GTEQ, GTEQ, 2},
			{LT, LT, 2},
			{LTEQ, LTEQ, 2},
			{LBRACKET, LBRACKET, 2},
			{RBRACKET, RBRACKET, 2},
			{LBRACE, LBRACE, 2},
			{RBRACE, RBRACE, 2},
			{LPAREN, LPAREN, 2},
			{RPAREN, RPAREN, 2},
			{POW, POW, 2},
			{PIPE, PIPE, 2},
			{NE, NE, 2},
			{EXPRESSION_END, "}}", 2},
			{Type: EOF},
		}},
		{"{{ 1234 1234.5678 }}", []Token{
			{EXPRESSION_BEGIN, "{{", 1},
			{INTEGER, "1234", 1},
			{FLOAT, "1234.5678", 1},
			{EXPRESSION_END, "}}", 1},
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
