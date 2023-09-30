package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Order struct {
	ent.Schema
}

// Fields of the Product.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("phone"),
		field.String("address"),
		field.String("note").Optional(),
		field.Float32("total"),
		field.Float32("discount"),
	}
}

// Edges of the Product.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cart", Cart.Type),
		edge.To("user", User.Type).Unique(),
	}
}
