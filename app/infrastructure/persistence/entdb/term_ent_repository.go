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

func NewTermEntRepository(db *ent.Client) repository.TermRepository {
	return &TermEntRepository{db: db}
}

func (r *TermEntRepository) FindTermById(ctx context.Context, uuid uuid.UUID) (*entities.Term, *errors.Error) {
	term, err := r.db.Term.Query().Where(term.ID(uuid)).Only(ctx)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			return nil, errors.NewError(code.NotExist, err)
		}
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertTermEntSchemaToEntity(term), nil
}

func (r *TermEntRepository) BulkUpsertTerm(ctx context.Context, terms *[]entities.TermCreate) (*[]entities.Term, *errors.Error) {
	return nil, nil
}

func (r *TermEntRepository) FindTermsByWords(ctx context.Context, words *[]string) (*[]entities.Term, *errors.Error) {
	return nil, nil
}

func convertTermEntSchemaToEntity(t *ent.Term) *entities.Term {
	return &entities.Term{
		Uuid:      t.ID,
		Word:      t.Word,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
