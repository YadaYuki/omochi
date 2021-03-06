package entdb

import (
	"context"
	"strings"

	"github.com/YadaYuki/omochi/pkg/common/constant"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/repository"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/ent/document"
	"github.com/YadaYuki/omochi/pkg/ent/predicate"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
)

type DocumentEntRepository struct {
	db *ent.Client
}

func NewDocumentEntRepository(db *ent.Client) repository.DocumentRepository {
	return &DocumentEntRepository{db: db}
}

func (r *DocumentEntRepository) CreateDocument(ctx context.Context, doc *entities.DocumentCreate) (*entities.Document, *errors.Error) {
	docCreated, err := r.db.Document.
		Create().
		SetContent(doc.Content).
		SetTokenizedContent(strings.Join(doc.TokenizedContent, constant.WhiteSpace)).
		Save(ctx)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	return convertDocumentEntSchemaToEntity(docCreated), nil
}

func (r *DocumentEntRepository) FindDocumentsByIds(ctx context.Context, ids *[]int64) ([]*entities.Document, *errors.Error) {
	predicatesForIds := make([]predicate.Document, len(*ids))
	for i, id := range *ids {
		predicatesForIds[i] = document.ID(int(id))
	}
	documents, queryErr := r.
		db.
		Document.
		Query().
		Where(document.Or(predicatesForIds...)).
		All(ctx)
	if queryErr != nil {
		return nil, errors.NewError(code.Unknown, queryErr)
	}
	return convertDocumentsEntSchemaToEntity(documents), nil
}

func convertDocumentsEntSchemaToEntity(t []*ent.Document) []*entities.Document {
	documents := []*entities.Document{}
	for _, entDocument := range t {
		documents = append(documents, convertDocumentEntSchemaToEntity(entDocument))
	}
	return documents
}

func convertDocumentEntSchemaToEntity(t *ent.Document) *entities.Document {
	return &entities.Document{
		Id:               int64(t.ID),
		Content:          t.Content,
		TokenizedContent: strings.Split(t.TokenizedContent, constant.WhiteSpace),
		CreatedAt:        t.CreatedAt,
		UpdatedAt:        t.UpdatedAt,
	}
}
