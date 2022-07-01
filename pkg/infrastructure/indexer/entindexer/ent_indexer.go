package entindexer

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/infrastructure/indexer"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/transaction/wrapper"
)

type EntIndexer struct {
	db                    *ent.Client
	t                     *wrapper.EntTransactionWrapper
	tokenizer             service.Tokenizer
	invertIndexCompresser service.InvertIndexCompresser
}

func NewEntIndexer(db *ent.Client, t *wrapper.EntTransactionWrapper, tokenizer service.Tokenizer, invertIndexCompresser service.InvertIndexCompresser) *EntIndexer {
	return &EntIndexer{db: db, t: t, tokenizer: tokenizer, invertIndexCompresser: invertIndexCompresser}
}

// IndexingDocumentWithTx is a function for indexing a document with RDB Transaction.
func (entIndexer *EntIndexer) IndexingDocumentWithTx(ctx context.Context, document *entities.DocumentCreate) *errors.Error {

	indexingDocumentFunc := func(transactionClient *ent.Client) *errors.Error {
		documentRepository := entdb.NewDocumentEntRepository(transactionClient)
		termRepository := entdb.NewTermEntRepository(transactionClient)
		indexer := indexer.NewIndexer(documentRepository, termRepository, entIndexer.tokenizer, entIndexer.invertIndexCompresser)
		return indexer.IndexingDocument(ctx, document)
	}

	err := entIndexer.t.WithTx(ctx, entIndexer.db, indexingDocumentFunc)
	if err != nil {
		return err
	}
	return nil
}
