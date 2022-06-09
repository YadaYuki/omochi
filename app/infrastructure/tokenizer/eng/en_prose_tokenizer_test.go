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
