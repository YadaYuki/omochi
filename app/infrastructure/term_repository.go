package infrastructure

import (
	"context"
	"fmt"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/term"
	"github.com/google/uuid"
)

type TermRepository struct {
	db *ent.Client
}

func NewTermRepository(db *ent.Client) repository.ITermRepository {
	return &TermRepository{db: db}
}

func (r *TermRepository) FindTermById(ctx context.Context, id uuid.UUID) (*entities.Term, error) {
	term, err := r.db.Term.Query().Where(term.ID(uuid.UUID(id))).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to find term by id: %w", err)
	}
	return convertEntSchemaToEntity(term), nil
}

func convertEntSchemaToEntity(t *ent.Term) *entities.Term {
	return &entities.Term{
		Uuid:      t.ID,
		Word:      t.Word,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
