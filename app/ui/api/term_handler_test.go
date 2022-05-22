package api

import (
	"net/http/httptest"
	"testing"

	"github.com/YadaYuki/omochi/app/ent/enttest"
	"github.com/YadaYuki/omochi/app/infrastructure"
	"github.com/YadaYuki/omochi/app/usecase"
)

func TestTermHandler_FindTermByIdHandler(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	termRepository := infrastructure.NewTermRepository(client)
	useCase := usecase.NewTermUseCase(termRepository)
	termHandler := NewTermHandler(useCase)

	testServer := httptest.NewServer(termHandler.FindTermByIdHandler) // サーバを立てる
	t.Cleanup(func() {
		testServer.Close()
	})

}
