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
	TermId      uuid.UUID  `json:"term_id"`
	PostingList *[]Posting `json:"posting_list"`
}

type InvertIndexCompressed struct {
	Uuid                  uuid.UUID `json:"uuid"`
	TermId                uuid.UUID `json:"term_id"`
	PostingListCompressed []byte    `json:"posting_list_compressed"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type InvertIndexCompressedCreate struct {
	TermId                uuid.UUID `json:"term_id"`
	PostingListCompressed []byte    `json:"posting_list_compressed"`
}

func NewInvertIndexCreate(termId uuid.UUID, postingList *[]Posting) *InvertIndexCreate {
	return &InvertIndexCreate{TermId: termId, PostingList: postingList}
}

func NewInvertIndexCompressedCreate(termId uuid.UUID, postingListCompressed []byte) *InvertIndexCompressedCreate {
	return &InvertIndexCompressedCreate{TermId: termId, PostingListCompressed: postingListCompressed}
}
