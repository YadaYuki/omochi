package searcher

import (
	"context"
	"strings"
	"testing"

	"github.com/YadaYuki/omochi/pkg/common/constant"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/ent/enttest"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/documentranker/tfidfranker"
	"github.com/YadaYuki/omochi/pkg/infrastructure/indexer"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/eng"

	_ "github.com/mattn/go-sqlite3"
)

func TestSearch(t *testing.T) {

	documentContents := []string{
		"java c js ruby cpp ts golang python", "c js ruby cpp ts golang python", "java c js ruby cpp ts golang python java",
	}
	documentCreates := []*entities.DocumentCreate{}
	for _, documentContent := range documentContents {
		documentCreates = append(documentCreates, entities.NewDocumentCreate(documentContent, strings.Split(documentContent, constant.WhiteSpace)))
	}

	testCases := []struct {
		query            string
		expectedContents []string
	}{
		{
			query:            "java",
			expectedContents: []string{"java c js ruby cpp ts golang python java", "java c js ruby cpp ts golang python"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.query, func(tt *testing.T) {
			client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
			defer client.Close()
			documentRepository := entdb.NewDocumentEntRepository(client)
			termRepository := entdb.NewTermEntRepository(client)
			tokenizer := eng.NewEnProseTokenizer()
			invertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
			indexer := indexer.NewIndexer(documentRepository, termRepository, tokenizer, invertIndexCompresser)
			for _, doc := range documentCreates {
				indexingErr := indexer.IndexingDocument(context.Background(), doc)
				if indexingErr != nil {
					t.Fatal(indexingErr)
				}
			}
			invertIndexCompressedCached := map[string]*entities.InvertIndex{}
			searcher := NewSearcher(invertIndexCompressedCached, termRepository, documentRepository, compresser.NewZlibInvertIndexCompresser(), tfidfranker.NewTfIdfDocumentRanker())

			searchResultDocs, searchErr := searcher.Search(context.Background(), &entities.Query{Keyword: tc.query})
			if searchErr != nil {
				t.Fatal(searchErr)
			}
			for i, expectedContent := range tc.expectedContents {
				if searchResultDocs[i].Content != expectedContent {
					t.Fatalf("searchResultDocs[i].Content should be %v, but got %v", expectedContent, searchResultDocs[i].Content)
				}
			}
		})
	}
}
