package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/env"
	"github.com/YadaYuki/omochi/app/infrastructure/persistence/entdb"
	handler "github.com/YadaYuki/omochi/app/interface/handler"
	usecase "github.com/YadaYuki/omochi/app/usecase/term"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	db, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	termRepository := entdb.NewTermEntRepository(db)
	useCase := usecase.NewTermUseCase(termRepository)
	termHandler := handler.NewTermHandler(useCase)

	log.Println("Successfully connected to MySQL")
	log.Println("application started")
	r := mux.NewRouter()
	r.HandleFunc("/term/{uuid}", termHandler.FindTermByIdHandler)
	http.ListenAndServe(":8081", r)
}
