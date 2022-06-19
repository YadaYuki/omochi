package entdb

import (
	"context"
	"reflect"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent/enttest"
	"github.com/google/uuid"
)

func TestBulkCreateInvertIndexesCompressed(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	invertedIndexCompressedRepository := NewInvertedIndexCompressedEntRepository(client)
	dummyContent := []byte("DUMMY POSTING LIST COMPRESSED")
	testCases := []struct {
		postingListsCompressed [][]byte
	}{
		{postingListsCompressed: [][]byte{dummyContent, dummyContent, dummyContent, dummyContent, dummyContent}},
	}
	for _, tc := range testCases {
		invertIndexesCompressed := make([]entities.InvertedIndexCompressed, 0)
		for _, postingListCompressed := range tc.postingListsCompressed {
			dummyUuid := uuid.New()
			invertIndexesCompressed = append(invertIndexesCompressed, *entities.NewInvertIndexCompressed(dummyUuid, postingListCompressed))
		}
		_, err := invertedIndexCompressedRepository.BulkCreateInvertIndexesCompressed(context.Background(), &invertIndexesCompressed)
		if err != nil {
			t.Fatal(err)
		}
		d, _ := client.InvertIndexCompressed.Query().All(context.Background())
		for _, item := range d {
			if !reflect.DeepEqual(item.PostingListCompressed, []byte("DUMMY POSTING LIST COMPRESSED")) {
				t.Fatalf("")
			}
		}
	}
}
