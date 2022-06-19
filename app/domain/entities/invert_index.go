package entities

import (
	"time"

	"github.com/google/uuid"
)

type InvertIndexDetail struct {
	Uuid        uuid.UUID  `json:"uuid"`
	TermId      uuid.UUID  `json:"term_id"`
	PostingList *[]Posting `json:"posting_list"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type InvertIndex struct {
	TermId      uuid.UUID  `json:"term_id"`
	PostingList *[]Posting `json:"posting_list"`
}

func NewInvertIndex(termId uuid.UUID, postingList *[]Posting) *InvertIndex {
	return &InvertIndex{TermId: termId, PostingList: postingList}
}

type InvertedIndexCompressedDetail struct {
	Uuid                  uuid.UUID `json:"uuid"`
	TermId                uuid.UUID `json:"term_id"`
	PostingListCompressed []byte    `json:"posting_list_compressed"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type InvertedIndexCompressed struct {
	TermId                uuid.UUID `json:"term_id"`
	PostingListCompressed []byte    `json:"posting_list_compressed"`
}

func NewInvertIndexCompressed(termId uuid.UUID, postingListCompressed []byte) *InvertedIndexCompressed {
	return &InvertedIndexCompressed{TermId: termId, PostingListCompressed: postingListCompressed}
}
