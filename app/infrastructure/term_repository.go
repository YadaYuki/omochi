package infrastructure

import "github.com/YadaYuki/omochi/app/ent"

type TermRepository struct {
	db *ent.Client
}

func NewTermRepository(db *ent.Client) *TermRepository {
	return &TermRepository{db: db}
}
