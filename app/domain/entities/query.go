package entities

type SearchModeType string

const (
	And    SearchModeType = "And"
	Or     SearchModeType = "Or"
	Phrase SearchModeType = "Phrase"
)

type Query struct {
	// Keywords   *[]string `json:"keywords"` // TODO: Correspond to multi keywords
	Keyword    string `json:"keyword"`
	SearchMode string `json:"mode"`
}
