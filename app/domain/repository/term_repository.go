package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/google/uuid"
)

type ITermRepository interface {
	FindTermById(ctx context.Context, uuid uuid.UUID) (*entities.TermDetail, *errors.Error)
}
