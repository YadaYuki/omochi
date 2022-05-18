package repository

import "github.com/YadaYuki/omochi/app/domain/entities"

type TermRepository interface {
	FindTermById(uuid string) (*entities.Term, error)
}
