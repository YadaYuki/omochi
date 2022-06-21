package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/google/uuid"
)

type TermRepository interface {
	FindTermById(ctx context.Context, uuid uuid.UUID) (*entities.Term, *errors.Error)
	BulkUpsertTerm(ctx context.Context, terms *[]entities.TermCreate) (*[]entities.Term, *errors.Error)
	FindTermsByWords(ctx context.Context, words *[]string) (*[]entities.Term, *errors.Error)
}
