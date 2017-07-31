package jinja

//func TestTokenizer(t *testing.T) {
//	testStrings := []string{
//		"before {{ foo }} middle {{ bar }} after",
//		//"before {% if 'a string with spaces in it' == foo %} middle {{ bar }} after",
//	}
//
//	for _, s := range testStrings {
//
//		fmt.Printf("tokenizing: %v\n", s)
//
//		tokenizer := NewTokenizer(s)
//		tokensChan := make(chan Token)
//		tokens := []Token{}
//		go tokenizer.GetTokens(tokensChan)
//
//		for token := range tokensChan {
//			tokens = append(tokens, token)
//		}
//
//		fmt.Printf("tokens made: %+v\n", tokens)
//	}
//	t.Errorf("")
//}
