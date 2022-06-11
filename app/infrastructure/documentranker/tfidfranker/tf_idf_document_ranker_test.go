package tfidfranker

import (
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
)

func TestCalculateTermFrequency(t *testing.T) {
	ranker := NewTfIdfDocumentRanker()
	testCases := []struct {
		doc        entities.DocumentDetail
		word       string
		expectedTf int
	}{
		{entities.DocumentDetail{TokenizedContent: []string{"sun", "is", "shining"}}, "is", 1},
	}
	for _, tc := range testCases {
		t.Run(tc.word, func(tt *testing.T) {
			tf := ranker.calculateTermFrequency(tc.word, tc.doc)
			if tc.expectedTf != tf {
				t.Fatalf("expected %v, but got %v", tc.expectedTf, tf)
			}
		})
	}
}
