package service

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type Tokenizer interface {
	Tokenize(ctx context.Context, content string) (*[]entities.TermCreate, *errors.Error)
}
