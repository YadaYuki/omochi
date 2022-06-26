package service

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type Indexer interface {
	IndexingDocument(ctx context.Context, document *entities.DocumentCreate) *errors.Error
}
