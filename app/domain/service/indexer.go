package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type Indexer interface {
	IndexingDocument(ctx context.Context, document *entities.DocumentCreate) *errors.Error
}
