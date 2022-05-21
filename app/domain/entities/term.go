package entities

import (
	"time"

	"github.com/google/uuid"
)

type Term struct {
	Uuid      uuid.UUID `json:"uuid"`
	Word      string    `json:"word"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Terms = []Term
