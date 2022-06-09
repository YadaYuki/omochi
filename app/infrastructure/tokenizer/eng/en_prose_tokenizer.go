package eng

import (
	"context"

	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
	"github.com/jdkato/prose/v2"

	"github.com/YadaYuki/omochi/app/domain/entities"
)

type EnProseTokenizer struct{}

func NewEnProseTokenizer() *EnProseTokenizer {
	return &EnProseTokenizer{}
}

func (tokenizer *EnProseTokenizer) Tokenize(ctx context.Context, content string) (*[]entities.Term, *errors.Error) {
	doc, err := prose.NewDocument(content)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}

	terms := []entities.Term{}
	for _, token := range doc.Tokens() {
		terms = append(terms, *entities.NewTerm(token.Text))
	}
	return &terms, nil
}
