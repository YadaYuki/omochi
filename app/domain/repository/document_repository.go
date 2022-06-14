package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type DocumentRepository interface {
	CreateDocument(ctx context.Context, doc *entities.Document) (*entities.DocumentDetail, *errors.Error)
}
