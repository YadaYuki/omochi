package indexer

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/infrastructure/transaction/wrapper"
)

type Indexer struct {
	documentRepository repository.DocumentRepository
	tokenizer          service.Tokenizer
	// invertIndexCompresser service.InvertIndexCompresser
}

func NewIndexer(wrapper *wrapper.EntTransactionWrapper, documentRepository repository.DocumentRepository, tokenizer service.Tokenizer) service.Indexer {
	return &Indexer{}
}

func (i *Indexer) IndexingDocument(ctx context.Context, document *entities.DocumentCreate) (*[]entities.Document, *errors.Error) {

	// Create Invert Index from document
	tokenizedContent, tokenizeErr := i.tokenizer.Tokenize(ctx, document.Content)
	if tokenizeErr != nil {
		return nil, errors.NewError(tokenizeErr.Code, tokenizeErr.Error())
	}
	document.TokenizedContent = make([]string, len(*tokenizedContent))
	for i, term := range *tokenizedContent {
		document.TokenizedContent[i] = term.Word
	}
	documentDetail, documentCreateErr := i.documentRepository.CreateDocument(ctx, document)
	if documentCreateErr != nil {
		return nil, errors.NewError(documentCreateErr.Code, documentCreateErr.Error())
	}
	// create invert indexes
	documentId := documentDetail.Id
	wordToPostingMap := make(map[string]*entities.Posting)
	for position, word := range document.TokenizedContent {
		if _, ok := wordToPostingMap[word]; ok {
			wordToPostingMap[word].PositionsInDocument = append(wordToPostingMap[word].PositionsInDocument, position)
		} else {
			positionsInDocument := []int{position}
			wordToPostingMap[word] = entities.NewPosting(documentId, positionsInDocument)
		}
	}

	return nil, nil
}
