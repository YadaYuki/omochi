package api

import (
	"github.com/YadaYuki/omochi/app/interface/api/document"
	"github.com/YadaYuki/omochi/app/interface/api/term"
	susecase "github.com/YadaYuki/omochi/app/usecase/search"
	tusecase "github.com/YadaYuki/omochi/app/usecase/term"
	"github.com/go-chi/chi/v5"
)

func InitRoutes(r chi.Router, termUsecase tusecase.TermUseCase, searchUsecase susecase.SearchUseCase) {
	// teerm
	termController := term.NewTermController(termUsecase)
	r.Route("/term", func(r chi.Router) {
		r.Get("/{uuid}", termController.FindTermCompressedById)
	})
	// document
	documentController := document.NewDocumentController(searchUsecase)
	r.Route("/document", func(r chi.Router) {
		r.Get("/search", documentController.SearchDocuments)
	})
}
