package ja

import (
	"context"
	"fmt"
	"testing"
)

func TestTokenize(t *testing.T) {

	testCases := []struct {
		content           string
		expectedTermWords []string
	}{
		{"私は犬が好きです。", []string{"私", "犬", "好き"}},
	}
	tokenizer := NewJaKagomeTokenizer()
	for _, tc := range testCases {
		t.Run(tc.content, func(tt *testing.T) {
			terms, err := tokenizer.Tokenize(context.Background(), tc.content)
			if err != nil {
				t.Fatalf(err.Error())
			}
			fmt.Println(*terms)
			if len(*terms) != len(tc.expectedTermWords) {
				t.Fatalf("len(*terms) should be %v but got %v", len(tc.expectedTermWords), len(*terms))
			}
			for i, term := range *terms {
				if term.Word != tc.expectedTermWords[i] {
					t.Fatalf("Tokenize() should return %s, but got %s", tc.expectedTermWords[i], term.Word)
				}
			}
		})
	}
}
