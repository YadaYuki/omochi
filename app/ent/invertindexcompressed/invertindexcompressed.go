// Code generated by entc, DO NOT EDIT.

package invertindexcompressed

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the invertindexcompressed type in the database.
	Label = "invert_index_compressed"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uuid"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldPostingListCompressed holds the string denoting the posting_list_compressed field in the database.
	FieldPostingListCompressed = "posting_list_compressed"
	// EdgeTerm holds the string denoting the term edge name in mutations.
	EdgeTerm = "term"
	// Table holds the table name of the invertindexcompressed in the database.
	Table = "invert_index_compresseds"
	// TermTable is the table that holds the term relation/edge.
	TermTable = "terms"
	// TermInverseTable is the table name for the Term entity.
	// It exists in this package in order to avoid circular dependency with the "term" package.
	TermInverseTable = "terms"
	// TermColumn is the table column denoting the term relation/edge.
	TermColumn = "invert_index_compressed_term"
)

// Columns holds all SQL columns for invertindexcompressed fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldPostingListCompressed,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// PostingListCompressedValidator is a validator for the "posting_list_compressed" field. It is called by the builders before save.
	PostingListCompressedValidator func([]byte) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)