package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Term holds the schema definition for the Term entity.
type Term struct {
	ent.Schema
}

// Fields of the Term.
func (Term) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Term.
func (Term) Edges() []ent.Edge {
	return nil
}
