package entities

import "github.com/google/uuid"

type Posting struct {
	DocumentRelatedId   uuid.UUID `json:"document_related_id"`
	PositionsInDocument []int     `json:"positions_in_document"`
}
