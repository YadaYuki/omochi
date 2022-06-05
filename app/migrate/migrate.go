package main

import (
	"context"
	"fmt"
	"log"

	"github.com/YadaYuki/omochi/app/ent"
	"github.com/YadaYuki/omochi/app/ent/migrate"
	"github.com/YadaYuki/omochi/app/env"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
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
