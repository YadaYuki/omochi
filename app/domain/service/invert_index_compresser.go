package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type InvertIndexCompresser interface {
	Compress(ctx context.Context, invertIndexes *[]entities.InvertIndexDetail) (*[]entities.InvertedIndexCompressedDetail, *errors.Error)
	Decompress(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressedDetail) (*[]entities.InvertIndexDetail, *errors.Error)
}
