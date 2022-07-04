package entities

type SearchModeType string

const (
	And SearchModeType = "And"
	Or  SearchModeType = "Or"
)

type Query struct {
	Keywords   *[]string      `json:"keywords"`
	SearchMode SearchModeType `json:"mode"`
}

func NewQuery(keyword []string, searchMode SearchModeType) *Query {
	return &Query{Keywords: &keyword, SearchMode: searchMode}
}
