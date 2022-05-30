package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Term holds the schema definition for the Term entity.
type Term struct {
	ent.Schema
}

// Fields of the Term.
func (Term) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("uuid").Default(uuid.New),
		field.String("word"),
	}
}

// Mixin of the Term.
func (Term) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStampMixin{},
	}
}

// Edges of the Term.
func (Term) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("invert_index", Term.Type).
			Ref("term").
			Unique(),
	}
}
