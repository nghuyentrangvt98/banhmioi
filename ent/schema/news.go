package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.String("subtitle").Optional(),
		field.Text("content"),
		field.String("author"),
		field.Time("created_at"),
		field.String("category"),
		field.String("product_url").Optional(),
		field.String("image_url"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return nil
}
