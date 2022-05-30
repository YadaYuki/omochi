package entities

import (
	"time"

	"github.com/google/uuid"
)

type InvertIndex struct {
	Uuid        uuid.UUID
	TermId      uuid.UUID
	PostingList *[]Posting
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
