package compresser

import (
	"context"
	"fmt"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/google/uuid"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		invertIndex *entities.InvertIndexCreate
	}{
		{invertIndex: entities.NewInvertIndexCreate(uuid.New(), &[]entities.Posting{{DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}})},
	}
	for _, tc := range testCases {
		compresser := NewGobInvertIndexCompresser()
		t.Run(fmt.Sprintf("%v", tc.invertIndex), func(tt *testing.T) {
			compressed, err := compresser.Compress(context.Background(), tc.invertIndex)
			if err != nil {
				t.Fatalf(err.Error())
			}
			if len(compressed.PostingListCompressed) <= 0 {
				t.Fatalf("compressed PostingList should be longer than 0")
			}
		})
	}
}

// E2E
func TestCompressToDecompress(t *testing.T) {
	testCases := []struct {
		invertIndex *entities.InvertIndexCreate
	}{
		// {invertIndex: entities.NewInvertIndex(uuid.New(), &[]entities.Posting{{DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}})},
		{invertIndex: entities.NewInvertIndexCreate(uuid.New(), &[]entities.Posting{{DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}, {DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}})},
	}
	for _, tc := range testCases {
		compresser := NewGobInvertIndexCompresser()
		t.Run(fmt.Sprintf("%v", tc.invertIndex), func(tt *testing.T) {
			ctx := context.Background()
			invertIndexCompressed, compressErr := compresser.Compress(ctx, tc.invertIndex)
			if compressErr != nil {
				t.Fatalf(compressErr.Error())
			}
			invertIndexDecompressed, decompressErr := compresser.Decompress(ctx, invertIndexCompressed)
			if decompressErr != nil {
				t.Fatalf(decompressErr.Error())
			}
			for i, postingDecompressed := range *invertIndexDecompressed.PostingList {
				if postingDecompressed.DocumentRelatedId != (*tc.invertIndex.PostingList)[i].DocumentRelatedId {
					t.Fatalf("postingDecompressed.DocumentRelatedId should be %v, but got %v ", (*tc.invertIndex.PostingList)[i].DocumentRelatedId, postingDecompressed.DocumentRelatedId)
				}
				for j, positionInDocDecompressed := range postingDecompressed.PositionsInDocument {
					if positionInDocDecompressed != (*tc.invertIndex.PostingList)[i].PositionsInDocument[j] {
						t.Fatalf("positionInDocDecompressed should be %v, but got %v ", positionInDocDecompressed, (*tc.invertIndex.PostingList)[i].PositionsInDocument[j])
					}
				}
			}
		})
	}
}
