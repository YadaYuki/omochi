package entdb

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/repository"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/ent/predicate"
	"github.com/YadaYuki/omochi/pkg/ent/term"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
	"github.com/google/uuid"
)

type TermEntRepository struct {
	db *ent.Client
}

func NewTermEntRepository(db *ent.Client) repository.TermRepository {
	return &TermEntRepository{db: db}
}

func (r *TermEntRepository) FindTermCompressedById(ctx context.Context, uuid uuid.UUID) (*entities.TermCompressed, *errors.Error) {
	term, err := r.db.Term.Query().Where(term.ID(uuid)).Only(ctx)
	if err != nil {
		_, ok := err.(*ent.NotFoundError)
		if ok {
			return nil, errors.NewError(code.NotExist, err)
		}
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertTermCompressedEntSchemaToEntity(term), nil
}

func (r *TermEntRepository) BulkUpsertTerm(ctx context.Context, terms *[]entities.TermCompressedCreate) *errors.Error {
	termCreates := make([]*ent.TermCreate, len(*terms))
	for i, term := range *terms {
		termCreates[i] = r.db.Term.Create().SetWord(term.Word).SetPostingListCompressed(term.InvertIndexCompressed.PostingListCompressed)
	}
	err := r.db.Term.
		CreateBulk(termCreates...).
		OnConflict().
		Update(func(tu *ent.TermUpsert) {
			tu.UpdatePostingListCompressed()
			tu.UpdateUpdatedAt()
		}).Exec(ctx)
	if err != nil {
		return errors.NewError(code.Unknown, err)
	}
	return nil
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
		All(ctx)
	if queryErr != nil {
		return nil, errors.NewError(code.Unknown, queryErr)
	}
	return convertTermCompressedsEntSchemaToEntity(termCompresseds), nil
}

func (r *TermEntRepository) FindTermCompressedByWord(ctx context.Context, word string) (*entities.TermCompressed, *errors.Error) {
	term, queryErr := r.db.Term.Query().Where(term.Word(word)).Only(ctx)
	if queryErr != nil {
		_, ok := queryErr.(*ent.NotFoundError)
		if ok {
			return nil, errors.NewError(code.NotExist, queryErr)
		}
		return nil, errors.NewError(code.Unknown, queryErr)
	}
	return convertTermCompressedEntSchemaToEntity(term), nil
}

func convertTermCompressedsEntSchemaToEntity(entTerms []*ent.Term) *[]entities.TermCompressed {
	termCompresseds := make([]entities.TermCompressed, len(entTerms))
	for i, entTerm := range entTerms {
		invertIndexCompressed := &entities.InvertIndexCompressed{
			PostingListCompressed: entTerm.PostingListCompressed,
		}
		termCompresseds[i] = entities.TermCompressed{
			Uuid:                  entTerm.ID,
			Word:                  entTerm.Word,
			InvertIndexCompressed: invertIndexCompressed,
			CreatedAt:             entTerm.CreatedAt,
			UpdatedAt:             entTerm.UpdatedAt,
		}
	}
	return &termCompresseds
}

func convertTermCompressedEntSchemaToEntity(entTerm *ent.Term) *entities.TermCompressed {
	invertIndexCompressed := &entities.InvertIndexCompressed{
		PostingListCompressed: entTerm.PostingListCompressed,
	}
	termCompressed := entities.TermCompressed{
		Uuid:                  entTerm.ID,
		Word:                  entTerm.Word,
		InvertIndexCompressed: invertIndexCompressed,
		CreatedAt:             entTerm.CreatedAt,
		UpdatedAt:             entTerm.UpdatedAt,
	}

	return &termCompressed
}
