package entindexer

import (
	"github.com/YadaYuki/omochi/pkg/ent"
)

type EntIndexer struct {
	db *ent.Client
}

func NewEntIndexer(db *ent.Client) *EntIndexer {
	return &EntIndexer{db: db}
}
