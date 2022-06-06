package entdb

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/term"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
	"github.com/google/uuid"
)

type TermEntRepository struct {
	db *ent.Client
}

func NewTermEntRepository(db *ent.Client) repository.ITermRepository {
	return &TermEntRepository{db: db}
}

func (r *TermEntRepository) FindTermById(ctx context.Context, id uuid.UUID) (*entities.TermDetail, *errors.Error) {
	term, err := r.db.Term.Query().Where(term.ID(uuid.UUID(id))).Only(ctx)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			return nil, errors.NewError(code.NotExist, err)
		}
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertEntSchemaToEntity(term), nil
}

func convertEntSchemaToEntity(t *ent.Term) *entities.TermDetail {
	return &entities.TermDetail{
		Uuid:      t.ID,
		Word:      t.Word,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
