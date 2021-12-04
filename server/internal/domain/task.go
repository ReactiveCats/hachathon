package domain

import (
	"context"
	"server/internal/ent"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Icon        int       `json:"icon"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	Deadline    time.Time `json:"deadline"`
	Estimated   int       `json:"estimated,omitempty"`
	Complexity  int       `json:"complexity"`
	Priority    int       `json:"priority"`
	F           float64   `json:"f"`
	Lo          float64   `json:"lo"`
	Hi          float64   `json:"hi"`
	Creator     *User     `json:"-"`
}

type Question struct {
	Text          string `json:"text"`
	CompareTaskID int    `json:"compare_task_id"`
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
		Complexity:  int(task.Complexity),
		Priority:    int(task.Priority),
		F:           task.F,
		Lo:          task.Lo,
		Hi:          task.Hi,
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
	Fetch(ctx context.Context, dto GetTaskDTO) ([]*Task, error)

	Create(ctx context.Context, task CreateTaskDTO) (*Task, *Question, error)
	AnswerQuestion(ctx context.Context, params AnswerQuestionDTO) (*Question, error)
	AskQuestion(ctx context.Context, task *Task) (*Question, error)

	Update(ctx context.Context, dto TaskPutDTO) (*Task, error)

	Delete(ctx context.Context, taskID int) error
}

type GetTaskDTO struct {
	UserID     int     `json:"-"`
	Estimated  *int    `json:"estimated,omitempty"`
	Complexity *int    `json:"complexity" binding:"omitempty,max=10,min=0"`
	Priority   *int    `json:"priority" binding:"omitempty,max=10,min=0"`
	Order      *string `json:"order" binding:"omitempty,oneof=desc asc"`
	OrderBy    *string `json:"orderBy" binding:"omitempty,oneof=created_at deadline estimated complexity priority"`
}

type AnswerQuestionDTO struct {
	UserID        int `json:"-"`
	Response      int `json:"response"`
	TaskID        int `json:"-"`
	CompareTaskID int `json:"task_id"`
}

type CreateTaskDTO struct {
	UserID      int        `json:"-"`
	Icon        *int       `json:"icon" binding:"omitempty,max=16,min=0"`
	Title       string     `json:"title" binding:"required,max=64"`
	Description *string    `json:"description,omitempty" binding:"omitempty,max=256"`
	Deadline    *time.Time `json:"deadline"`
	Estimated   *int       `json:"estimated,omitempty" binding:"omitempty,max=86400,min=0"`
	Complexity  int        `json:"complexity" binding:"omitempty,max=10,min=0"`
	Priority    int        `json:"priority" binding:"omitempty,max=10,min=0"`
}

type CreateTaskAnswer struct {
	Task     *Task     `json:"task"`
	Question *Question `json:"question"`
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
	Complexity      int        `json:"complexity" binding:"max=10,min=0"`
	Priority        int        `json:"priority" binding:"max=10,min=0"`
}
