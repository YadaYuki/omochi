package tfidfranker

import (
	"context"
	"fmt"
	"math"

	"github.com/YadaYuki/omochi/app/common/slices"
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type TfIdfDocumentRanker struct{}

func NewTfIdfDocumentRanker() *TfIdfDocumentRanker {
	return &TfIdfDocumentRanker{}
}

func (ranker *TfIdfDocumentRanker) SortDocumentByScore(ctx context.Context, query string, docs *[]entities.DocumentDetail) (*[]entities.DocumentDetail, *errors.Error) {
	fmt.Println(ranker.calculateDocumentScores(ctx, query, docs))
	return nil, nil
}

func (ranker *TfIdfDocumentRanker) calculateDocumentScores(ctx context.Context, query string, docs *[]entities.DocumentDetail) ([]float64, *errors.Error) {
	return []float64{}, nil
}

func (ranker *TfIdfDocumentRanker) calculateTermFrequency(query string, doc entities.DocumentDetail) int {
	termFrequency := 0
	for _, term := range doc.TokenizedContent {
		if term == query {
			termFrequency++
		}
	}
	return termFrequency
}
func (ranker *TfIdfDocumentRanker) calculateInverseDocumentFrequency(query string, docs *[]entities.DocumentDetail) float64 {
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
