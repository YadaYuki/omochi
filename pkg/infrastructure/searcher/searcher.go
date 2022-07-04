package searcher

import (
	"context"
	"fmt"

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

	case entities.And:
		return s.searchAnd(ctx, query)

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

	documents, documentErr := s.documentRepository.FindDocumentsByIds(ctx, &documentIds)
	if documentErr != nil {
		return nil, errors.NewError(documentErr.Code, documentErr)
	}
	return documents, nil
}

func (s *Searcher) searchAnd(ctx context.Context, query *entities.Query) ([]*entities.Document, *errors.Error) {
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

	// 辞書に登録されていない単語が含まれている場合は、その時点で空配列を返す
	for _, word := range *query.Keywords {
		if _, ok := wordToInvertIndex[word]; !ok {
			return []*entities.Document{}, nil
		}
	}

	documentIdToValidMap := map[int64]bool{}
	// keywordの一つ目のposting listから検索結果となるdocument idの候補を取得
	firstKeyword := (*query.Keywords)[0]
	for _, posting := range *(*wordToInvertIndex[firstKeyword]).PostingList {
		documentIdToValidMap[posting.DocumentRelatedId] = true
	}

	// keywordの一つ目以降のposting listから検索結果となるdocument idの候補を取得
	for _, keyword := range (*query.Keywords)[1:] {
		for id := range documentIdToValidMap {
			valid := documentIdToValidMap[id]
			if valid {
				// keywordに対応するposting list内にdocument idが存在するかを二分探索で検索
				postingList := (*wordToInvertIndex[keyword]).PostingList
				low := -1
				high := len(*postingList)
				for (high - low) > 1 {
					mid := (low + high) / 2
					if (*postingList)[mid].DocumentRelatedId < id {
						low = mid
					} else {
						high = mid
					}
				}
				if high == len(*postingList) || (*postingList)[high].DocumentRelatedId != id {
					documentIdToValidMap[id] = false
				}
			}
		}
	}

	documentIds := []int64{}
	for id := range documentIdToValidMap {
		valid := documentIdToValidMap[id]
		if valid {
			documentIds = append(documentIds, id)
		}
	}

	documents, documentErr := s.documentRepository.FindDocumentsByIds(ctx, &documentIds)
	if documentErr != nil {
		return nil, errors.NewError(documentErr.Code, documentErr)
	}
	return documents, nil
}
