package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/env"
	"github.com/YadaYuki/omochi/app/infrastructure/persistence/entdb"
	"github.com/YadaYuki/omochi/app/ui/api"
	termUsecase "github.com/YadaYuki/omochi/app/usecase/term"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jdkato/prose/v2"
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

	// terms := []string{"hello", "world", "omochi"}
	// for _, term := range terms {
	// 	_, err := CreateTerm(term, context.Background(), db)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	doc, err := prose.NewDocument("Go is an open-source programming language created at Google.")
	if err != nil {
		log.Fatal(err)
	}

	for _, tok := range doc.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}

	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
	}

	for _, sent := range doc.Sentences() {
		fmt.Println(sent.Text)
	}
	termRepository := entdb.NewTermEntRepository(db)
	useCase := termUsecase.NewTermUseCase(termRepository)
	termHandler := api.NewTermHandler(useCase)

	log.Println("Successfully connected to MySQL")
	log.Println("application started")
	r := mux.NewRouter()
	r.HandleFunc("/term/{uuid}", termHandler.FindTermByIdHandler)
	http.ListenAndServe(":8081", r)
}

func CreateTerm(word string, ctx context.Context, client *ent.Client) (*ent.Term, error) {
	u, err := client.Term.
		Create().
		SetWord(word).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating term: %w", err)
	}
	log.Println("term was created: ", u)
	return u, nil
}
