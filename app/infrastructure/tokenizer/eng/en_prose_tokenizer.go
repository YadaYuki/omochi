package eng

import (
	"context"

	"github.com/YadaYuki/omochi/app/errors"

	"github.com/YadaYuki/omochi/app/domain/entities"
)

type EnProseTokenizer struct{}

func NewEnProseTokenizer() *EnProseTokenizer {
	return &EnProseTokenizer{}
}

func (tokenizer *EnProseTokenizer) Tokenize(ctx context.Context, content string) (*[]entities.Term, *errors.Error) {
	return nil, nil
}
