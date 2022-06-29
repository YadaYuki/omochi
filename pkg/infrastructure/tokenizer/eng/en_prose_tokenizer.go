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
	EnIndexableTokenPOSPrefix := []string{
		"JJ", "MD", "NN", "PDT", "PRP", "RB", "RPP", "UH", "VB", "WP", "WRB",
	}
	terms := []entities.TermCreate{}
	for _, token := range doc.Tokens() {
		indexableToken := false
		for _, prefix := range EnIndexableTokenPOSPrefix {
			if strings.HasPrefix(token.Tag, prefix) {
				indexableToken = true
			}
		}
		if indexableToken {
			terms = append(terms, *entities.NewTermCreate(strings.ToLower(token.Text), nil))
		}
	}
	return &terms, nil
}
