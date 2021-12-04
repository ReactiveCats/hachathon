package domain

import (
	"context"
	"math"
	"server/internal/ent"
	"time"
)

const (
	Epsilon = 0.01
)

type Task struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	Icon          int       `json:"icon"`
	Title         string    `json:"title"`
	Description   string    `json:"description,omitempty"`
	Deadline      time.Time `json:"deadline"`
	Estimated     int       `json:"estimated,omitempty"`
	Importance    int       `json:"importance"`
	Urgency       int       `json:"urgency"`
	CustomMult    float64   `json:"custom_mult"`
	Lo            float64   `json:"-"`
	Hi            float64   `json:"-"`
	Creator       *User     `json:"-"`
	FinalPriority *float64  `json:"priority,omitempty"`
}

type Tasks []*Task

const MULTIPLIER = 0.2

func (ts Tasks) Len() int {
	return len(ts)
}

func (ts Tasks) Less(i, j int) bool {
	return ts[i].P() < ts[j].P()
}

func (ts Tasks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (t *Task) P() float64 {
	a := t.F() * t.CustomMult * (float64(t.Importance) * MULTIPLIER) * (float64(t.Urgency) * MULTIPLIER)
	t.FinalPriority = &a
	return a
}

func (t Task) F() float64 {
	timeLeft := t.Deadline.Unix() - time.Now().Unix()
	return (math.Log10((float64(t.Estimated)/float64(timeLeft))+0.015) + 2) / 2
	//return math.Sqrt(float64(t.Estimated) / float64(timeLeft))
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
		Importance:  int(task.Importance),
		Urgency:     int(task.Urgency),
		CustomMult:  task.CustomMult,
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
	GenerateQuestion(ctx context.Context, task *Task) (*Question, error)

	Update(ctx context.Context, dto TaskPutDTO) (*Task, error)

	Delete(ctx context.Context, taskID int) error
}

type GetTaskDTO struct {
	UserID     int     `json:"-"`
	Estimated  *int    `json:"estimated,omitempty"`
	Importance *int    `json:"importance" binding:"omitempty,max=10,min=0"`
	Urgency    *int    `json:"urgency" binding:"omitempty,max=10,min=0"`
	Order      *string `json:"order" binding:"omitempty,oneof=desc asc"`
	OrderBy    *string `json:"orderBy" binding:"omitempty,oneof=created_at deadline estimated importance urgency"`
	Priority   *bool   `json:"priority,omitempty"`
}

type AnswerQuestionDTO struct {
	UserID        int `json:"-"`
	Response      int `json:"response" binding:"oneof=-1 0 1"` // -1,0,1
	TaskID        int `json:"-"`
	CompareTaskID int `json:"compare_task_id"`
}

type CreateTaskDTO struct {
	UserID      int        `json:"-"`
	Icon        *int       `json:"icon" binding:"omitempty,max=16,min=0"`
	Title       string     `json:"title" binding:"required,max=64"`
	Description *string    `json:"description,omitempty" binding:"omitempty,max=256"`
	Deadline    *time.Time `json:"deadline"`
	Estimated   *int       `json:"estimated,omitempty" binding:"omitempty,max=86400,min=0"`
	Importance  int        `json:"importance" binding:"omitempty,max=10,min=0"`
	Urgency     int        `json:"urgency" binding:"omitempty,max=10,min=0"`
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
	Importance      int        `json:"importance" binding:"max=10,min=0"`
	Urgency         int        `json:"urgency" binding:"max=10,min=0"`
}
