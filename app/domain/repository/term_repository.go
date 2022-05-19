package repository

import "github.com/YadaYuki/omochi/app/domain/entities"

type ITermRepository interface {
	FindTermById(uuid string) (*entities.Term, error)
}
