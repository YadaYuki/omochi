package entities

type InvertIndex struct {
	PostingList *[]Posting `json:"posting_list"`
}

type InvertIndexCompressed struct {
	PostingListCompressed []byte `json:"posting_list_compressed"`
}

func NewInvertIndex(postingList *[]Posting) *InvertIndex {
	return &InvertIndex{PostingList: postingList}
}

func NewInvertIndexCompressed(postingListCompressed []byte) *InvertIndexCompressed {
	return &InvertIndexCompressed{PostingListCompressed: postingListCompressed}
}
