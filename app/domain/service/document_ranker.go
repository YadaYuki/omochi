package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type DocumentRanker interface {
	SortDocumentByScore(ctx context.Context, query string, docs *[]entities.DocumentDetail) (*[]entities.DocumentDetail, *errors.Error)
	calculateDocumentScores(ctx context.Context, query string, docs *[]entities.DocumentDetail) ([]float64, *errors.Error)
}

// type DocumentRanker[T entities.Document | entities.DocumentDetail] interface {
// 	SortDocumentByScore(ctx context.Context, query T, docs *[]T) (*[]T, *errors.Error)
// 	calculateDocumentScore(ctx context.Context, query string, docs *[]T) (float64, *errors.Error)
// }
