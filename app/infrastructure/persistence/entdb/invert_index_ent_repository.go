package entdb

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type InvertIndexCompressedEntRepository struct {
	db *ent.Client
}

func NewInvertIndexCompressedEntRepository(db *ent.Client) repository.InvertIndexCompressedRepository {
	return &InvertIndexCompressedEntRepository{db: db}
}

func (r *InvertIndexCompressedEntRepository) BulkCreateInvertIndexesCompressed(ctx context.Context, invertIndexes *[]entities.InvertIndexCompressedCreate) (*[]entities.InvertIndexCompressed, *errors.Error) {
	bulk := make([]*ent.InvertIndexCompressedCreate, len(*invertIndexes))
	for i, invertIndex := range *invertIndexes {
		bulk[i] = r.db.InvertIndexCompressed.Create().SetTermRelatedID(invertIndex.TermId).SetPostingListCompressed(invertIndex.PostingListCompressed)
	}
	invertIndexesCreated, err := r.db.InvertIndexCompressed.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertInvertIndexesEntSchemaToEntity(invertIndexesCreated), nil
}

func convertInvertIndexesEntSchemaToEntity(t []*ent.InvertIndexCompressed) *[]entities.InvertIndexCompressed {
	invertIndexes := make([]entities.InvertIndexCompressed, len(t))
	for i, item := range t {
		invertIndexes[i] = entities.InvertIndexCompressed{
			Uuid:                  item.ID,
			PostingListCompressed: item.PostingListCompressed,
			CreatedAt:             item.CreatedAt,
			UpdatedAt:             item.UpdatedAt,
		}
	}
	return &invertIndexes
}
