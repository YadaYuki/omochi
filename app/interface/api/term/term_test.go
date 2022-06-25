package term

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/enttest"
	"github.com/YadaYuki/omochi/app/infrastructure/persistence/entdb"
	usecase "github.com/YadaYuki/omochi/app/usecase/term"
	"github.com/go-chi/chi/v5"

	_ "github.com/mattn/go-sqlite3"
)

func TestTermController_FindTermById(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termController := createTermController(t, client)
	testCases := []struct {
		word string
	}{
		{"sample"},
	}
	for _, tc := range testCases {
		// create mock data
		termCreated, _ := client.Term.
			Create().
			SetWord(tc.word).
			SetPostingListCompressed([]byte("sample")).
			Save(context.Background())
		req, err := http.NewRequest("GET", fmt.Sprintf("/term/%s", termCreated.ID.String()), nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		r := chi.NewRouter()
		r.Get("/term/{uuid}", termController.FindTermCompressedById)
		r.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			t.Fatalf("expected %d, but got %d", http.StatusOK, res.Code)
		}
		var term entities.Term
		if err := json.Unmarshal(res.Body.Bytes(), &term); err != nil {
			t.Fatal(err)
		}
		if term.Word != tc.word {
			t.Fatalf("expected %s, but got %s", tc.word, term.Word)
		}
	}
}

func createTermController(t testing.TB, client *ent.Client) *TermController {
	termRepository := entdb.NewTermEntRepository(client)
	useCase := usecase.NewTermUseCase(termRepository)
	termController := NewTermController(useCase)
	return termController
}
