// Code generated by entc, DO NOT EDIT.

package task

import (
	"server/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// Complexity applies equality check predicate on the "complexity" field. It's identical to ComplexityEQ.
func Complexity(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComplexity), v))
	})
}

// HardDeadline applies equality check predicate on the "hard_deadline" field. It's identical to HardDeadlineEQ.
func HardDeadline(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHardDeadline), v))
	})
}

// SoftDeadline applies equality check predicate on the "soft_deadline" field. It's identical to SoftDeadlineEQ.
func SoftDeadline(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSoftDeadline), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriority), v))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPriority), v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPriority), v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriority), v))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriority), v))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriority), v))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriority), v))
	})
}

// PriorityContains applies the Contains predicate on the "priority" field.
func PriorityContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPriority), v))
	})
}

// PriorityHasPrefix applies the HasPrefix predicate on the "priority" field.
func PriorityHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPriority), v))
	})
}

// PriorityHasSuffix applies the HasSuffix predicate on the "priority" field.
func PriorityHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPriority), v))
	})
}

// PriorityEqualFold applies the EqualFold predicate on the "priority" field.
func PriorityEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPriority), v))
	})
}

// PriorityContainsFold applies the ContainsFold predicate on the "priority" field.
func PriorityContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPriority), v))
	})
}

// ComplexityEQ applies the EQ predicate on the "complexity" field.
func ComplexityEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldComplexity), v))
	})
}

// ComplexityNEQ applies the NEQ predicate on the "complexity" field.
func ComplexityNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldComplexity), v))
	})
}

// ComplexityIn applies the In predicate on the "complexity" field.
func ComplexityIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldComplexity), v...))
	})
}

// ComplexityNotIn applies the NotIn predicate on the "complexity" field.
func ComplexityNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldComplexity), v...))
	})
}

// ComplexityGT applies the GT predicate on the "complexity" field.
func ComplexityGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldComplexity), v))
	})
}

// ComplexityGTE applies the GTE predicate on the "complexity" field.
func ComplexityGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldComplexity), v))
	})
}

// ComplexityLT applies the LT predicate on the "complexity" field.
func ComplexityLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldComplexity), v))
	})
}

// ComplexityLTE applies the LTE predicate on the "complexity" field.
func ComplexityLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldComplexity), v))
	})
}

// ComplexityContains applies the Contains predicate on the "complexity" field.
func ComplexityContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldComplexity), v))
	})
}

// ComplexityHasPrefix applies the HasPrefix predicate on the "complexity" field.
func ComplexityHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldComplexity), v))
	})
}

// ComplexityHasSuffix applies the HasSuffix predicate on the "complexity" field.
func ComplexityHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldComplexity), v))
	})
}

// ComplexityEqualFold applies the EqualFold predicate on the "complexity" field.
func ComplexityEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldComplexity), v))
	})
}

// ComplexityContainsFold applies the ContainsFold predicate on the "complexity" field.
func ComplexityContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldComplexity), v))
	})
}

// HardDeadlineEQ applies the EQ predicate on the "hard_deadline" field.
func HardDeadlineEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineNEQ applies the NEQ predicate on the "hard_deadline" field.
func HardDeadlineNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineIn applies the In predicate on the "hard_deadline" field.
func HardDeadlineIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHardDeadline), v...))
	})
}

// HardDeadlineNotIn applies the NotIn predicate on the "hard_deadline" field.
func HardDeadlineNotIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHardDeadline), v...))
	})
}

// HardDeadlineGT applies the GT predicate on the "hard_deadline" field.
func HardDeadlineGT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineGTE applies the GTE predicate on the "hard_deadline" field.
func HardDeadlineGTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineLT applies the LT predicate on the "hard_deadline" field.
func HardDeadlineLT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineLTE applies the LTE predicate on the "hard_deadline" field.
func HardDeadlineLTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHardDeadline), v))
	})
}

// HardDeadlineIsNil applies the IsNil predicate on the "hard_deadline" field.
func HardDeadlineIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHardDeadline)))
	})
}

// HardDeadlineNotNil applies the NotNil predicate on the "hard_deadline" field.
func HardDeadlineNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHardDeadline)))
	})
}

// SoftDeadlineEQ applies the EQ predicate on the "soft_deadline" field.
func SoftDeadlineEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineNEQ applies the NEQ predicate on the "soft_deadline" field.
func SoftDeadlineNEQ(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineIn applies the In predicate on the "soft_deadline" field.
func SoftDeadlineIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSoftDeadline), v...))
	})
}

// SoftDeadlineNotIn applies the NotIn predicate on the "soft_deadline" field.
func SoftDeadlineNotIn(vs ...time.Time) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSoftDeadline), v...))
	})
}

// SoftDeadlineGT applies the GT predicate on the "soft_deadline" field.
func SoftDeadlineGT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineGTE applies the GTE predicate on the "soft_deadline" field.
func SoftDeadlineGTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineLT applies the LT predicate on the "soft_deadline" field.
func SoftDeadlineLT(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineLTE applies the LTE predicate on the "soft_deadline" field.
func SoftDeadlineLTE(v time.Time) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSoftDeadline), v))
	})
}

// SoftDeadlineIsNil applies the IsNil predicate on the "soft_deadline" field.
func SoftDeadlineIsNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSoftDeadline)))
	})
}

// SoftDeadlineNotNil applies the NotNil predicate on the "soft_deadline" field.
func SoftDeadlineNotNil() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSoftDeadline)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Task {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Task(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CreatorTable, CreatorPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CreatorTable, CreatorPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
