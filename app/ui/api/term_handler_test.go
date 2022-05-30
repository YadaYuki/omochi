package api

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
	"github.com/YadaYuki/omochi/app/infrastructure/entdb"
	"github.com/YadaYuki/omochi/app/usecase"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func TestTermHandler_FindTermByIdHandler(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termHandler := createTermHandler(t, client)
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
			Save(context.Background())
		req, err := http.NewRequest("GET", fmt.Sprintf("/term/%s", termCreated.ID.String()), nil)
		if err != nil {
			t.Fatal(err)
		}
		res := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/term/{uuid}", termHandler.FindTermByIdHandler)
		router.ServeHTTP(res, req)
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

func TestTermHandler_FindTermByIdHandlerError(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termHandler := createTermHandler(t, client)

	req, err := http.NewRequest("GET", fmt.Sprintf("/term/%s", ""), nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/term/{uuid}", termHandler.FindTermByIdHandler)
	router.ServeHTTP(res, req)
	if res.Code != http.StatusBadRequest {
		t.Fatalf("expected %d, but got %d", http.StatusBadRequest, res.Code)
	}
}

func createTermHandler(t testing.TB, client *ent.Client) *TermHandler {
	termRepository := entdb.NewTermEntRepository(client)
	useCase := usecase.NewTermUseCase(termRepository)
	termHandler := NewTermHandler(useCase)
	return termHandler
}
