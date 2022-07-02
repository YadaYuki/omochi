package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/YadaYuki/omochi/cmd/seeds/data"
	"github.com/YadaYuki/omochi/pkg/common/slices"
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
	"golang.org/x/sync/errgroup"
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
	fmt.Println(entIndexer)
	// load documents
	docs, err := data.LoadDocumentsFromTsv(data.DoraemonDocumentTsvPath)
	if err != nil {
		log.Fatalf("failed loading documents: %v", err)
	}
	size := 200
	docLists := slices.SplitSlice(*docs, size)
	goroutines := len(docLists)
	ctx := context.Background()

	// index documents concurrently
	log.Println("start indexing documents")
	var eg errgroup.Group
	for i := 0; i < goroutines; i++ {
		docList := docLists[i]
		eg.Go(func() error {
			for _, doc := range docList {
				docCreate := entities.NewDocumentCreate(doc, []string{})
				if err := entIndexer.IndexingDocumentWithTx(ctx, docCreate); err != nil {
					return err
				}
				log.Println("indexed:", doc)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
}
