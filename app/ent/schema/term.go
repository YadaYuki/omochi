package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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

func (Term) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("word"),
	}
}

// Edges of the Term.
func (Term) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invert_index_compressed", InvertIndexCompressed.Type).
			Unique(),
	}
}
