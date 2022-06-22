package term

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/google/uuid"
)

type TermUseCase interface {
	FindTermCompressedById(ctx context.Context, id uuid.UUID) (*entities.TermCompressed, *errors.Error)
}

type termUseCase struct {
	r repository.TermRepository
}

func NewTermUseCase(repository repository.TermRepository) TermUseCase {
	return &termUseCase{r: repository}
}

func (u *termUseCase) FindTermCompressedById(ctx context.Context, id uuid.UUID) (*entities.TermCompressed, *errors.Error) {
	term, err := u.r.FindTermCompressedById(ctx, id)
	if err != nil {
		return nil, err
	}
	return term, nil
}
