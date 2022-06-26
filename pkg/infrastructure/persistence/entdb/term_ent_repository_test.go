package entdb

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/YadaYuki/omochi/pkg/common/slices"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/ent/enttest"
)

func TestFindTermCompressedById(t *testing.T) {
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
			SetPostingListCompressed([]byte("hoge")).
			Save(context.Background())
		term, err := termRepository.FindTermCompressedById(context.Background(), termCreated.ID)
		if err != nil {
			t.Fatal(err)
		}
		if term.Word != tc.word {
			t.Fatalf("expected %s, but got %s", tc.word, term.Word)
		}
	}
}

func TestFindTermCompressedByWord(t *testing.T) {

	testCases := []struct {
		word                  string
		postingListCompressed []byte
	}{
		{"sample", []byte("hoge")},
	}
	for _, tc := range testCases {
		t.Run(tc.word, func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			termRepository := NewTermEntRepository(client)
			ctx := context.Background()
			client.Term.
				Create().
				SetWord(tc.word).
				SetPostingListCompressed(tc.postingListCompressed).
				Save(ctx)
			term, err := termRepository.FindTermCompressedByWord(ctx, tc.word)
			if err != nil {
				t.Fatal(err)
			}
			if term.Word != tc.word {
				t.Fatalf("expected %s, but got %s", tc.word, term.Word)
			}
		})
	}
}

func TestFindTermCompressedsByWords(t *testing.T) {

	dummyInvertIndexCompressedCreate := entities.NewInvertIndexCompressed([]byte("DUMMY INVERT INDEX COMPRESSED"))
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
				client.Term.
					Create().
					SetWord(word).
					SetPostingListCompressed(dummyInvertIndexCompressedCreate.PostingListCompressed).
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
				if !bytes.Equal(dummyInvertIndexCompressedCreate.PostingListCompressed, term.InvertIndexCompressed.PostingListCompressed) {
					t.Fatalf("")
				}
			}
		})
	}
}

func TestBulkUpsertTerm(t *testing.T) {
	dummyInvertIndexCompressedCreate := entities.NewInvertIndexCompressed([]byte("DUMMY INVERT INDEX COMPRESSED"))
	dummyInvertIndexCompressedUpdate := entities.NewInvertIndexCompressed([]byte("DUMMY INVERT INDEX COMPRESSED UPDATED"))
	testCases := []struct {
		wordsForAdvanceInsert []string
		wordsToUpsert         []string
		wordsAfterUpsert      []string // wordsForQueryとwordsToInsertの和集合になる.
	}{
		{
			wordsForAdvanceInsert: []string{"hoge", "fuga"},
			wordsToUpsert:         []string{"hoge", "piyo"},
			wordsAfterUpsert:      []string{"hoge", "fuga", "piyo"},
		},
		{
			wordsForAdvanceInsert: []string{},
			wordsToUpsert:         []string{"ruby", "js", "cpp"},
			wordsAfterUpsert:      []string{"ruby", "js", "cpp"},
		},
		{
			wordsForAdvanceInsert: []string{"ruby", "js", "java", "python"},
			wordsToUpsert:         []string{"ruby", "js", "java", "python"},
			wordsAfterUpsert:      []string{"ruby", "js", "java", "python"},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(tt *testing.T) {
			ctx := context.Background()
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			termRepository := NewTermEntRepository(client)
			for _, word := range tc.wordsForAdvanceInsert {
				client.Term.
					Create().
					SetWord(word).
					SetPostingListCompressed(dummyInvertIndexCompressedCreate.PostingListCompressed).
					Save(ctx)
			}
			termsUpsert := make([]entities.TermCompressedCreate, len(tc.wordsToUpsert))
			for i := 0; i < len(tc.wordsToUpsert); i++ {
				term := entities.NewTermCompressedCreate(tc.wordsToUpsert[i], dummyInvertIndexCompressedUpdate)
				termsUpsert[i] = *term
			}
			err := termRepository.BulkUpsertTerm(ctx, &termsUpsert)
			if err != nil {
				t.Fatal(err)
			}
			entTerms, _ := client.
				Term.
				Query().
				All(ctx)
			if len(tc.wordsAfterUpsert) != len(entTerms) {
				t.Fatalf("len(entTerms) should be %v,but got %v", len(tc.wordsAfterUpsert), len(entTerms))
			}
			for _, entTerm := range entTerms {
				if !slices.Contains(tc.wordsAfterUpsert, entTerm.Word) {
					t.Fatalf("%v does not contain %v", tc.wordsAfterUpsert, entTerm.Word)
				}
				if slices.Contains(tc.wordsToUpsert, entTerm.Word) {
					if !bytes.Equal(dummyInvertIndexCompressedUpdate.PostingListCompressed, entTerm.PostingListCompressed) {
						t.Fatalf("PostingListCompressed after update should be %v. but got %v", string(dummyInvertIndexCompressedUpdate.PostingListCompressed), string(entTerm.PostingListCompressed))
					}
				}
			}
		})
	}
}
