package usecase

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/google/uuid"
)

type ITermUseCase interface {
	FindTermById(ctx context.Context, id uuid.UUID) (*entities.Term, error)
}

type TermUseCase struct {
	r repository.ITermRepository
}

func NewTermUseCase(repository repository.ITermRepository) ITermUseCase {
	return &TermUseCase{r: repository}
}

func (u *TermUseCase) FindTermById(ctx context.Context, id uuid.UUID) (*entities.Term, error) {
	term, err := u.r.FindTermById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find term by id: %w", err)
	}
	return term, nil
}
