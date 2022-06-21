package compresser

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/domain/service"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
)

type ZlibInvertIndexCompresser struct {
}

func NewZlibInvertIndexCompresser() service.InvertIndexCompresser {
	return &ZlibInvertIndexCompresser{}
}

func (c *ZlibInvertIndexCompresser) Compress(ctx context.Context, invertIndex *entities.InvertIndex) (*entities.InvertIndexCompressed, *errors.Error) {

	// Encode posting list to gob
	var postingListGobBuffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&postingListGobBuffer)
	postings := make([]entities.Posting, 0)
	for i := 0; i < len(*invertIndex.PostingList); i++ {
		postings = append(postings, (*invertIndex.PostingList)[i])
	}
	gobEncodeErr := gobEncoder.Encode(&postings)
	if gobEncodeErr != nil {
		return nil, errors.NewError(code.Unknown, gobEncodeErr)
	}

	// Compress posting list by zlib
	var compressedPostingListBuffer bytes.Buffer
	zlibWriter := zlib.NewWriter(&compressedPostingListBuffer)
	_, zlibError := zlibWriter.Write(postingListGobBuffer.Bytes())
	if zlibError != nil {
		return nil, errors.NewError(code.Unknown, zlibError)
	}
	defer zlibWriter.Close()
	flushErr := zlibWriter.Flush() // compressedPostingListBufferに圧縮したデータを全て書き込む
	if flushErr != nil {
		return nil, errors.NewError(code.Unknown, flushErr)
	}
	compressedPostingList := compressedPostingListBuffer.Bytes()

	invertIndexCompressed := entities.NewInvertIndexCompressed(compressedPostingList)

	return invertIndexCompressed, nil
}

func (c *ZlibInvertIndexCompresser) Decompress(ctx context.Context, invertIndex *entities.InvertIndexCompressed) (*entities.InvertIndex, *errors.Error) {

	// decompress posting list by zlib
	compressedPostingListBuffer := bytes.NewBuffer(invertIndex.PostingListCompressed)
	zlibReader, zlibError := zlib.NewReader(compressedPostingListBuffer)
	if zlibError != nil {
		return nil, errors.NewError(code.Unknown, fmt.Sprintf("zlib: %v", zlibError.Error()))
	}
	var decompressedDataBuffer bytes.Buffer
	io.Copy(&decompressedDataBuffer, zlibReader)

	// Decode gob to PostingList
	var postingList []entities.Posting
	gobDecoder := gob.NewDecoder(&decompressedDataBuffer)
	gobDecodeErr := gobDecoder.Decode(&postingList)
	if gobDecodeErr != nil {
		return nil, errors.NewError(code.Unknown, fmt.Sprintf("gob: %v", gobDecodeErr.Error()))
	}
	invertIndexes := entities.NewInvertIndex(&postingList)
	return invertIndexes, nil
}
