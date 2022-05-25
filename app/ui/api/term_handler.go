package api

import (
	"encoding/json"
	"net/http"

	"github.com/YadaYuki/omochi/app/errors/code"
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
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	term, err := h.u.FindTermById(r.Context(), uuid.MustParse(uuidStr))
	if err != nil {
		if err.Code == code.NotExist {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	termBody, _ := json.Marshal(term)
	w.Write(termBody)
}
