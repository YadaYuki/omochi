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

func NewDocument(content string, tokenizedConetnt []string) *Document {
	return &Document{Content: content, TokenizedContent: tokenizedConetnt}
}

type Documents = []Document
