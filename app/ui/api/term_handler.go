package api

import (
	"encoding/json"
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
		return
	}
	term, err := h.u.FindTermById(r.Context(), uuid.MustParse(uuidStr))
	if err != nil { // TODO:
		return
	}

	termBody, err := json.Marshal(term)
	if err != nil { // TODO:
		return
	}
	w.Write(termBody)
}
