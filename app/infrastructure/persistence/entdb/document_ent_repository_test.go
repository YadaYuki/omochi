package entdb

import (
	"context"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent/document"
	"github.com/YadaYuki/omochi/app/ent/enttest"
)

func TestCreateDocument(t *testing.T) {
	// TODO: Migrate sqlite3 to config
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	documentRepository := NewDocumentEntRepository(client)
	testCases := []struct {
		content          string
		tokenizedContent []string
	}{
		{"hoge hoge hoge", []string{"hoge", "hoge", "hoge"}},
	}
	for _, tc := range testCases {
		doc := entities.NewDocument(tc.content, tc.tokenizedContent)
		documentDetail, err := documentRepository.CreateDocument(context.Background(), doc)
		if err != nil {
			t.Fatal(err)
		}
		d, _ := client.Document.Query().Where(document.ID(int(documentDetail.Id))).Only(context.Background())
		if d.Content != tc.content {
			t.Fatalf("expected %s, but got %s", tc.content, d.Content)
		}
	}
}
