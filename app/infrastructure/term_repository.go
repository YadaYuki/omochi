package infrastructure

import (
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/repository"
	"github.com/YadaYuki/omochi/app/ent"
)

type TermRepository struct {
	db *ent.Client
}

func NewTermRepository(db *ent.Client) repository.ITermRepository {
	return &TermRepository{db: db}
}

func (r *TermRepository) FindTermById(uuid string) (*entities.Term, error) {
	return &entities.Term{}, nil
}
