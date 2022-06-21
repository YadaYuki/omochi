package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/google/uuid"
)

type TermRepository interface {
	FindTermById(ctx context.Context, uuid uuid.UUID) (*entities.TermDetail, *errors.Error)
	BulkCreateTerm(ctx context.Context, terms *[]entities.Term) (*[]entities.TermDetail, *errors.Error)
	FindTermsByWords(ctx context.Context, words *[]string, compresser service.InvertIndexCompresser) (*[]entities.TermDetail, *errors.Error)
}
