package eng

import (
	"context"
	"strings"

	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
	"github.com/jdkato/prose/v2"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"
)

type EnProseTokenizer struct{}

func NewEnProseTokenizer() service.Tokenizer {
	return &EnProseTokenizer{}
}

func (tokenizer *EnProseTokenizer) Tokenize(ctx context.Context, content string) (*[]entities.Term, *errors.Error) {
	doc, err := prose.NewDocument(content)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	INDEXABLE_TOKEN_TAG_PREFIX := []string{
		"JJ", "MD", "NN", "PDT", "PRP", "RB", "RPP", "UH", "VB", "WP", "WRB",
	}
	terms := []entities.Term{}
	for _, token := range doc.Tokens() {
		indexable_token := true
		for _, prefix := range INDEXABLE_TOKEN_TAG_PREFIX {
			if !strings.HasPrefix(token.Tag, prefix) {
				indexable_token = false
			}
		}
		if indexable_token {
			terms = append(terms, *entities.NewTerm(strings.ToLower(token.Text)))
		}
	}
	return &terms, nil
}
