package repository

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type DocumentRepository interface {
	CreateDocument(ctx context.Context, doc *entities.DocumentCreate) (*entities.Document, *errors.Error)
	FindDocumentsByIds(ctx context.Context, ids *[]int64) ([]*entities.Document, *errors.Error)
}
