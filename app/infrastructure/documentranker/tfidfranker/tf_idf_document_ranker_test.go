package tfidfranker

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
)

func TestCalculateTermFrequency(t *testing.T) {
	ranker := &TfIdfDocumentRanker{}
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
	ranker := &TfIdfDocumentRanker{}
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

func TestNormalize(t *testing.T) {
	ranker := &TfIdfDocumentRanker{}

	testCases := []struct {
		nums               []float64
		expectedNormalized []float64
	}{
		{[]float64{1.0, 1.0, 1.0}, []float64{0.577, 0.577, 0.577}},
		{[]float64{1.0, 2.0, 3.0}, []float64{0.267, 0.535, 0.802}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.nums), func(tt *testing.T) {
			normalized := ranker.normalize(tc.nums)
			for i, item := range normalized {
				if math.Abs(item-tc.expectedNormalized[i]) > 1e-3 {
					t.Fatalf("expected %v, but got %v", tc.expectedNormalized[i], item)
				}
			}
		})
	}
}

func TestCalculateDocumentScores(t *testing.T) {
	ranker := &TfIdfDocumentRanker{}
	documents := &[]entities.DocumentDetail{
		{TokenizedContent: []string{"sun", "is", "shining"}},
		{TokenizedContent: []string{"weather", "is", "sweet"}},
		{TokenizedContent: []string{"sun", "is", "shining", "weather", "is", "sweet"}},
	}
	testCases := []struct {
		word           string
		expectedScores []float64
	}{
		{"sun", []float64{0.707, 0.0, 0.707}},
		{"is", []float64{0.408, 0.408, 0.816}},
		{"shining", []float64{0.707, 0.0, 0.707}},
	}
	for _, tc := range testCases {
		t.Run(tc.word, func(tt *testing.T) {
			documentScores, _ := ranker.calculateDocumentScores(context.Background(), tc.word, documents)
			for i, item := range documentScores {
				if math.Abs(item-tc.expectedScores[i]) > 1e-3 {
					t.Fatalf("expected %v, but got %v", tc.expectedScores[i], item)
				}
			}
		})
	}
}

func TestSortDocumentByScore(t *testing.T) {
	ranker := &TfIdfDocumentRanker{}
	documents := []entities.DocumentDetail{
		{Content: "sun is shining", TokenizedContent: []string{"sun", "is", "shining"}},
		{Content: "weather is sweet", TokenizedContent: []string{"weather", "is", "sweet"}},
		{Content: "sun is shining weather is sweet", TokenizedContent: []string{"sun", "is", "shining", "weather", "is", "sweet"}},
	}
	testCases := []struct {
		word                   string
		expectedSortedContents []string
	}{
		{"sun", []string{"sun is shining", "sun is shining weather is sweet", "weather is sweet"}},
		{"is", []string{"sun is shining weather is sweet", "sun is shining", "weather is sweet"}},
		{"weather", []string{"weather is sweet", "sun is shining weather is sweet", "sun is shining"}},
	}
	for _, tc := range testCases {
		t.Run(tc.word, func(tt *testing.T) {
			ranker.SortDocumentByScore(context.Background(), tc.word, &documents)
			for i, doc := range documents {
				if doc.Content != tc.expectedSortedContents[i] {
					t.Fatalf("expected %v, but got %v", tc.expectedSortedContents[i], doc.Content)
				}
			}
		})
	}
}
