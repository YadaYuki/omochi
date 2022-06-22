package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type Searcher interface {
	Search(ctx context.Context, query entities.SearchModeType) ([]*entities.Document, *errors.Error)
}
