package entities

import (
	"time"
)

type DocumentDetail struct {
	Id               int64     `json:"id"`
	Content          string    `json:"content"`
	TokenizedContent []string  `json:"tokenized_content"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Document struct {
	Content          string   `json:"content"`
	TokenizedContent []string `json:"tokenized_content"`
}

type Documents = []Document
