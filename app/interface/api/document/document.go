package document

import (
	"encoding/json"
	"net/http"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
	"github.com/YadaYuki/omochi/app/usecase/search"
)

type DocumentController struct {
	searchUsecase search.SearchUseCase
}

func NewDocumentController(searchUsecase search.SearchUseCase) *DocumentController {
	return &DocumentController{searchUsecase}
}

type RequestSearchDocument struct {
	Keyword string `json:"keyword"`
	Mode    string `json:"mode"`
}

type ReseponseSearchDocument struct {
	Documents []entities.Document `json:"documents"`
}

func (controller *DocumentController) SearchDocuments(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestSearchDocument
	parseErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if parseErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	query := entities.NewQuery(requestBody.Keyword, entities.SearchModeType(requestBody.Mode))
	documents, searchErr := controller.searchUsecase.SearchDocuments(r.Context(), query)
	if searchErr != nil {
		covertErrorToResponse(searchErr, w)
		return
	}
	responseBody := &ReseponseSearchDocument{}
	for _, doc := range documents {
		responseBody.Documents = append(responseBody.Documents, *doc)
	}
	documentBody, jsonErr := json.Marshal(responseBody)
	if jsonErr != nil {
		covertErrorToResponse(errors.NewError(code.Unknown, jsonErr), w)
		return
	}

	w.Write(documentBody)
}

func covertErrorToResponse(err *errors.Error, w http.ResponseWriter) {
	switch err.Code {
	case code.NotExist:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
