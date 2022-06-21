package entdb

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/predicate"
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

func (r *TermEntRepository) BulkUpsertTerm(ctx context.Context, terms *[]entities.TermCompressedCreate) (*[]entities.TermCompressed, *errors.Error) {
	return nil, nil
}

//
func (r *TermEntRepository) FindTermCompressedsByWords(ctx context.Context, words *[]string) (*[]entities.TermCompressed, *errors.Error) {
	predicatesForWords := make([]predicate.Term, len(*words))
	for i, word := range *words {
		predicatesForWords[i] = term.Word(word)
	}
	termCompresseds, queryErr := r.
		db.
		Term.
		Query().
		Where(term.Or(predicatesForWords...)).
		WithInvertIndexCompressed().
		All(ctx)
	if queryErr != nil {
		return nil, errors.NewError(code.Unknown, queryErr)
	}
	return convertTermCompressedsEntSchemaToEntity(termCompresseds), nil
}

func convertTermCompressedsEntSchemaToEntity(entTerms []*ent.Term) *[]entities.TermCompressed {
	termCompresseds := make([]entities.TermCompressed, len(entTerms))
	for i, entTerm := range entTerms {
		invertIndexCompressed := &entities.InvertIndexCompressed{
			Uuid:                  entTerm.Edges.InvertIndexCompressed.ID,
			TermId:                entTerm.ID,
			PostingListCompressed: entTerm.Edges.InvertIndexCompressed.PostingListCompressed,
			CreatedAt:             entTerm.Edges.InvertIndexCompressed.CreatedAt,
			UpdatedAt:             entTerm.Edges.InvertIndexCompressed.UpdatedAt,
		}
		termCompresseds[i] = entities.TermCompressed{
			Uuid:                 entTerm.ID,
			Word:                 entTerm.Word,
			InvertIndexCompressd: invertIndexCompressed,
			CreatedAt:            entTerm.CreatedAt,
			UpdatedAt:            entTerm.UpdatedAt,
		}
	}
	return &termCompresseds
}

func convertTermEntSchemaToEntity(t *ent.Term) *entities.Term {
	return &entities.Term{
		Uuid:      t.ID,
		Word:      t.Word,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
