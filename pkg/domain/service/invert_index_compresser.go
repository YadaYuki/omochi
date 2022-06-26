package service

import (
	"context"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
)

type InvertIndexCompresser interface {
	Compress(ctx context.Context, invertIndexes *entities.InvertIndex) (*entities.InvertIndexCompressed, *errors.Error)
	Decompress(ctx context.Context, invertIndexes *entities.InvertIndexCompressed) (*entities.InvertIndex, *errors.Error)
}
