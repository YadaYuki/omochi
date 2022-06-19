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
		invertIndex *entities.InvertIndex
	}{
		{invertIndex: entities.NewInvertIndex(uuid.New(), &[]entities.Posting{{DocumentRelatedId: -1, PositionsInDocument: []int{1, 2, 3}}})},
	}
	for _, tc := range testCases {
		compresser := NewGobInvertIndexCompresser()
		t.Run(fmt.Sprintf("%v", tc.invertIndex), func(tt *testing.T) {
			_, err := compresser.Compress(context.Background(), tc.invertIndex)
			if err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}
