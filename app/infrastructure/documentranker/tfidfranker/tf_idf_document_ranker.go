package tfidfranker

import (
	"context"
	"math"
	"sort"

	"github.com/YadaYuki/omochi/app/common/slices"
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
)

type TfIdfDocumentRanker struct{}

func NewTfIdfDocumentRanker() service.DocumentRanker {
	return &TfIdfDocumentRanker{}
}

func (ranker *TfIdfDocumentRanker) SortDocumentByScore(ctx context.Context, query string, docs *[]entities.Document) (*[]entities.Document, *errors.Error) {
	documentScores, _ := ranker.calculateDocumentScores(ctx, query, docs)
	contentToScoreMap := make(map[string]float64)
	for i := 0; i < len(*docs); i++ {
		contentToScoreMap[(*docs)[i].Content] = documentScores[i]
	}

	sort.Slice(*docs, func(i, j int) bool {
		// Scoreが同じだった場合は、より単語の密度が大きい、短い文章を前に.
		scoreI := contentToScoreMap[(*docs)[i].Content]
		scoreJ := contentToScoreMap[(*docs)[j].Content]
		if scoreI == scoreJ {
			return len((*docs)[i].Content) < len((*docs)[j].Content)
		}
		// Scoreが大きい方が前
		return scoreI > scoreJ
	})
	return docs, nil
}

func (ranker *TfIdfDocumentRanker) calculateDocumentScores(ctx context.Context, query string, docs *[]entities.Document) ([]float64, *errors.Error) {
	documentScores := make([]float64, len(*docs))
	idf := ranker.calculateInverseDocumentFrequency(query, docs)
	for i, doc := range *docs {
		tf := ranker.calculateTermFrequency(query, doc)
		documentScores[i] = float64(tf) * (idf + 1)
	}
	return ranker.normalize(documentScores), nil
}

func (ranker *TfIdfDocumentRanker) calculateTermFrequency(query string, doc entities.Document) int {
	termFrequency := 0
	for _, term := range doc.TokenizedContent {
		if term == query {
			termFrequency++
		}
	}
	return termFrequency
}

func (ranker *TfIdfDocumentRanker) calculateInverseDocumentFrequency(query string, docs *[]entities.Document) float64 {
	nDocs := len(*docs)
	documentFrequency := 0 // docsのうち、何個のドキュメントに、queryが含まれているか
	for _, doc := range *docs {
		if slices.Contains(doc.TokenizedContent, query) {
			documentFrequency++
		}
	}
	idf := math.Log10(float64(1+nDocs) / float64(1+documentFrequency))
	return idf
}

func (ranker *TfIdfDocumentRanker) normalize(nums []float64) []float64 {
	norm := 0.0
	for _, num := range nums {
		norm += math.Pow(num, 2)
	}
	norm = math.Pow(norm, 0.5)
	normalizeNums := make([]float64, len(nums))
	for i := 0; i < len(nums); i++ {
		normalizeNums[i] = nums[i] / norm
	}
	return normalizeNums
}
