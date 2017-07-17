package jinja

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	s := "before {{ foo }} middle {{ bar }} after"
	fmt.Printf("tokenizing: %v\n", s)

	tokenizer := NewTokenizer(s)
	tokensChan := make(chan Token)
	go tokenizer.GetTokens(tokensChan)
	for token := range tokensChan {
		fmt.Printf("%v\n", token.String())
		if token.Type == EOF {
			close(tokensChan)
		}
	}

	t.Errorf("")
}
