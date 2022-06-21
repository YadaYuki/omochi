package entities

import (
	"time"

	"github.com/google/uuid"
)

type Term struct {
	Uuid        uuid.UUID    `json:"uuid"`
	Word        string       `json:"word"`
	InvertIndex *InvertIndex `json:"invert_index"` // タームに対応した転置インデックス.
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type TermCompressed struct {
	Uuid                  uuid.UUID              `json:"uuid"`
	Word                  string                 `json:"word"`
	InvertIndexCompressed *InvertIndexCompressed `json:"invert_index_compressed"` // タームに対応した転置インデックス.
	CreatedAt             time.Time              `json:"created_at"`
	UpdatedAt             time.Time              `json:"updated_at"`
}

type TermCreate struct {
	Word              string       `json:"word"`
	InvertIndexCreate *InvertIndex `json:"invert_index"` // タームに対応した転置インデックス.
}

type TermCompressedCreate struct {
	Word                  string                 `json:"word"`
	InvertIndexCompressed *InvertIndexCompressed `json:"invert_index_compressed"` // タームに対応した転置インデックス.
}

func NewTermCreate(word string, invertIndex *InvertIndex) *TermCreate {
	return &TermCreate{Word: word}
}

func NewTermCompressedCreate(word string, invertIndex *InvertIndexCompressed) *TermCompressedCreate {
	return &TermCompressedCreate{Word: word, InvertIndexCompressed: invertIndex}
}
