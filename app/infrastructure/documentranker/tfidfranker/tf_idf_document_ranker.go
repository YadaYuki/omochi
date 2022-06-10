package tfidfranker

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type TfIdfDocumentRanker struct{}

func NewTfIdfDocumentRanker() *TfIdfDocumentRanker {
	return &TfIdfDocumentRanker{}
}

func (ranker *TfIdfDocumentRanker) SortDocumentByScore(ctx context.Context, query string, docs *[]entities.DocumentDetail) (*[]entities.DocumentDetail, *errors.Error) {
	fmt.Println(ranker.calculateDocumentScore(ctx, query, docs))
	return nil, nil
}

func (ranker *TfIdfDocumentRanker) calculateDocumentScore(ctx context.Context, query string, docs *[]entities.DocumentDetail) (float64, *errors.Error) {
	return 0.0, nil
}
