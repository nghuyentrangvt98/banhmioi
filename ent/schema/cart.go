package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Cart struct {
	ent.Schema
}

// Fields of the Product.
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		// field.Int("product_id"),
		// field.Int("user_id"),
		field.JSON("variation", make(map[string]interface{})),
		field.Int32("qty"),
		field.Int64("price"),
		field.Enum("status").Values("done", "processing", "fail"),
		field.String("note").Optional(),
	}
}

// Edges of the Product.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique(),
		edge.To("product", Product.Type).
			Unique(),
	}
}
