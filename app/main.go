package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/infrastructure/entdb"
	"github.com/YadaYuki/omochi/app/ui/api"
	"github.com/YadaYuki/omochi/app/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	DB_USER := os.Getenv("MYSQL_USER")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_HOST := os.Getenv("MYSQL_HOST")
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	DB_PORT := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
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
	termRepository := entdb.NewTermEntRepository(db)
	useCase := usecase.NewTermUseCase(termRepository)
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
