package domain

import (
	"context"
	"server/internal/ent"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Icon        int       `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Deadline    time.Time `json:"deadline"`
	Estimated   int       `json:"estimated,omitempty"`
	Complexity  string    `json:"complexity"`
	Priority    string    `json:"priority"`
	Creator     *User     `json:"-"`
}

func TaskFromEnt(task *ent.Task) *Task {
	if task == nil {
		return nil
	}

	return &Task{
		ID:          task.ID,
		CreatedAt:   task.CreatedAt,
		Icon:        task.Icon,
		Title:       task.Title,
		Description: task.Description,
		Deadline:    task.Deadline,
		Estimated:   task.Estimated,
		Complexity:  string(task.Complexity),
		Priority:    string(task.Priority),
		Creator:     UserFromEnt(task.Edges.Creator),
	}
}

func TasksFromEnt(slice []*ent.Task) []*Task {
	tasks := make([]*Task, len(slice))
	for i, e := range slice {
		tasks[i] = TaskFromEnt(e)
	}
	return tasks
}

type TaskService interface {
	ByID(ctx context.Context, taskID int) (*Task, error)
	Update(ctx context.Context, taskID TaskPutDTO) (*Task, error)
	Delete(ctx context.Context, taskID int) error
}

type TaskPutDTO struct {
	UserID          int        `json:"-"`
	TaskID          int        `json:"-"`
	Icon            *int       `json:"icon"`
	Title           string     `json:"title"`
	Description     *string    `json:"description,omitempty"`
	DeadlineDateStr *string    `json:"deadline"`
	Deadline        *time.Time `json:"-"`
	Estimated       *int       `json:"estimated,omitempty" `
	Complexity      *string    `json:"complexity" binding:"oneof=very_low low mid high very_high"`
	Priority        *string    `json:"priority" binding:"oneof=very_low low mid high very_high"`
}
