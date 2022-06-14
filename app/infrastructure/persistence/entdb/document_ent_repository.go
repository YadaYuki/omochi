package entdb

import (
	"context"
	"strings"

	"github.com/YadaYuki/omochi/app/common/constant"
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type DocumentEntRepository struct {
	db *ent.Client
}

func NewDocumentEntRepository(db *ent.Client) repository.DocumentRepository {
	return &DocumentEntRepository{db: db}
}

func (r *DocumentEntRepository) CreateDocument(ctx context.Context, doc *entities.Document) (*entities.DocumentDetail, *errors.Error) {
	docCreated, err := r.db.Document.
		Create().
		SetContent(doc.Content).
		SetTokenizedContent(strings.Join(doc.TokenizedContent, constant.WHITE_SPACE)).
		Save(ctx)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertDocumentEntSchemaToEntity(docCreated), nil
}

func convertDocumentEntSchemaToEntity(t *ent.Document) *entities.DocumentDetail {
	return &entities.DocumentDetail{
		Id:               int64(t.ID),
		Content:          t.Content,
		TokenizedContent: strings.Split(t.TokenizedContent, constant.WHITE_SPACE),
		CreatedAt:        t.CreatedAt,
		UpdatedAt:        t.UpdatedAt,
	}
}
