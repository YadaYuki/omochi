package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/YadaYuki/omochi/app/domain/entities"
	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/env"
	"github.com/YadaYuki/omochi/app/infrastructure/compresser"
	"github.com/YadaYuki/omochi/app/infrastructure/documentranker/tfidfranker"
	"github.com/YadaYuki/omochi/app/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/app/infrastructure/searcher"
	api "github.com/YadaYuki/omochi/app/interface/api"
	susecase "github.com/YadaYuki/omochi/app/usecase/search"
	tusecase "github.com/YadaYuki/omochi/app/usecase/term"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	db, err := ent.Open("mysql", connectionString)
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
