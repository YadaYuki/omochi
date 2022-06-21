package repository

import (
	"context"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
)

type InvertIndexCompressedRepository interface {
	BulkCreateInvertIndexesCompressed(ctx context.Context, invertIndexes *[]entities.InvertIndexCompressedCreate) (*[]entities.InvertIndexCompressed, *errors.Error)
}
