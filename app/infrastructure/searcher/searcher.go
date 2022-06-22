package searcher

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
)

type Searcher struct {
	invertIndexCached  map[string]*entities.InvertIndex
	termRepository     repository.TermRepository
	documentRepository repository.DocumentRepository
	compresser         service.InvertIndexCompresser
	documentRanker     service.DocumentRanker
}

func NewSearcher(invertIndexCached map[string]*entities.InvertIndex, termRepository repository.TermRepository, documentRepository repository.DocumentRepository, compresser service.InvertIndexCompresser, documentRanker service.DocumentRanker) service.Searcher {
	return &Searcher{invertIndexCached, termRepository, documentRepository, compresser, documentRanker}
}

func (s *Searcher) Search(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {

	invertIndex, ok := s.invertIndexCached[query.Keyword]
	if !ok {
		termCompressed, err := s.termRepository.FindTermCompressedByWord(ctx, query.Keyword)
		if err != nil {
			return nil, errors.NewError(err.Code, err)
		}
		invertIndexCompressed := termCompressed.InvertIndexCompressed
		invertIndex, err = s.compresser.Decompress(ctx, invertIndexCompressed)
		if err != nil {
			return nil, errors.NewError(err.Code, err)
		}
	}

	documentIds := []int64{}
	for _, postingList := range *invertIndex.PostingList {
		documentIds = append(documentIds, postingList.DocumentRelatedId)
	}

	documents, documentErr := s.documentRepository.FindDocumentsByIds(ctx, &documentIds)
	if documentErr != nil {
		return nil, errors.NewError(documentErr.Code, documentErr)
	}
	sortedDocument, sortErr := s.documentRanker.SortDocumentByScore(ctx, query.Keyword, documents)
	if sortErr != nil {
		return nil, errors.NewError(sortErr.Code, sortErr)
	}

	return sortedDocument, nil
}
