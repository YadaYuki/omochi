package service

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type InvertIndexCompresser interface {
	Compress(ctx context.Context, invertIndexes *entities.InvertIndexCreate) (*entities.InvertIndexCompressedCreate, *errors.Error)
	Decompress(ctx context.Context, invertIndexes *entities.InvertIndexCompressedCreate) (*entities.InvertIndexCreate, *errors.Error)
}
