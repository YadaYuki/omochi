package main

import (
	"context"
	"flag"
	"log"

	"github.com/YadaYuki/omochi/cmd/seeds/data"
	"github.com/YadaYuki/omochi/pkg/config"
	"github.com/YadaYuki/omochi/pkg/domain/entities"
	"github.com/YadaYuki/omochi/pkg/domain/service"
	"github.com/YadaYuki/omochi/pkg/ent"
	"github.com/YadaYuki/omochi/pkg/infrastructure/compresser"
	"github.com/YadaYuki/omochi/pkg/infrastructure/indexer/entindexer"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/eng"
	"github.com/YadaYuki/omochi/pkg/infrastructure/tokenizer/ja"
	"github.com/YadaYuki/omochi/pkg/infrastructure/transaction/wrapper"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := ent.Open("mysql", config.MysqlConnection)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer db.Close()

	// read arguments
	language := *(flag.String("lang", "ja", "language of the documents"))
	flag.Parse()

	// initialize term usecase
	t := wrapper.NewEntTransactionWrapper()
	zlibInvertIndexCompresser := compresser.NewZlibInvertIndexCompresser()
	// create tokenizer
	var tokenizer service.Tokenizer
	if language == "ja" {
		tokenizer = ja.NewJaKagomeTokenizer()
	} else if language == "eng" {
		tokenizer = eng.NewEnProseTokenizer()
	} else {
		log.Fatalf("unknown language: %s", language)
	}
	entIndexer := entindexer.NewEntIndexer(db, t, tokenizer, zlibInvertIndexCompresser)

	// load documents

	docs, loadErr := data.LoadDocumentsFromTsv(data.DoraemonDocumentTsvPath)
	if loadErr != nil {
		panic(loadErr)
	}

	ctx := context.Background()

	for _, doc := range *docs {
		EmptyTokenizedContent := []string{}
		documentCreate := entities.NewDocumentCreate(doc, EmptyTokenizedContent)
		indexingErr := entIndexer.IndexingDocumentWithTx(ctx, documentCreate)
		if indexingErr != nil {
			panic(indexingErr)
		}
	}

}
