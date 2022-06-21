package indexer

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type Indexer struct {
	documentRepository    repository.DocumentRepository
	termRepository        repository.TermRepository
	tokenizer             service.Tokenizer
	invertIndexCompresser service.InvertIndexCompresser
}

func NewIndexer(documentRepository repository.DocumentRepository, termRepository repository.TermRepository, tokenizer service.Tokenizer, invertIndexCompresser service.InvertIndexCompresser) service.Indexer {
	return &Indexer{documentRepository, termRepository, tokenizer, invertIndexCompresser}
}

func (indexer *Indexer) IndexingDocument(ctx context.Context, document *entities.DocumentCreate) *errors.Error {

	// ドキュメント(文書)の新規作成
	tokenizedContent, tokenizeErr := indexer.tokenizer.Tokenize(ctx, document.Content)
	if tokenizeErr != nil {
		return errors.NewError(tokenizeErr.Code, tokenizeErr.Error())
	}
	document.TokenizedContent = make([]string, len(*tokenizedContent))
	for i, term := range *tokenizedContent {
		document.TokenizedContent[i] = term.Word
	}

	documentDetail, documentCreateErr := indexer.documentRepository.CreateDocument(ctx, document)
	if documentCreateErr != nil {
		return errors.NewError(documentCreateErr.Code, documentCreateErr.Error())
	}

	// ポスティングの作成
	wordToPostingMap := make(map[string]*entities.Posting)
	for position, word := range document.TokenizedContent {
		if _, ok := wordToPostingMap[word]; ok {
			wordToPostingMap[word].PositionsInDocument = append(wordToPostingMap[word].PositionsInDocument, position)
		} else {
			positionsInDocument := []int{position}
			wordToPostingMap[word] = entities.NewPosting(documentDetail.Id, positionsInDocument)
		}
	}

	// 文書内に登場する単語の中で、既にストレージに登録済みのものに関しては、転置インデックスを取得する
	termCompresseds, termErr := indexer.termRepository.FindTermCompressedsByWords(ctx, &document.TokenizedContent)
	if termErr != nil {
		return errors.NewError(documentCreateErr.Code, termErr.Error())
	}

	// 取得した圧縮済み転置インデックスの解凍 & wordToTermsMapの作成
	terms := make([]entities.TermCreate, len(*termCompresseds))
	wordToTermsMap := make(map[string]*entities.TermCreate)
	for i, termCompressed := range *termCompresseds {
		invertIndex, decompressErr := indexer.invertIndexCompresser.Decompress(ctx, termCompressed.InvertIndexCompressed)
		if decompressErr != nil {
			panic(decompressErr)
		}
		terms[i].Word = termCompressed.Word
		terms[i].InvertIndex = invertIndex
		wordToTermsMap[termCompressed.Word] = &terms[i]
	}
	// PostingをAppendする
	for wordInDocument, posting := range wordToPostingMap {
		if _, ok := wordToTermsMap[wordInDocument]; ok {
			*wordToTermsMap[wordInDocument].InvertIndex.PostingList = append(*wordToTermsMap[wordInDocument].InvertIndex.PostingList, *posting)
		} else {
			invertIndex := entities.NewInvertIndex(&[]entities.Posting{*posting})
			wordToTermsMap[wordInDocument] = entities.NewTermCreate(wordInDocument, invertIndex)
		}
	}

	upsertTermCompresseds := &[]entities.TermCompressedCreate{}
	for wordInDocument := range wordToTermsMap {
		invertIndexCompressed, compressErr := indexer.invertIndexCompresser.Compress(ctx, wordToTermsMap[wordInDocument].InvertIndex)
		if compressErr != nil {
			panic(compressErr)
		}
		termCompressed := entities.NewTermCompressedCreate(wordInDocument, invertIndexCompressed)
		*upsertTermCompresseds = append(*upsertTermCompresseds, *termCompressed)
	}

	// 転置インデックスの永続化
	upsertErr := indexer.termRepository.BulkUpsertTerm(ctx, upsertTermCompresseds)
	if upsertErr != nil {
		return errors.NewError(code.Unknown, upsertErr)
	}

	return nil
}
