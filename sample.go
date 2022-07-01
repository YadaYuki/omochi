package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func NewTsvReader(reader io.Reader) *csv.Reader {
	r := csv.NewReader(reader)
	r.Comma = '\t'
	return r
}

func main() {
	reader, openErr := os.Open("./pkg/data/ja/doraemon.tsv")
	if openErr != nil {
		panic(openErr)
	}
	defer reader.Close()
	tsvReader := NewTsvReader(reader)

	data, readErr := tsvReader.ReadAll()
	if readErr != nil {
		panic(readErr)
	}
	fmt.Println(data)
	fmt.Println(len(data) - 1)
}
