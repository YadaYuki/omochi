package entdb

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/YadaYuki/omochi/app/common/slices"
	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent/enttest"
)

func TestFindTermById(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termRepository := NewTermEntRepository(client)
	testCases := []struct {
		word string
	}{
		{"sample"},
	}
	for _, tc := range testCases {
		termCreated, _ := client.Term.
			Create().
			SetWord(tc.word).
			Save(context.Background())
		term, err := termRepository.FindTermById(context.Background(), termCreated.ID)
		if err != nil {
			t.Fatal(err)
		}
		if term.Word != tc.word {
			t.Fatalf("expected %s, but got %s", tc.word, term.Word)
		}
	}
}

func TestFindTermCompressedsByWords(t *testing.T) {

	dummyInvertIndexCompressedCreate := entities.NewInvertIndexCompressedCreate([]byte("DUMMY INVERT INDEX COMPRESSED"))
	testCases := []struct {
		wordsForQuery []string
		wordsToInsert []string
		wordsToFind   []string // wordsForQueryとwordsToInsertの積集合になる.
	}{
		{
			wordsToInsert: []string{"hoge", "fuga", "piyo"},
			wordsForQuery: []string{"hoge", "piyo"},
			wordsToFind:   []string{"hoge", "piyo"},
		},
		{
			wordsToInsert: []string{"ruby", "js", "java", "python"},
			wordsForQuery: []string{"ruby", "js", "cpp"},
			wordsToFind:   []string{"ruby", "js"},
		},
		{
			wordsToInsert: []string{"ruby", "js", "java", "python"},
			wordsForQuery: []string{"cpp"},
			wordsToFind:   []string{},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			termRepository := NewTermEntRepository(client)
			for _, word := range tc.wordsToInsert {
				termCreated, _ := client.Term.
					Create().
					SetWord(word).
					Save(context.Background())
				client.InvertIndexCompressed.
					Create().
					SetPostingListCompressed(dummyInvertIndexCompressedCreate.PostingListCompressed).
					SetTermRelatedID(termCreated.ID).
					Save(context.Background())
			}
			termCompresseds, err := termRepository.FindTermCompressedsByWords(context.Background(), &tc.wordsForQuery)
			if err != nil {
				t.Fatal(err)
			}
			if len(tc.wordsToFind) != len(*termCompresseds) {
				t.Fatalf("len(*term) should be %v,but got %v", len(tc.wordsToFind), len(*termCompresseds))
			}
			for _, term := range *termCompresseds {
				if !slices.Contains(tc.wordsToFind, term.Word) {
					t.Fatalf("%v does not contain %v", tc.wordsToFind, term.Word)
				}
				if !bytes.Equal(dummyInvertIndexCompressedCreate.PostingListCompressed, term.InvertIndexCompressd.PostingListCompressed) {
					t.Fatalf("")
				}
			}
		})
	}
}
