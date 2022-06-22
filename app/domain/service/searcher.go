package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type Searcher interface {
	Search(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error)
}
