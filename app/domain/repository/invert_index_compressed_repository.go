package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type InvertedIndexCompressedRepository interface {
	BulkCreateInvertIndexesCompressed(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressed) (*[]entities.InvertedIndexCompressedDetail, *errors.Error)
}
