package searcher

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
)

type Searcher struct {
	invertIndexCached map[string]*entities.InvertIndex
}

func NewSearcher(invertIndexCached map[string]*entities.InvertIndex) service.Searcher {
	return &Searcher{invertIndexCached}
}

func (s *Searcher) Search(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {

	return nil, nil
}
