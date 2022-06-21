package entdb

import (
	"context"
	"log"
	"reflect"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent/enttest"
	"github.com/google/uuid"
)

func TestBulkCreateInvertIndexesCompressed(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	invertedIndexCompressedRepository := NewInvertedIndexCompressedEntRepository(client)
	dummyContent := []byte("DUMMY POSTING LIST COMPRESSED")
	testCases := []struct {
		postingListsCompressed [][]byte
	}{
		{postingListsCompressed: [][]byte{dummyContent, dummyContent, dummyContent, dummyContent, dummyContent}},
	}
	for _, tc := range testCases {
		invertIndexesCompressed := make([]entities.InvertIndexCompressedCreate, 0)
		// 転置インデックスと紐づくタームを事前作成する.
		ctx := context.Background()
		for _, postingListCompressed := range tc.postingListsCompressed {
			wordDummy := uuid.NewString()
			term := entities.NewTerm(wordDummy)
			termCreated, _ := client.Term.Create().SetWord(term.Word).Save(ctx)
			invertIndexesCompressed = append(invertIndexesCompressed, *entities.NewInvertIndexCompressedCreate(termCreated.ID, postingListCompressed))
		}
		_, err := invertedIndexCompressedRepository.BulkCreateInvertIndexesCompressed(ctx, &invertIndexesCompressed)
		if err != nil {
			t.Fatal(err)
		}
		d, allErr := client.InvertIndexCompressed.Query().All(ctx)
		if allErr != nil {
			t.Fatalf(allErr.Error())
		}
		if len(d) != len(tc.postingListsCompressed) {
			t.Fatalf("len(d) should be %v but got %v", len(tc.postingListsCompressed), len(d))
		}
		for _, item := range d {
			if !reflect.DeepEqual(item.PostingListCompressed, []byte("DUMMY POSTING LIST COMPRESSED")) {
				t.Fatalf("")
			}
		}
	}
}
