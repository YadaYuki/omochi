package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// InvertIndexCompressed holds the schema definition for the InvertIndexCompressed entity.
type InvertIndexCompressed struct {
	ent.Schema
}

// Fields of the InvertIndexCompressed.
func (InvertIndexCompressed) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("uuid").Default(uuid.New),
		field.Bytes("posting_list_compressed").MaxLen(1 << 30),
	}
}

// Mixin of the Term.
func (InvertIndexCompressed) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStampMixin{},
	}
}

// Edges of the InvertIndexCompressed.
func (InvertIndexCompressed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("term_related", Term.Type).Ref("invert_index_compressed").Unique().Required(),
	}
}
