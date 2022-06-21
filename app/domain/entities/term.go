package entities

import (
	"time"

	"github.com/google/uuid"
)

type TermDetail struct {
	Uuid        uuid.UUID          `json:"uuid"`
	Word        string             `json:"word"`
	InvertIndex *InvertIndexDetail `json:"invert_index"` // タームに対応した転置インデックス.
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type TermCompressedDetail struct {
	Uuid        uuid.UUID                `json:"uuid"`
	Word        string                   `json:"word"`
	InvertIndex *InvertedIndexCompressed `json:"invert_index"` // タームに対応した転置インデックス.
	CreatedAt   time.Time                `json:"created_at"`
	UpdatedAt   time.Time                `json:"updated_at"`
}

type Term struct {
	Word string `json:"word"`
}

func NewTerm(word string) *Term {
	return &Term{Word: word}
}

type Terms = []Term
