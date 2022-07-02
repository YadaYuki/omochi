package data

import (
	"encoding/csv"
	"io"
	"os"
)

func newTsvReader(reader io.Reader) *csv.Reader {
	r := csv.NewReader(reader)
	r.Comma = '\t'
	return r
}

func LoadDocumentsFromTsv(pathTo string) (*[]string, error) {
	reader, openErr := os.Open(pathTo)
	if openErr != nil {
		return nil, openErr
	}
	defer reader.Close()
	tsvReader := newTsvReader(reader)

	data, readErr := tsvReader.ReadAll()
	if readErr != nil {
		return nil, readErr
	}

	DocumentColIndex := 0
	documents := make([]string, len(data)-1)
	for i, row := range data[1:] {
		documents[i] = row[DocumentColIndex]
	}
	return &documents, nil
}
