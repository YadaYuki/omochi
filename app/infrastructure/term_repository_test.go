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

	termCreated, err := client.Term.
		Create().
		SetWord("sample").
		Save(ctx)
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}
	t.Log("created term", termCreated)

	term, err := termRepository.FindTermById(termCreated.ID.String())
	if err != nil {
		t.Fatal("failed finding user using UserService", err)
	}
	if term.Word != "sample" {
		t.Fatal("failed finding user using UserService")
	}
}
