package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Document holds the schema definition for the Document entity.
type Document struct {
	ent.Schema
}

// Fields of the Document.
func (Document) Fields() []ent.Field {
	return []ent.Field{
		field.String("content"),
		field.String("tokenized_content"), // トークナイズしたコンテンツを" "(WHITE_SPACE)区切りで保存する
	}
}

// Mixin of the Document.
func (Document) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeStampMixin{},
	}
}

// Edges of the Document.
func (Document) Edges() []ent.Edge {
	return nil
}
