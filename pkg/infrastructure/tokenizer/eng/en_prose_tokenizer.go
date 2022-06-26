package eng

import (
	"context"
	"strings"

	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
	"github.com/jdkato/prose/v2"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
)

type EnProseTokenizer struct{}

func NewEnProseTokenizer() service.Tokenizer {
	return &EnProseTokenizer{}
}

func (tokenizer *EnProseTokenizer) Tokenize(ctx context.Context, content string) (*[]entities.TermCreate, *errors.Error) {
	doc, err := prose.NewDocument(content)
	if err != nil {
		return nil, errors.NewError(code.Unknown, err)
	}
	INDEXABLE_TOKEN_TAG_PREFIX := []string{
		"JJ", "MD", "NN", "PDT", "PRP", "RB", "RPP", "UH", "VB", "WP", "WRB",
	}
	terms := []entities.TermCreate{}
	for _, token := range doc.Tokens() {
		indexable_token := false
		for _, prefix := range INDEXABLE_TOKEN_TAG_PREFIX {
			if strings.HasPrefix(token.Tag, prefix) {
				indexable_token = true
			}
		}
		if indexable_token {
			terms = append(terms, *entities.NewTermCreate(strings.ToLower(token.Text), nil))
		}
	}
	return &terms, nil
}
