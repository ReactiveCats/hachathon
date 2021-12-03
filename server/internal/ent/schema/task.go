package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"server/internal/consts"
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
			Values(
				consts.ComplexityVeryLow,
				consts.ComplexityLow,
				consts.ComplexityMid,
				consts.ComplexityHigh,
				consts.ComplexityVeryHigh,
			).
			Default(consts.ComplexityMid),
		field.Enum("priority").
			Values(
				consts.PriorityVeryLow,
				consts.PriorityLow,
				consts.PriorityMid,
				consts.PriorityHigh,
				consts.PriorityVeryHigh,
			).
			Default(consts.PriorityMid),
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
