package indexer

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/app/infrastructure/transaction/wrapper"
)

type Indexer struct {
	transactionWrapper *wrapper.EntTransactionWrapper
}

func NewIndexer(wrapper *wrapper.EntTransactionWrapper) *Indexer {
	return &Indexer{transactionWrapper: wrapper}
}

func (indexer *Indexer) RegisterDocument(ctx context.Context, document *entities.Document) (*[]entities.DocumentDetail, *errors.Error) {

	indexer.transactionWrapper.WithTx(ctx, func(transactionClient *ent.Client) error { // TODO: entの知識が入るのは微妙
		termEntRepository := entdb.NewTermEntRepository(transactionClient)
		documentEntRepository := entdb.NewTermEntRepository(transactionClient)
		fmt.Println(termEntRepository, documentEntRepository)
		return nil
	})

	return nil, nil

}
