package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Int("icon").Default(0),
		field.String("title"),
		field.String("description").Optional(),
		field.Time("deadline").Optional(),
		field.Int("estimated").Optional(),
		field.Enum("complexity").
			Values("low", "mid", "high").
			Default("mid"),
		field.Enum("priority").
			Values("low", "mid", "high").
			Default("mid"),
		field.Int("creator_id"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("tasks").
			Field("creator_id").Required().
			Unique(),
	}
}
