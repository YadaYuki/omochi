package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DB_USER := os.Getenv("MYSQL_USER")
	DB_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	DB_HOST := os.Getenv("MYSQL_HOST")
	DB_NAME := os.Getenv("MYSQL_DATABASE")
	DB_PORT := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
