package searcher

import (
	"context"
	"fmt"
	"sort"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/repository"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
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
	if len(*query.Keywords) == 1 {
		return s.searchBySingleKeyword(ctx, query)
	}
	switch query.SearchMode {
	case entities.Or:
		return s.searchOr(ctx, query)
	default:
		return nil, errors.NewError(code.Unknown, fmt.Sprintf("unsupported search mode: %s", query.SearchMode))
	}
}

func (s *Searcher) searchBySingleKeyword(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {
	invertIndex, ok := s.invertIndexCached[(*query.Keywords)[0]]
	if !ok {
		termCompressed, err := s.termRepository.FindTermCompressedByWord(ctx, (*query.Keywords)[0])
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
	sortedDocument, sortErr := s.documentRanker.SortDocumentByScore(ctx, (*query.Keywords)[0], documents)
	if sortErr != nil {
		return nil, errors.NewError(sortErr.Code, sortErr)
	}
	return sortedDocument, nil
}

func (s *Searcher) searchOr(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {
	wordToInvertIndex := map[string]*entities.InvertIndex{}
	wordsNotInCache := []string{}
	for _, word := range *query.Keywords {
		invertIndex, ok := s.invertIndexCached[word]
		if !ok {
			wordsNotInCache = append(wordsNotInCache, word)
		} else {
			wordToInvertIndex[word] = invertIndex
		}
	}

	if len(wordsNotInCache) > 0 {
		termCompresseds, err := s.termRepository.FindTermCompressedsByWords(ctx, &wordsNotInCache)
		if err != nil {
			return nil, errors.NewError(err.Code, err)
		}
		for _, termCompressed := range *termCompresseds {
			invertIndexCompressed := termCompressed.InvertIndexCompressed
			invertIndex, decompressErr := s.compresser.Decompress(ctx, invertIndexCompressed)
			if decompressErr != nil {
				return nil, errors.NewError(err.Code, decompressErr)
			}
			wordToInvertIndex[termCompressed.Word] = invertIndex
		}
	}

	documentIdsMap := map[int64]bool{}

	for _, keyword := range *query.Keywords {
		for _, posting := range *(*wordToInvertIndex[keyword]).PostingList {
			documentIdsMap[posting.DocumentRelatedId] = true
		}
	}

	documentIds := []int64{}
	for id := range documentIdsMap {
		documentIds = append(documentIds, id)
	}

	sort.Slice(documentIds, func(i, j int) bool {
		return documentIds[i] < documentIds[j]
	})

	documents, documentErr := s.documentRepository.FindDocumentsByIds(ctx, &documentIds)
	if documentErr != nil {
		return nil, errors.NewError(documentErr.Code, documentErr)
	}
	return documents, nil
}
