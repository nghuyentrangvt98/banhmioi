package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Text("description"),
		field.Int64("list_price"),
		field.Int64("sale_price"),
		field.JSON("variation", make(map[string]interface{})),
		field.Int32("stock_qty"),
		field.String("product_url"),
		field.String("category"),
		field.String("measurement_unit"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
