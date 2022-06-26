package entities

type Posting struct {
	DocumentRelatedId   int64 `json:"document_related_id"`
	PositionsInDocument []int `json:"positions_in_document"`
}

func NewPosting(documentRelatedId int64, positionsInDocument []int) *Posting {
	return &Posting{DocumentRelatedId: documentRelatedId, PositionsInDocument: positionsInDocument}
}
