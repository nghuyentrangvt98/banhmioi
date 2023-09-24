package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("username").Unique(),
		field.String("fullname"),
		field.Enum("gender").Values("male", "female"),
		field.String("dob").Nillable(),
		field.String("password"),
		field.String("email").Nillable(),
		field.String("phone"),
		field.Time("created_at"),
		field.Time("updated_at"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
