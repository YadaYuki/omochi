package infrastructure

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/YadaYuki/omochi/app/ent/enttest"
)

func TestFindTermById(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termRepository := NewTermRepository(client)

	termCreated, _ := client.Term.
		Create().
		SetWord("sample").
		Save(ctx)

	term, err := termRepository.FindTermById(termCreated.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("found term", term)
}
