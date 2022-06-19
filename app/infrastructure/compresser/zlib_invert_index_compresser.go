package compresser

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/gob"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type ZlibInvertIndexCompresser struct {
}

func NewGobInvertIndexCompresser() *ZlibInvertIndexCompresser {
	return &ZlibInvertIndexCompresser{}
}

func (c *ZlibInvertIndexCompresser) Compress(ctx context.Context, invertIndex *entities.InvertIndex) (*entities.InvertedIndexCompressed, *errors.Error) {

	// Encode posting list to gob
	var postingListGobBuffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&postingListGobBuffer)
	gobEncodeErr := gobEncoder.Encode(invertIndex.PostingList)
	if gobEncodeErr != nil {
		return nil, errors.NewError(code.Unknown, gobEncodeErr)
	}

	// Compress posting list by zlib
	var compressedPostingListBuffer bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedPostingListBuffer)
	defer zlibWriter.Close()

	_, zlibError := zlibWriter.Write(postingListGobBuffer.Bytes())
	if zlibError != nil {
		return nil, errors.NewError(code.Unknown, zlibError)
	}
	compressedPostingList := compressedPostingListBuffer.Bytes()

	invertIndexCompressed := entities.NewInvertIndexCompressed(invertIndex.TermId, compressedPostingList)

	return invertIndexCompressed, nil
}

func (c *ZlibInvertIndexCompresser) Decompress(ctx context.Context, invertIndexes *[]entities.InvertedIndexCompressed) (*[]entities.InvertIndex, *errors.Error) {
	return nil, nil
}
