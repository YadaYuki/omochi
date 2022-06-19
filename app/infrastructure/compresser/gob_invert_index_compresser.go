package compresser

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type GobInvertIndexCompresser struct {
}

func NewGobInvertIndexCompresser() *GobInvertIndexCompresser {
	return &GobInvertIndexCompresser{}
}

func (c *GobInvertIndexCompresser) Compress(ctx context.Context, invertIndexes *[]entities.InvertIndexDetail) (*[]entities.InvertedIndexCompressedDetail, *errors.Error) {
	return nil, nil
}
func (c *GobInvertIndexCompresser) Decompress(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressedDetail) (*[]entities.InvertIndexDetail, *errors.Error) {
	return nil, nil
}
