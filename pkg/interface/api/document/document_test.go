package document

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/YadaYuki/omochi/pkg/common/constant"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/ent/enttest"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/documentranker/tfidfranker"
	"github.com/YadaYuki/omochi/pkg/infrastructure/indexer"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/searcher"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/eng"
	"github.com/go-chi/chi/v5"

	susecase "github.com/YadaYuki/omochi/pkg/usecase/search"

	_ "github.com/mattn/go-sqlite3"
)

func TestTermController_FindTermById(t *testing.T) {

	documentContents := []string{
		"java c js ruby cpp ts golang python", "c js ruby cpp ts golang python", "JAVA C JS RUBY CPP TS GOLANG PYTHON JAVA",
	}
	documentCreates := []*entities.DocumentCreate{}
	for _, documentContent := range documentContents {
		documentCreates = append(documentCreates, entities.NewDocumentCreate(documentContent, strings.Split(documentContent, constant.WhiteSpace)))
	}

	testCases := []struct {
		query            string
		mode             entities.SearchModeType
		expectedContents []string
	}{
		{
			query:            "java",
			mode:             entities.Or,
			expectedContents: []string{"JAVA C JS RUBY CPP TS GOLANG PYTHON JAVA", "java c js ruby cpp ts golang python"},
		},
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	documentController := createDocumentController(t, client)
	indexer := createIndexer(t, client)
	for _, doc := range documentCreates {
		indexingErr := indexer.IndexingDocument(context.Background(), doc)
		if indexingErr != nil {
			t.Fatal(indexingErr)
		}
	}
	DummyPath := "/search_test"
	for _, tc := range testCases {

		reqBody, err := json.Marshal(&RequestSearchDocument{Keyword: tc.query, Mode: string(tc.mode)})
		if err != nil {
			t.Fatal(err)
		}
		reqBodyReader := bytes.NewReader(reqBody)
		req, _ := http.NewRequest("GET", DummyPath, reqBodyReader)

		res := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get(DummyPath, documentController.SearchDocuments)
		r.ServeHTTP(res, req)

		if res.Code != http.StatusOK {
			t.Fatalf("expected %d, but got %d", http.StatusOK, res.Code)
		}
		var respBody ReseponseSearchDocument
		if err := json.Unmarshal(res.Body.Bytes(), &respBody); err != nil {
			t.Fatal(err)
		}
		if len(respBody.Documents) != len(tc.expectedContents) {
			t.Fatalf("expected %d, but got %d", len(tc.expectedContents), len(respBody.Documents))
		}
		fmt.Println(res.Body.String())
		for i, doc := range respBody.Documents {
			if doc.Content != tc.expectedContents[i] {
				t.Fatalf("expected %s, but got %s", tc.expectedContents[i], doc.Content)
			}
		}
	}
}

func createDocumentController(t testing.TB, client *ent.Client) *DocumentController {
	documentRepository := entdb.NewDocumentEntRepository(client)
	invertIndexCached := map[string]*entities.InvertIndex{} // TODO: initialize by frequent words
	zlibInvertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
	tfIdfDocumentRanker := tfidfranker.NewTfIdfDocumentRanker()
	termRepository := entdb.NewTermEntRepository(client)
	searcher := searcher.NewSearcher(invertIndexCached, termRepository, documentRepository, zlibInvertIndexCompresser, tfIdfDocumentRanker)
	searchUseCase := susecase.NewSearchUseCase(searcher)
	documentController := NewDocumentController(searchUseCase)
	return documentController
}

func createIndexer(t testing.TB, client *ent.Client) service.Indexer {
	documentRepository := entdb.NewDocumentEntRepository(client)
	termRepository := entdb.NewTermEntRepository(client)
	tokenizer := eng.NewEnProseTokenizer()
	invertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
	indexer := indexer.NewIndexer(documentRepository, termRepository, tokenizer, invertIndexCompresser)
	return indexer
}
