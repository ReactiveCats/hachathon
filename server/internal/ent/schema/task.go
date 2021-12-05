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
		field.Int("icon").Min(0).Max(16).Default(0),
		field.String("title").MaxLen(64),
		field.String("description").MaxLen(256).Optional(),
		field.Time("deadline").Optional(),
		field.Int("estimated").Min(0).Max(86400).Optional(),
		field.Int8("importance").Min(0).Max(11).Default(5),
		field.Int8("urgency").Min(0).Max(11).Default(5),

		field.Float("custom_mult").Default(1),

		field.Float("lo"),
		field.Float("hi"),

		field.Int("tag_id").Optional(),
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
		edge.From("tagg", Tag.Type).
			Ref("tasks").
			Field("tag_id").
			Unique(),
	}
}
