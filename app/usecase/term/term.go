package term

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/google/uuid"
)

type TermUseCase interface {
	FindTermById(ctx context.Context, id uuid.UUID) (*entities.TermDetail, *errors.Error)
}

type termUseCase struct {
	r repository.TermRepository
}

func NewTermUseCase(repository repository.TermRepository) TermUseCase {
	return &termUseCase{r: repository}
}

func (u *termUseCase) FindTermById(ctx context.Context, id uuid.UUID) (*entities.TermDetail, *errors.Error) {
	term, err := u.r.FindTermById(ctx, id)
	if err != nil {
		return nil, err
	}
	return term, nil
}
