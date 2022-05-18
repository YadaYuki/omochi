package entities

import "time"

type Term struct {
	Uuid      string
	Word      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Terms = []Term
