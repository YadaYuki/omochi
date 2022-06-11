package tfidfranker

import (
	"math"
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
		{entities.DocumentDetail{TokenizedContent: []string{"sun", "is", "shining"}}, "hoge", 0},
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

func TestCalculateInverseDocumentFrequency(t *testing.T) {
	ranker := NewTfIdfDocumentRanker()
	documents := &[]entities.DocumentDetail{
		{TokenizedContent: []string{"sun", "is", "shining"}},
		{TokenizedContent: []string{"weather", "is", "sweet"}},
		{TokenizedContent: []string{"sun", "is", "shining", "weather", "is", "sweet"}},
	}
	testCases := []struct {
		docs        *[]entities.DocumentDetail
		word        string
		expectedIdf float32
	}{
		{documents, "is", 0.0},
		{documents, "sun", 0.125},
		{documents, "weather", 0.125},
	}
	for _, tc := range testCases {
		t.Run(tc.word, func(tt *testing.T) {
			idf := ranker.calculateInverseDocumentFrequency(tc.word, tc.docs)
			// 小数点第3位までが一致しているかどうかで比較.
			if math.Abs(float64(tc.expectedIdf)-float64(idf)) > 1e-3 {
				t.Fatalf("expected %v, but got %v", tc.expectedIdf, idf)
			}
		})
	}
}
