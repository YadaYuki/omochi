package main

import (
	"context"
	"log"
	"net/http"

	"github.com/YadaYuki/omochi/pkg/config"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/documentranker/tfidfranker"
	"github.com/YadaYuki/omochi/pkg/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/pkg/infrastructure/searcher"
	api "github.com/YadaYuki/omochi/pkg/interface/api"
	susecase "github.com/YadaYuki/omochi/pkg/usecase/search"
	tusecase "github.com/YadaYuki/omochi/pkg/usecase/term"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := ent.Open("mysql", config.MysqlConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Successfully connected to MySQL")

	// initialize term usecase
	termRepository := entdb.NewTermEntRepository(db)
	termUseCase := tusecase.NewTermUseCase(termRepository)

	// initialize search usecase
	documentRepository := entdb.NewDocumentEntRepository(db)
	invertIndexCached := map[string]*entities.InvertIndex{} // TODO: initialize by frequent words
	zlibInvertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
	tfIdfDocumentRanker := tfidfranker.NewTfIdfDocumentRanker()
	searcher := searcher.NewSearcher(invertIndexCached, termRepository, documentRepository, zlibInvertIndexCompresser, tfIdfDocumentRanker)
	searchUseCase := susecase.NewSearchUseCase(searcher)

	// init & start api
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		api.InitRoutes(r, termUseCase, searchUseCase)
	})
	log.Println("application started")
	http.ListenAndServe(":8081", r)
}
