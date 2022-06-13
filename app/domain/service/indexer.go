package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type Indexer interface {
	RegisterDocument(ctx context.Context, document *entities.Document) (*[]entities.DocumentDetail, *errors.Error)
}
