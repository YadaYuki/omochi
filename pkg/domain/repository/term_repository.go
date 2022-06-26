package repository

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/google/uuid"
)

type TermRepository interface {
	FindTermCompressedById(ctx context.Context, uuid uuid.UUID) (*entities.TermCompressed, *errors.Error)
	FindTermCompressedByWord(ctx context.Context, word string) (*entities.TermCompressed, *errors.Error)
	BulkUpsertTerm(ctx context.Context, terms *[]entities.TermCompressedCreate) *errors.Error
	FindTermCompressedsByWords(ctx context.Context, words *[]string) (*[]entities.TermCompressed, *errors.Error)
}
