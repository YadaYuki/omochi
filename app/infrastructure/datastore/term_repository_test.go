package infrastructure

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/YadaYuki/omochi/app/ent/enttest"
)

func TestFindTermById(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termRepository := NewTermRepository(client)
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
