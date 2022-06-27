package wrapper

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
)

type EntTransactionWrapper struct {
}

func NewEntTransactionWrapper() *EntTransactionWrapper {
	return &EntTransactionWrapper{}
}

func (m *EntTransactionWrapper) WithTx(ctx context.Context, db *ent.Client, fn func(t *ent.Client) *errors.Error) *errors.Error {
	tx, err := db.Tx(ctx)
	if err != nil {
		return errors.NewError(code.Unknown, err)
	}
	if err := fn(tx.Client()); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.NewError(code.Unknown, fmt.Errorf("rolling back transaction: %w", rollbackErr))
		}
		return err
	}
	tx.Commit()
	return nil
}
