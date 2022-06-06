package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type Tokenizer interface {
	Tokenize(ctx context.Context, content string) (*[]entities.Term, *errors.Error)
}
