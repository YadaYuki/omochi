package infrastructure

import "github.com/YadaYuki/omochi/app/ent"

type TermRepository struct {
	db *ent.Client
}
