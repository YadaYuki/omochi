// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/YadaYuki/omochi/pkg/ent/document"
)

// Document is the model entity for the Document schema.
type Document struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// TokenizedContent holds the value of the "tokenized_content" field.
	TokenizedContent string `json:"tokenized_content,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Document) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			values[i] = new(sql.NullInt64)
		case document.FieldContent, document.FieldTokenizedContent:
			values[i] = new(sql.NullString)
		case document.FieldCreatedAt, document.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Document", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Document fields.
func (d *Document) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case document.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case document.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case document.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case document.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				d.Content = value.String
			}
		case document.FieldTokenizedContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tokenized_content", values[i])
			} else if value.Valid {
				d.TokenizedContent = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Document.
// Note that you need to call Document.Unwrap() before calling this method if this Document
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Document) Update() *DocumentUpdateOne {
	return (&DocumentClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Document entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Document) Unwrap() *Document {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Document is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Document) String() string {
	var builder strings.Builder
	builder.WriteString("Document(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", content=")
	builder.WriteString(d.Content)
	builder.WriteString(", tokenized_content=")
	builder.WriteString(d.TokenizedContent)
	builder.WriteByte(')')
	return builder.String()
}

// Documents is a parsable slice of Document.
type Documents []*Document

func (d Documents) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
