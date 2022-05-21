package repository

import (
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/google/uuid"
)

type ITermRepository interface {
	FindTermById(uuid uuid.UUID) (*entities.Term, error)
}
