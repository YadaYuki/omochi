package term

import (
	"encoding/json"
	"net/http"

	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
	usecase "github.com/YadaYuki/omochi/pkg/usecase/term"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TermController struct {
	u usecase.TermUseCase
}

func NewTermController(u usecase.TermUseCase) *TermController {
	return &TermController{u: u}
}

func (controller *TermController) FindTermCompressedById(w http.ResponseWriter, r *http.Request) {
	uuidStr := chi.URLParam(r, "uuid")

	id, errId := uuid.Parse(uuidStr)
	if errId != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	term, err := controller.u.FindTermCompressedById(r.Context(), id)
	if err != nil {
		covertErrorToResponse(err, w)
		return
	}
	termBody, _ := json.Marshal(term)
	w.Write(termBody)
}

func covertErrorToResponse(err *errors.Error, w http.ResponseWriter) {
	switch err.Code {
	case code.NotExist:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
