package entities

import (
	"time"
)

type Document struct {
	Id               int64     `json:"id"`
	Content          string    `json:"content"`
	TokenizedContent []string  `json:"tokenized_content"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type DocumentCreate struct {
	Content          string   `json:"content"`
	TokenizedContent []string `json:"tokenized_content"`
}

func NewDocumentCreate(content string, tokenizedConetnt []string) *DocumentCreate {
	return &DocumentCreate{Content: content, TokenizedContent: tokenizedConetnt}
}
