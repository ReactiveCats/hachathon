package domain

import (
	"context"
	"server/internal/ent"
	"time"
)

type Task struct {
	ID           int
	Title        string
	Description  string
	Priority     string
	Complexity   string
	HardDeadline time.Time
	SoftDeadline time.Time
	Status       string
}

func TaskFromEnt(task *ent.Task) *Task {
	if task == nil {
		return nil
	}

	return &Task{
		ID:           task.ID,
		Title:        task.Title,
		Description:  task.Description,
		Priority:     task.Priority,
		Complexity:   task.Complexity,
		HardDeadline: task.HardDeadline,
		SoftDeadline: task.SoftDeadline,
		Status:       task.Status,
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
	Delete(ctx context.Context, taskID int) error
}
