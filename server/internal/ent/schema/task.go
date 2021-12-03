package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description").Optional(),
		field.String("priority").Default("5"),
		field.String("complexity").Default("5"),
		field.String("hard_deadline").Optional(),
		field.String("soft_deadline").Optional(),
		field.String("status").Default("123"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("creator", User.Type),
	}
}
