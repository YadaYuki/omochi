package indexer

import (
	"context"
	"fmt"
	"testing"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/ent/enttest"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/eng"

	_ "github.com/mattn/go-sqlite3"
)

//

func TestIndexingDocument(t *testing.T) {

	// Define Deps
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	documentRepository := entdb.NewDocumentEntRepository(client)
	termRepository := entdb.NewTermEntRepository(client)
	tokenizer := eng.NewEnProseTokenizer()
	invertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
	indexer := NewIndexer(documentRepository, termRepository, tokenizer, invertIndexCompresser)

	testCases := []struct {
		content string
	}{
		{"hoge hoge hoge fuga fuga fuga piyo piyo piyo"},
		{"hoge hoge hoge fuga fuga fuga piyo piyo piyo"},
		{"hoge hoge hoge fuga fuga fuga piyo piyo piyo"},
		{"hoge hoge hoge fuga fuga fuga piyo piyo piyo"},
	}
	for _, tc := range testCases {
		doc := entities.NewDocumentCreate(tc.content, []string{})
		indexingErr := indexer.IndexingDocument(context.Background(), doc)
		if indexingErr != nil {
			t.Fatal(indexingErr)
		}
	}
	a, _ := client.Term.Query().All(context.Background())
	c := compresser.NewZlibInvertIndexCompresser()
	invertIdxCps := entities.NewInvertIndexCompressed(a[0].PostingListCompressed)
	invertIndex, _ := c.Decompress(context.Background(), invertIdxCps)
	fmt.Println(*invertIndex.PostingList)
}
