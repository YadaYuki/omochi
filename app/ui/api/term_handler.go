package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YadaYuki/omochi/app/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TermHandler struct {
	u usecase.ITermUseCase
}

func NewTermHandler(u usecase.ITermUseCase) *TermHandler {
	return &TermHandler{u: u}
}

func (h *TermHandler) FindTermByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidStr, ok := vars["uuid"]
	if !ok { // TODO:
		fmt.Println("uuid is not found")
	}
	term, err := h.u.FindTermById(r.Context(), uuid.MustParse(uuidStr))
	if err != nil { // TODO:
		panic(err)
	}
	termBody, err := json.Marshal(term)
	if err != nil { // TODO:
		panic(err)
		// return
	}
	w.Write(termBody)
}
