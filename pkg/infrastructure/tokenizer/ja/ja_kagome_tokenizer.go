package ja

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

type JaKagomeTokenizer struct {
	t *tokenizer.Tokenizer
}

func NewJaKagomeTokenizer() service.Tokenizer {
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		panic(err)
	}
	return &JaKagomeTokenizer{t: t}
}

func (tokenizer *JaKagomeTokenizer) Tokenize(ctx context.Context, japaneseContent string) (*[]entities.TermCreate, *errors.Error) {
	tokens := tokenizer.t.Tokenize(japaneseContent)
	var JaIndexableTokenPOS map[string]bool = map[string]bool{"感動詞": true, "形容詞": true, "動詞": true, "名詞": true, "副詞": true}
	terms := []entities.TermCreate{}
	for _, token := range tokens {
		POS := token.Features()[0]
		if _, ok := JaIndexableTokenPOS[POS]; ok {
			terms = append(terms, *entities.NewTermCreate(token.Surface, nil))
		}
	}
	return &terms, nil
}
