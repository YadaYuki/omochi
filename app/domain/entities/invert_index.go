package entities

import (
	"time"

	"github.com/google/uuid"
)

type InvertIndex struct {
	Uuid        uuid.UUID  `json:"uuid"`
	TermId      uuid.UUID  `json:"term_id"`
	PostingList *[]Posting `json:"posting_list"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type InvertIndexCreate struct {
	TermId      uuid.UUID  `json:"term_id"` // REVIEW: TermCreate may be better ?
	PostingList *[]Posting `json:"posting_list"`
}

type InvertedIndexCompressed struct {
	Uuid                  uuid.UUID `json:"uuid"`
	TermId                uuid.UUID `json:"term_id"`
	PostingListCompressed []byte    `json:"posting_list_compressed"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
