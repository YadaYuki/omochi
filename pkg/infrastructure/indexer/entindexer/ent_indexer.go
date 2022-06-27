package entindexer

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/indexer"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/eng"
	"github.com/YadaYuki/omochi/pkg/infrastructure/transaction/wrapper"
)

type EntIndexer struct {
	db *ent.Client
	t  *wrapper.EntTransactionWrapper
}

func NewEntIndexer(db *ent.Client, t *wrapper.EntTransactionWrapper) *EntIndexer {
	return &EntIndexer{db: db, t: t}
}

// IndexingDocumentWithTx is a function for indexing a document with RDB Transaction.
func (entIndexer *EntIndexer) IndexingDocumentWithTx(ctx context.Context, document *entities.DocumentCreate) *errors.Error {

	if err := entIndexer.t.WithTx(ctx, entIndexer.db, func(transactionClient *ent.Client) *errors.Error {
		documentRepository := entdb.NewDocumentEntRepository(transactionClient)
		termRepository := entdb.NewTermEntRepository(transactionClient)
		tokenizer := eng.NewEnProseTokenizer()
		invertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
		indexer := indexer.NewIndexer(documentRepository, termRepository, tokenizer, invertIndexCompresser)
		return indexer.IndexingDocument(ctx, document)

	}); err != nil {
		return err
	}
	return nil
}
