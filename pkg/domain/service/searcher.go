package service

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type Searcher interface {
	Search(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error)
}
