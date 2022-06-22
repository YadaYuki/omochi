package entdb

import (
	"context"
	"fmt"
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
		doc := entities.NewDocumentCreate(tc.content, tc.tokenizedContent)
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

func TestFindDocumentsByIds(t *testing.T) {
	testCases := []struct {
		documentsCreate []*entities.DocumentCreate
		ids             []int64
		expectedContent []string
	}{
		{
			[]*entities.DocumentCreate{entities.NewDocumentCreate("hoge hoge hoge", []string{"hoge", "hoge", "hoge"}), entities.NewDocumentCreate("fuga fuga fuga", []string{"fuga", "fuga", "fuga"}), entities.NewDocumentCreate("piyo piyo piyo", []string{"piyo", "piyo", "piyo"})},
			[]int64{1, 2, 3},
			[]string{"hoge hoge hoge", "fuga fuga fuga", "piyo piyo piyo"},
		},
		{
			[]*entities.DocumentCreate{entities.NewDocumentCreate("hoge hoge hoge", []string{"hoge", "hoge", "hoge"}), entities.NewDocumentCreate("fuga fuga fuga", []string{"fuga", "fuga", "fuga"}), entities.NewDocumentCreate("piyo piyo piyo", []string{"piyo", "piyo", "piyo"})},
			[]int64{1, 3},
			[]string{"hoge hoge hoge", "piyo piyo piyo"},
		},
		{
			[]*entities.DocumentCreate{entities.NewDocumentCreate("hoge hoge hoge", []string{"hoge", "hoge", "hoge"}), entities.NewDocumentCreate("fuga fuga fuga", []string{"fuga", "fuga", "fuga"}), entities.NewDocumentCreate("piyo piyo piyo", []string{"piyo", "piyo", "piyo"})},
			[]int64{1, 2},
			[]string{"hoge hoge hoge", "fuga fuga fuga"},
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			documentRepository := NewDocumentEntRepository(client)
			ctx := context.Background()
			for _, documentCreate := range tc.documentsCreate {
				documentRepository.CreateDocument(ctx, documentCreate)
			}
			documents, findErr := documentRepository.FindDocumentsByIds(ctx, &tc.ids)
			if findErr != nil {
				t.Fatal(findErr)
			}
			if len(documents) != len(tc.expectedContent) {
				t.Fatalf("len(documents) should be %v, but got %v", len(tc.expectedContent), len(documents))
			}
			for i, doc := range documents {
				if doc.Content != tc.expectedContent[i] {
					t.Fatalf("doc.Content should be %v, but got %v", tc.expectedContent[i], doc.Content)
				}
			}
		})
	}
}
