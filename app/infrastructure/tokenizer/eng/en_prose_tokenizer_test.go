package eng

import (
	"context"
	"testing"
)

func TestTokenize(t *testing.T) {

	testCases := []struct {
		content           string
		expectedTermWords []string
	}{
		{"hoge fuga piyo", []string{"hoge", "fuga", "piyo"}},
		{"I have a pen", []string{"i", "have", "pen"}},                                                        // a,theなどの冠詞は除去 / 単語は小文字に統一.
		{"I have a pen , you don't have pens.", []string{"i", "have", "pen", "you", "don't", "have", "pens"}}, // .も除去
	}
	for _, tc := range testCases {
		tokenizer := NewEnProseTokenizer()
		t.Run(tc.content, func(tt *testing.T) {
			terms, err := tokenizer.Tokenize(context.Background(), tc.content)
			if err != nil {
				t.Fatalf(err.Error())
			}
			for i, term := range *terms {
				if term.Word != tc.expectedTermWords[i] {
					t.Fatalf("Tokenize() should return %s, but got %s", tc.expectedTermWords[i], term.Word)
				}
			}
		})
	}
}
