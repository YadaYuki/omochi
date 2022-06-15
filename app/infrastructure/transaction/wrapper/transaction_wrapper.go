package wrapper

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/app/ent"
)

type EntTransactionWrapper struct {
	db *ent.Client
}

func NewEntTransactionWrapper(db *ent.Client) *EntTransactionWrapper {
	return &EntTransactionWrapper{db: db}
}

func (m *EntTransactionWrapper) WithTx(ctx context.Context, fn func(transactionClient *ent.Client) error) error {
	tx, err := m.db.Tx(ctx)
	if err != nil {
		return err
	}
	if err := fn(tx.Client()); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("rolling back transaction: %w", rollbackErr)
		}
		return err
	}
	tx.Commit()
	return nil
}
