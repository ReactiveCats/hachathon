package task

import (
	"context"
	"server/internal/config"
	"server/internal/domain"
	"server/internal/ent"
	taskent "server/internal/ent/task"
	"server/internal/platform"
)

type Service struct {
	client    *ent.TaskClient
	jwtConfig config.Jwt
}

func NewService(client *ent.Client, config config.Config) *Service {
	return &Service{
		client:    client.Task,
		jwtConfig: config.Jwt,
	}
}

func (s Service) ByID(ctx context.Context, taskID int) (*domain.Task, error) {
	task, err := s.client.Query().
		Where(taskent.ID(taskID)).
		WithCreator().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, platform.NotFound("Task is not found")
		}
		return nil, platform.WrapInternal(err)
	}

	return domain.TaskFromEnt(task), nil
}

func (s Service) Update(ctx context.Context, taskID int) (*domain.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, taskID int) error {
	_, err := s.client.Delete().
		Where(taskent.ID(taskID)).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return platform.NotFound("Task is not found")
		}
		return platform.WrapInternal(err)
	}

	return nil
}

func (s Service) Create(ctx context.Context, taskDTO domain.CreateTaskDTO) error {
	err := s.client.Create().
		SetTitle(taskDTO.Title).
		SetNillableIcon(taskDTO.Icon).
		SetNillableDescription(taskDTO.Description).
		SetNillablePriority((*taskent.Priority)(taskDTO.Priority)).
		SetNillableComplexity((*taskent.Complexity)(taskDTO.Complexity)).
		SetNillableDeadline(taskDTO.Deadline).
		SetNillableEstimated(taskDTO.Estimated).
		SetCreatorID(taskDTO.UserID).
		Exec(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return platform.NotFound("User not found")
		}
		return platform.WrapInternal(err)
	}
	return nil
}
