package document

import (
	"encoding/json"
	"net/http"

	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/errors"
	"github.com/YadaYuki/omochi/pkg/errors/code"
	"github.com/YadaYuki/omochi/pkg/usecase/search"
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
	keyword := r.URL.Query().Get("keyword")
	mode := r.URL.Query().Get("mode")
	requestBody := RequestSearchDocument{
		Keyword: keyword,
		Mode:    mode,
	}
	// parseErr := json.NewDecoder(r.Body).Decode(&requestBody)
	// if parseErr != nil {
	// 	http.Error(w, parseErr.Error(), http.StatusUnprocessableEntity)
	// 	return
	// }
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
		http.Error(w, err.Error(), http.StatusNotFound)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
