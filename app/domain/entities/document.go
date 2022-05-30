package entities

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	Uuid      uuid.UUID `json:"uuid"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DocumentCreate struct {
	Content string `json:"content"`
}

type Documents = []Document
