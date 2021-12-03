package domain

import (
	"context"
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

type TaskService interface {
	Delete(ctx context.Context, taskID int) error
}
