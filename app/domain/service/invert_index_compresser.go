package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type InvertIndexCompresser interface {
	Compress(ctx context.Context, invertIndexes *[]entities.InvertIndex) (*[]entities.InvertedIndexCompressed, *errors.Error)
	Decompress(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressed) (*[]entities.InvertIndex, *errors.Error)
}
