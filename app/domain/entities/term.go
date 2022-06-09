package entities

import (
	"time"

	"github.com/google/uuid"
)

type TermDetail struct {
	Uuid      uuid.UUID `json:"uuid"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Term struct {
	Word string `json:"word"`
}

func NewTerm(word string) *Term {
	return &Term{Word: word}
}

type Terms = []Term
