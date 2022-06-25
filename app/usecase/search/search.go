package search

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"

	"github.com/YadaYuki/omochi/app/errors"
)

type SearchUseCase interface {
	SearchDocuments(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error)
}

type searchUseCase struct {
	seacher service.Searcher
}

func NewSearchUseCase(s service.Searcher) SearchUseCase {
	return &searchUseCase{s}
}

func (s *searchUseCase) SearchDocuments(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {
	documents, err := s.seacher.Search(ctx, query)
	if err != nil {
		return nil, err
	}
	return documents, nil
}
