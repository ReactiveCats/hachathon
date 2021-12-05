// Code generated by entc, DO NOT EDIT.

package task

import (
	"time"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldEstimated holds the string denoting the estimated field in the database.
	FieldEstimated = "estimated"
	// FieldImportance holds the string denoting the importance field in the database.
	FieldImportance = "importance"
	// FieldUrgency holds the string denoting the urgency field in the database.
	FieldUrgency = "urgency"
	// FieldCustomMult holds the string denoting the custom_mult field in the database.
	FieldCustomMult = "custom_mult"
	// FieldLo holds the string denoting the lo field in the database.
	FieldLo = "lo"
	// FieldHi holds the string denoting the hi field in the database.
	FieldHi = "hi"
	// FieldTagID holds the string denoting the tag_id field in the database.
	FieldTagID = "tag_id"
	// FieldCreatorID holds the string denoting the creator_id field in the database.
	FieldCreatorID = "creator_id"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeTagg holds the string denoting the tagg edge name in mutations.
	EdgeTagg = "tagg"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "tasks"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "creator_id"
	// TaggTable is the table that holds the tagg relation/edge.
	TaggTable = "tasks"
	// TaggInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TaggInverseTable = "tags"
	// TaggColumn is the table column denoting the tagg relation/edge.
	TaggColumn = "tag_id"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldIcon,
	FieldTitle,
	FieldDescription,
	FieldDeadline,
	FieldEstimated,
	FieldImportance,
	FieldUrgency,
	FieldCustomMult,
	FieldLo,
	FieldHi,
	FieldTagID,
	FieldCreatorID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultIcon holds the default value on creation for the "icon" field.
	DefaultIcon int
	// IconValidator is a validator for the "icon" field. It is called by the builders before save.
	IconValidator func(int) error
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// EstimatedValidator is a validator for the "estimated" field. It is called by the builders before save.
	EstimatedValidator func(int) error
	// DefaultImportance holds the default value on creation for the "importance" field.
	DefaultImportance int8
	// ImportanceValidator is a validator for the "importance" field. It is called by the builders before save.
	ImportanceValidator func(int8) error
	// DefaultUrgency holds the default value on creation for the "urgency" field.
	DefaultUrgency int8
	// UrgencyValidator is a validator for the "urgency" field. It is called by the builders before save.
	UrgencyValidator func(int8) error
	// DefaultCustomMult holds the default value on creation for the "custom_mult" field.
	DefaultCustomMult float64
)
