package main

import (
	"context"
	"log"

	"github.com/YadaYuki/omochi/pkg/config"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/ent/migrate"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	client, err := ent.Open("mysql", config.MysqlConnection)
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
	log.Println("Successfully migrated ! ")
}
