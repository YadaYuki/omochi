package service

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type DocumentRanker interface {
	SortDocumentByScore(ctx context.Context, query string, docs []*entities.Document) ([]*entities.Document, *errors.Error)
}
