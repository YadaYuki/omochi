package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ResponseBody struct {
	Message string `json:"message"`
}

func hello(w http.ResponseWriter, req *http.Request) {
	r := ResponseBody{Message: "Hello World"}
	resBody, _ := json.Marshal(r)
	w.Write(resBody)
}

func main() {
	DB_USER := os.Getenv("MYSQL_USER")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_HOST := os.Getenv("MYSQL_HOST")
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	DB_PORT := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to MySQL")
	log.Println("application started")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8082", nil)
}
